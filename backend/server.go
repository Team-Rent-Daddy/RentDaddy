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
)

func main() {
	// OS signal channel
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	dbUrl := os.Getenv("PG_URL")
	if dbUrl == "" {
		log.Fatal("[ENV] Error: No Database url")
	}
	// Get the secret key from the environment variable
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")

	if clerkSecretKey == "" {
		log.Fatal("[ENV] CLERK_SECRET_KEY environment vars are required")
	}
	webhookSecret := os.Getenv("CLERK_WEBHOOK")

	if webhookSecret == "" {
		log.Fatal("[ENV] CLERK_WEBHOOK environment vars are required")
	}

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
	// Tenants Routes
	r.Route("/tenants", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userHandler.GetAllUsers(w, r, gen.RoleTenant)
		})
		// r.Post("/", userHandler.CreateTenant)
		r.Get("/{clerk_id}", userHandler.GetTenantByClerkId)
		r.Patch("/{clerk_id}/credentials", userHandler.UpdateTenantCredentials)
	})
	// Admin Routes
	r.Route("/admins", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userHandler.GetAllUsers(w, r, gen.RoleAdmin)
		})
		// r.Post("/", userHandler.CreateAdmin)
		r.Get("/{clerk_id}", userHandler.GetAdminByClerkId)
	})

	r.Get("/test/get", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success in get"))
	})

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
