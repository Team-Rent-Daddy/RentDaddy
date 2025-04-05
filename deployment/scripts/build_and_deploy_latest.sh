#!/bin/bash
#
# Build and deploy RentDaddy Docker images to AWS ECR and force new ECS deployment
#
# This script builds the backend, frontend, and/or documenso-worker Docker images
# and pushes them to Amazon ECR. It can also force a new deployment of the ECS services.
#
# Requirements: docker, aws-cli, buildx enabled
#

# Load utility functions
source "$(dirname "$0")/utils.sh"

# Initialize
init_script "RentDaddy Build and Deploy" docker aws

# Find project root directly to avoid capturing banner output
PROJECT_ROOT=$(find_project_root)

# Function to build and push backend image
build_backend() {
  local tag="${1:-latest}"
  log "Building backend with tag: $tag"
  
  # Load backend environment variables
  if [ -f "$PROJECT_ROOT/backend/.env.production.local" ]; then
    load_env "$PROJECT_ROOT/backend/.env.production.local"
  else
    log "Error: Cannot find backend/.env.production.local in expected locations"
    exit 1
  fi
  
  # Ensure required variables are set
  if [ -z "$AWS_ACCOUNT_ID" ] || [ -z "$AWS_REGION" ]; then
    log "Error: AWS_ACCOUNT_ID and AWS_REGION must be set in backend/.env.production.local"
    exit 1
  fi
  
  # Authenticate with ECR
  ecr_login
  
  log "Starting backend build..."
  
  # Record the start time
  local start_time=$(date +%s)
  
  # Get absolute path to backend directory
  local backend_dir="$PROJECT_ROOT/backend"
  log "Using backend directory: $backend_dir"
  
  # Build and push backend
  docker buildx build \
    --platform linux/amd64 \
    -t $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/rentdaddy/backend:$tag \
    -f $backend_dir/Dockerfile.prod \
    --push \
    $backend_dir | tee "$PROJECT_ROOT/deployment/backend-build.log"
  
  # Calculate build duration
  local end_time=$(date +%s)
  local duration=$((end_time - start_time))
  log "Backend build completed in ${duration}s"
  
  log "Backend image pushed to: $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/rentdaddy/backend:$tag"
}

# Function to build and push frontend
build_frontend() {
  local tag="${1:-prod}"
  log "Building frontend with tag: $tag"
  
  # Load frontend environment variables
  if [ -f "$PROJECT_ROOT/frontend/app/.env.production.local" ]; then
    load_env "$PROJECT_ROOT/frontend/app/.env.production.local"
  else
    log "Error: Cannot find frontend/app/.env.production.local in expected locations"
    exit 1
  fi
  
  # Ensure required variables are set
  if [ -z "$AWS_ACCOUNT_ID" ] || [ -z "$AWS_REGION" ]; then
    log "Error: AWS_ACCOUNT_ID and AWS_REGION must be set in frontend/app/.env.production.local"
    exit 1
  fi
  
  if [ -z "$VITE_CLERK_PUBLISHABLE_KEY" ] || [ -z "$VITE_BACKEND_URL" ]; then
    log "Error: VITE_CLERK_PUBLISHABLE_KEY and VITE_BACKEND_URL must be set in frontend/app/.env.production.local"
    exit 1
  fi
  
  # Authenticate with ECR
  ecr_login
  
  log "Starting frontend build..."
  
  # Record the start time
  local start_time=$(date +%s)
  
  # Get absolute path to frontend directory
  local frontend_dir="$PROJECT_ROOT/frontend/app"
  log "Using frontend directory: $frontend_dir"
  
  # Build and push frontend
  docker buildx build \
    --platform linux/amd64 \
    --progress=plain \
    --build-arg VITE_CLERK_PUBLISHABLE_KEY="$VITE_CLERK_PUBLISHABLE_KEY" \
    --build-arg VITE_BACKEND_URL="$VITE_BACKEND_URL" \
    --build-arg VITE_DOCUMENSO_PUBLIC_URL="${VITE_DOCUMENSO_PUBLIC_URL:-https://docs.${DOMAIN_NAME:-curiousdev.net}}" \
    --build-arg VITE_ENV="${VITE_ENV:-production}" \
    -t $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/rentdaddy/frontend:$tag \
    -f $frontend_dir/Dockerfile.prod \
    --push \
    $frontend_dir | tee "$PROJECT_ROOT/deployment/frontend-build.log"
  
  # Calculate build duration
  local end_time=$(date +%s)
  local duration=$((end_time - start_time))
  log "Frontend build completed in ${duration}s"
  
  log "Frontend image pushed to: $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/rentdaddy/frontend:$tag"
}

