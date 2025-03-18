package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/careecodes/RentDaddy/internal/db"
	gen "github.com/careecodes/RentDaddy/internal/db/generated"
	"github.com/careecodes/RentDaddy/pkg/handlers"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

var items = make(map[string]Item)

func PutItemHandler(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")
	var updatedItem Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if itemID != updatedItem.ID {
		http.Error(w, "ID in path and body do not match", http.StatusBadRequest)
		return
	}
	if _, ok := items[itemID]; !ok {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	items[itemID] = updatedItem
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedItem)
}

// QuickDump is a function that dumps the request to the console for debugging purposes
//
//	func QuickDump(r *http.Request) {
//		dump, err := httputil.DumpRequest(r, true)
//		if err != nil {
//			http.Error(w, "Failed to dump request", http.StatusInternalServerError)
//			return
//		}
//		fmt.Printf("Request dump: %s\n", dump)
//	}

func main() {
	// OS signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dbUrl := os.Getenv("PG_URL")

	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")

	// Initialize the database

	ctx := context.Background()

	queries, pool, err := db.ConnectDB(ctx, dbUrl)
	if err != nil {
		log.Fatalf("[DB] Failed initializing: %v", err)
	}
	defer pool.Close()

	// Initialize Clerk with your secret key
	clerk.SetKey(clerkSecretKey)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Webhooks
	r.Post("/webhooks/clerk", func(w http.ResponseWriter, r *http.Request) {
		handlers.ClerkWebhookHandler(w, r, pool, queries)
	})

	// User Router
	userHandler := handlers.NewUserHandler(pool, queries)
	
	// Initialize locker handler
	lockerHandler := handlers.NewLockerHandler(pool, queries)

	// Tenants Routes
	r.Route("/tenants", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userHandler.GetAllUsers(w, r, gen.RoleTenant)
		})
		r.Get("/{clerk_id}", userHandler.GetTenantByClerkId)
		r.Patch("/{clerk_id}/credentials", userHandler.UpdateTenantCredentials)
	})
	// Admin Routes
	r.Route("/admins", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userHandler.GetAllUsers(w, r, gen.RoleAdmin)
		})
		r.Post("/tenant_invite/{clerk_id}/{tenant_email}/{tenant_unit_number}", userHandler.InviteTenant)
		r.Get("/{clerk_id}", userHandler.GetAdminByClerkId)

		r.Route("/lockers", func(r chi.Router) {
			// Get all lockers with pagination
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				log.Println("List Lockers")
				lockerHandler.GetLockers(w, r)
			})
			// Create many lockers (used for the initial apartment setup)
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				log.Println("Creating Multiple Lockers")
				lockerHandler.CreateManyLockers(w, r)
			})
			
			// Routes for specific locker operations
			r.Route("/{id}", func(r chi.Router) {
				// Get single locker
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					log.Println("Get Locker")
					lockerHandler.GetLocker(w, r)
				})
		
				// Update locker (user/status or access code)d
				r.Patch("/", func(w http.ResponseWriter, r *http.Request) {
					log.Println("Update Locker")
					lockerHandler.UpdateLocker(w, r)
				})
			})
		})
	})



	r.Get("/test/get", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success in get"))
	})
	// Sample data
	items["1"] = Item{ID: "1", Value: "initial value"}

	r.Post("/test/post", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v",items)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Printf("%s", body)
		fmt.Printf("post success")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		w.Write([]byte("Success in post"))
	})

	r.Put("/test/put", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v",items)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Printf("%s", body)
		fmt.Printf("put success")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		w.Write([]byte("Success in put"))
	})

	r.Delete("/test/delete", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v",items)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Printf("%s", body)
		fmt.Printf("delete success")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	r.Patch("/test/patch", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%v",items)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		fmt.Printf("%s", body)
		fmt.Printf("patch success")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	r.Put("/test/clerk/update-username", func(w http.ResponseWriter, r *http.Request) {
		// QuickDump(r) // Uncomment to see the request dump

		// Define a struct to parse the incoming JSON
		type UpdateUsernameRequest struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		}

		// Set the request body to the struct so that we can parse the request body
		var updateReq UpdateUsernameRequest

		// Parse the request body
		if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Log the parsed request
		log.Printf("Received update request - ID: %s, Username: %s", updateReq.ID, updateReq.Username)

		// Check if ID is provided
		if updateReq.ID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		log.Printf("Updating user with ID: %s", updateReq.ID)

		// Update the user with the provided ID and username
		resource, err := user.Update(r.Context(), updateReq.ID, &user.UpdateParams{
			Username: clerk.String(updateReq.Username),
		})
		if err != nil {
			log.Printf("Error updating user: %v", err)
			http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("User updated successfully: %v", resource.ID)

		// Return the updated user as JSON using the response writer and the resource
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resource)
	})
	// End of Clerk Routes

	// Server config
	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server
	go func() {
		log.Printf("Server is running on port %s....\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Block until we reveive an interrupt signal
	<-sigChan
	log.Println("shutting down server...")

	// Gracefully shutdown the server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
}
