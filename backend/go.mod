module github.com/careecodes/RentDaddy

go 1.23.0

toolchain go1.24.0

require github.com/go-chi/chi/v5 v5.2.1 // indirect gotcha

require (
	github.com/clerk/clerk-sdk-go/v2 v2.2.0
	github.com/go-chi/cors v1.2.1
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v5 v5.7.2
	github.com/svix/svix-webhooks v1.61.3
)

require (
	github.com/go-jose/go-jose/v3 v3.0.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
