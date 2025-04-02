#!/bin/sh
set -e

echo "[ENTRYPOINT-PROD] Backend starting..."

# In production, environment variables should be provided by the environment
# (e.g., container platform, Kubernetes, etc.)

echo "[ENTRYPOINT-PROD] POSTGRES_HOST: $POSTGRES_HOST"
echo "[ENTRYPOINT-PROD] PORT: $PORT"

# Set up database connection string
export PG_URL="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable"

# The postgres service in docker-compose already has a healthcheck,
# but we'll add a quick check to make sure our connection works
echo "Verifying PostgreSQL connection..."
attempt=0
max_attempts=30

until PGPASSWORD="$POSTGRES_PASSWORD" psql -h ${POSTGRES_HOST} -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q' > /dev/null 2>&1 || [ $attempt -eq $max_attempts ]; do
  attempt=$((attempt+1))
  echo "PostgreSQL connection attempt $attempt/$max_attempts"
  sleep 2
done

if [ $attempt -eq $max_attempts ]; then
  echo "Error: Could not connect to PostgreSQL after multiple attempts"
  exit 1
fi

echo "Running database migrations..."
cd /app
# Create migrations directory if it doesn't exist
mkdir -p /app/internal/db/migrations

# Run the migrations
set +e
# Use migrate command directly
migrate -path /app/internal/db/migrations -database "$PG_URL" -verbose up
set -e
echo "Database migrations complete."

# Create config directory if it doesn't exist
mkdir -p /app/config

# Execute the Go application
echo "Starting Go server..."
exec /app/server