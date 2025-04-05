#!/bin/bash
set -e

# Log function for better output
log() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') - $1"
}

# Function to find project root (works from any subdirectory)
find_project_root() {
  # Try to use git to find the repository root
  if git rev-parse --show-toplevel &> /dev/null; then
    git rev-parse --show-toplevel
  else
    # Fallback if not in a git repository
    echo "$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
  fi
}

# Function to load environment variables from a file
load_env() {
  local env_file="$1"
  if [ -f "$env_file" ]; then
    log "Loading environment variables from $env_file"
    set -a # automatically export all variables
    source "$env_file"
    set +a
  else
    log "Warning: Environment file $env_file not found!"
    return 1
  fi
}

# Function to check AWS CLI version compatibility
check_aws_cli_version() {
  local CLI_VERSION=$(aws --version 2>&1 | cut -d/ -f2 | cut -d' ' -f1)
  local MAJOR_VERSION=$(echo $CLI_VERSION | cut -d. -f1)
  
  if [[ "$MAJOR_VERSION" -lt 2 ]]; then
    log "Warning: AWS CLI version $CLI_VERSION detected"
    log "Some functionality may not be available with AWS CLI v1"
    log "Using alternative log retrieval method for compatibility"
    return 1
  fi
  
  return 0
}

# Find the project root
PROJECT_ROOT=$(find_project_root)
log "Project root directory: $PROJECT_ROOT"

# Navigate to the project root
cd "$PROJECT_ROOT"

# Load environment variables
load_env "$PROJECT_ROOT/backend/.env.production.local"

# Set default pattern and log group
FILTER_PATTERN=${1:-"documenso-worker"}
LOG_GROUP="/ecs/rentdaddy-documenso"

log "Monitoring Documenso worker logs with filter pattern: '$FILTER_PATTERN'"
log "Press Ctrl+C to exit"

# Check AWS CLI version
if check_aws_cli_version; then
  log "Using AWS CLI v2 tail command"
  # Modern method with AWS CLI v2
  aws logs tail "$LOG_GROUP" --filter-pattern "$FILTER_PATTERN" --follow
else
  log "Using AWS CLI v1 compatible method"
  # Alternative method for AWS CLI v1 compatibility
  
  # Create a temp file for the last event timestamp
  LAST_TIMESTAMP_FILE=$(mktemp)
  echo "0" > "$LAST_TIMESTAMP_FILE"  # Start with timestamp 0 (beginning of logs)
  
  # Monitor interval in seconds
  INTERVAL=5
  
  # Cleanup on exit
  trap 'rm -f "$LAST_TIMESTAMP_FILE"; log "Exiting log monitor"; exit 0' EXIT INT TERM
  
  log "Polling logs every $INTERVAL seconds..."
  
  # Continuous polling loop
  while true; do
    LAST_TIMESTAMP=$(cat "$LAST_TIMESTAMP_FILE")
    
    # Get log events after the last timestamp we've seen
    EVENTS=$(aws logs filter-log-events \
      --log-group-name "$LOG_GROUP" \
      --filter-pattern "$FILTER_PATTERN" \
      --start-time "$LAST_TIMESTAMP" \
      --query 'events[*].[timestamp,message]' \
      --output text)
    
    if [[ -n "$EVENTS" ]]; then
      echo "$EVENTS" | while read -r TIMESTAMP MESSAGE; do
        # Convert timestamp to human-readable format
        if [[ "$OSTYPE" == "darwin"* ]]; then
          # macOS
          DATE=$(date -r $((TIMESTAMP/1000)) "+%Y-%m-%d %H:%M:%S")
        else
          # Linux
          DATE=$(date -d @$((TIMESTAMP/1000)) "+%Y-%m-%d %H:%M:%S")
        fi
        echo "[$DATE] $MESSAGE"
        
        # Update the largest timestamp seen
        if [[ $TIMESTAMP -gt $LAST_TIMESTAMP ]]; then
          echo $TIMESTAMP > "$LAST_TIMESTAMP_FILE"
        fi
      done
    fi
    
    sleep $INTERVAL
  done
fi