# Function to build and push documenso-worker
build_worker() {
  local tag="${1:-latest}"
  log "Building documenso-worker with tag: $tag"
  
  # Remember original directory
  local original_dir=$(pwd)
  
  # Load backend environment variables for AWS credentials
  load_env "$PROJECT_ROOT/backend/.env.production.local"
  
  # Ensure required variables are set
  if [ -z "$AWS_ACCOUNT_ID" ] || [ -z "$AWS_REGION" ]; then
    log "Error: AWS_ACCOUNT_ID and AWS_REGION must be set in backend/.env.production.local"
    exit 1
  fi
  
  # Make sure we're logged in to ECR
  ecr_login
  
  log "Starting documenso-worker build..."
  
  # Record the start time
  local start_time=$(date +%s)
  
  # Check if repository exists, create if not
  log "Checking if repository exists..."
  if ! aws ecr describe-repositories --repository-names "rentdaddy/documenso-worker" &>/dev/null; then
    log "Creating ECR repository rentdaddy/documenso-worker..."
    aws ecr create-repository --repository-name "rentdaddy/documenso-worker" --image-scanning-configuration scanOnPush=true
  fi
  
  # Find the worker directory using the utility function
  WORKER_DIR=$(get_worker_dir "$PROJECT_ROOT")
  if [ $? -ne 0 ] || [ -z "$WORKER_DIR" ]; then
    log "Error: Could not find documenso-worker directory"
    exit 1
  fi
  
  log "Using worker directory: $WORKER_DIR"
  
  # Build and push worker
  log "Building Docker image..."
  docker buildx build \
    --platform linux/amd64 \
    -t $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/rentdaddy/documenso-worker:$tag \
    -f "$WORKER_DIR/Dockerfile" \
    --push \
    "$WORKER_DIR" | tee "$PROJECT_ROOT/deployment/worker-build.log"
  
  # Calculate build duration
  local end_time=$(date +%s)
  local duration=$((end_time - start_time))
  log "Documenso worker build completed in ${duration}s"
  
  log "Worker image pushed to: $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/rentdaddy/documenso-worker:$tag"
  
  # Return to original directory
  cd "$original_dir"
}

# Function to force new deployment
force_deployment() {
  log "Forcing new deployment of all services..."
  
  # Ensure AWS credentials are set
  if [ -z "$AWS_ACCOUNT_ID" ] || [ -z "$AWS_REGION" ]; then
    log "Loading AWS configuration for deployment..."
    load_aws_config "$PROJECT_ROOT"
  fi
  
  # Create directory to store deployment info
  local timestamp=$(date +%Y%m%d_%H%M%S)
  local deploy_dir="$PROJECT_ROOT/deployment/simplified_terraform/ecs_deployment_$timestamp"
  mkdir -p "$deploy_dir"
  
  # Force new deployment for all services
  for svc_arn in $(aws ecs list-services --cluster rentdaddy-cluster --query "serviceArns[]" --output text); do
    local svc_name=$(basename "$svc_arn")
    log "📦 Forcing new deployment for $svc_name..."
    aws ecs update-service \
      --cluster rentdaddy-cluster \
      --service "$svc_name" \
      --force-new-deployment \
      > "$deploy_dir/${svc_name}.json"
  done
  
  log "All deployments have been initiated. Check AWS ECS console for status."
}

# Display help
show_help() {
  echo "Usage: $0 [OPTIONS]"
  echo ""
  echo "Options:"
  echo "  -b, --backend    Build and push backend image only"
  echo "  -f, --frontend   Build and push frontend image only"
  echo "  -w, --worker     Build and push documenso-worker image only"
  echo "  -a, --all        Build and push all images (default)"
  echo "  -d, --deploy     Force new deployment after building"
  echo "  -t, --tag TAG    Specify a custom tag (default: latest for backend/worker, prod for frontend)"
  echo "  -h, --help       Show this help message"
  echo ""
  echo "Example:"
  echo "  $0 --all --deploy         # Build all and deploy"
  echo "  $0 --backend --tag v1.2.0 # Build backend only with specific tag"
  echo "  $0 --worker               # Build documenso-worker only"
  echo ""
  echo "Note: This script loads environment variables from:"
  echo "  - backend/.env.production.local (for backend builds)"
  echo "  - frontend/app/.env.production.local (for frontend builds)"
  echo ""
  echo "Required variables in the environment files:"
  echo "  - AWS_ACCOUNT_ID: Your AWS account ID"
  echo "  - AWS_REGION: AWS region (e.g., us-east-2)"
  echo "  - VITE_CLERK_PUBLISHABLE_KEY: Clerk publishable key (for frontend)"
  echo "  - VITE_BACKEND_URL: Backend URL (for frontend)"
}

# Default values
BUILD_BACKEND=false
BUILD_FRONTEND=false
BUILD_WORKER=false
DEPLOY=false
BACKEND_TAG="latest"
FRONTEND_TAG="prod"
WORKER_TAG="latest"

# Parse command line arguments
if [ $# -eq 0 ]; then
  # Default behavior if no args: build all
  BUILD_BACKEND=true
  BUILD_FRONTEND=true
  BUILD_WORKER=true
else
  while [ $# -gt 0 ]; do
    case "$1" in
      -b|--backend)
        BUILD_BACKEND=true
        shift
        ;;
      -f|--frontend)
        BUILD_FRONTEND=true
        shift
        ;;
      -w|--worker)
        BUILD_WORKER=true
        shift
        ;;
      -a|--all)
        BUILD_BACKEND=true
        BUILD_FRONTEND=true
        BUILD_WORKER=true
        shift
        ;;
      -d|--deploy)
        DEPLOY=true
        shift
        ;;
      -t|--tag)
        BACKEND_TAG="$2"
        FRONTEND_TAG="$2"
        WORKER_TAG="$2"
        shift 2
        ;;
      -h|--help)
        show_help
        exit 0
        ;;
      *)
        echo "Unknown option: $1"
        show_help
        exit 1
        ;;
    esac
  done
fi

# Build the specified components
if [ "$BUILD_BACKEND" = true ]; then
  build_backend $BACKEND_TAG
fi

if [ "$BUILD_FRONTEND" = true ]; then
  build_frontend $FRONTEND_TAG
fi

if [ "$BUILD_WORKER" = true ]; then
  build_worker $WORKER_TAG
fi

# Force deployment if requested
if [ "$DEPLOY" = true ]; then
  force_deployment
fi

log "Build and deployment process completed successfully!"