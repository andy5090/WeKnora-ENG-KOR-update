#!/bin/bash
# Development environment startup script - only starts infrastructure, app and frontend need to be run manually locally

# Set colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No color

# Get project root directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

# Log functions
log_info() {
    printf "%b\n" "${BLUE}[INFO]${NC} $1"
}

log_success() {
    printf "%b\n" "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    printf "%b\n" "${RED}[ERROR]${NC} $1"
}

log_warning() {
    printf "%b\n" "${YELLOW}[WARNING]${NC} $1"
}

# Select available Docker Compose command
DOCKER_COMPOSE_BIN=""
DOCKER_COMPOSE_SUBCMD=""

detect_compose_cmd() {
    if docker compose version &> /dev/null; then
        DOCKER_COMPOSE_BIN="docker"
        DOCKER_COMPOSE_SUBCMD="compose"
        return 0
    fi
    if command -v docker-compose &> /dev/null; then
        if docker-compose version &> /dev/null; then
            DOCKER_COMPOSE_BIN="docker-compose"
            DOCKER_COMPOSE_SUBCMD=""
            return 0
        fi
    fi
    return 1
}

# Show help information
show_help() {
    printf "%b\n" "${GREEN}WeKnora Development Environment Script${NC}"
    echo "Usage: $0 [command] [options]"
    echo ""
    echo "Commands:"
    echo "  start      Start infrastructure services (postgres, redis, docreader)"
    echo "  stop       Stop all services"
    echo "  restart    Restart all services"
    echo "  logs       View service logs"
    echo "  status     View service status"
    echo "  app        Start backend application (run locally)"
    echo "  frontend   Start frontend development server (run locally)"
    echo "  help       Show this help information"
    echo ""
    echo "Optional Profiles (for start command):"
    echo "  --minio    Start MinIO object storage"
    echo "  --qdrant   Start Qdrant vector database"
    echo "  --neo4j    Start Neo4j graph database"
    echo "  --jaeger   Start Jaeger distributed tracing"
    echo "  --full     Start all optional services"
    echo ""
    echo "Examples:"
    echo "  $0 start                    # Start basic services"
    echo "  $0 start --qdrant           # Start basic services + Qdrant"
    echo "  $0 start --qdrant --jaeger  # Start basic services + Qdrant + Jaeger"
    echo "  $0 start --full             # Start all services"
    echo "  $0 app                      # Start backend in another terminal"
    echo "  $0 frontend                 # Start frontend in another terminal"
}

# Check Docker
check_docker() {
    if ! command -v docker &> /dev/null; then
        log_error "Docker is not installed, please install Docker first"
        return 1
    fi
    
    if ! detect_compose_cmd; then
        log_error "Docker Compose not detected"
        return 1
    fi
    
    if ! docker info &> /dev/null; then
        log_error "Docker service is not running"
        return 1
    fi
    
    return 0
}

# Start infrastructure services
start_services() {
    log_info "Starting development environment infrastructure services..."
    
    check_docker
    if [ $? -ne 0 ]; then
        return 1
    fi
    
    cd "$PROJECT_ROOT"
    
    # Check .env file
    if [ ! -f ".env" ]; then
        log_error ".env file does not exist, please create it first"
        return 1
    fi
    
    # Parse profile parameters
    shift  # Remove "start" command itself
    PROFILES="--profile full"
    ENABLED_SERVICES=""
    
    while [ $# -gt 0 ]; do
        case "$1" in
            --minio)
                PROFILES="$PROFILES --profile minio"
                ENABLED_SERVICES="$ENABLED_SERVICES minio"
                ;;
            --qdrant)
                PROFILES="$PROFILES --profile qdrant"
                ENABLED_SERVICES="$ENABLED_SERVICES qdrant"
                ;;
            --neo4j)
                PROFILES="$PROFILES --profile neo4j"
                ENABLED_SERVICES="$ENABLED_SERVICES neo4j"
                ;;
            --jaeger)
                PROFILES="$PROFILES --profile jaeger"
                ENABLED_SERVICES="$ENABLED_SERVICES jaeger"
                ;;
            --full)
                PROFILES="--profile full"
                ENABLED_SERVICES="minio qdrant neo4j jaeger"
                break
                ;;
            *)
                log_warning "Unknown parameter: $1"
                ;;
        esac
        shift
    done
    
    # Start services
    "$DOCKER_COMPOSE_BIN" $DOCKER_COMPOSE_SUBCMD -f docker-compose.dev.yml $PROFILES up -d
    
    if [ $? -eq 0 ]; then
        log_success "Infrastructure services started"
        echo ""
        log_info "Service access URLs:"
        echo "  - PostgreSQL:    localhost:5432"
        echo "  - Redis:         localhost:6379"
        echo "  - DocReader:     localhost:50051"
        
        # Show additional services based on enabled profiles
        if [[ "$ENABLED_SERVICES" == *"minio"* ]]; then
            echo "  - MinIO:         localhost:9000 (Console: localhost:9001)"
        fi
        if [[ "$ENABLED_SERVICES" == *"qdrant"* ]]; then
            echo "  - Qdrant:        localhost:6333 (gRPC: localhost:6334)"
        fi
        if [[ "$ENABLED_SERVICES" == *"neo4j"* ]]; then
            echo "  - Neo4j:         localhost:7474 (Bolt: localhost:7687)"
        fi
        if [[ "$ENABLED_SERVICES" == *"jaeger"* ]]; then
            echo "  - Jaeger:        localhost:16686"
        fi
        
        echo ""
        log_info "Next steps:"
        printf "%b\n" "${YELLOW}1. Run backend in new terminal:${NC} make dev-app"
        printf "%b\n" "${YELLOW}2. Run frontend in new terminal:${NC} make dev-frontend"
        return 0
    else
        log_error "Service startup failed"
        return 1
    fi
}

# Stop services
stop_services() {
    log_info "Stopping development environment services..."
    
    check_docker
    if [ $? -ne 0 ]; then
        return 1
    fi
    
    cd "$PROJECT_ROOT"
    "$DOCKER_COMPOSE_BIN" $DOCKER_COMPOSE_SUBCMD -f docker-compose.dev.yml down
    
    if [ $? -eq 0 ]; then
        log_success "All services stopped"
        return 0
    else
        log_error "Service stop failed"
        return 1
    fi
}

# Restart services
restart_services() {
    stop_services
    sleep 2
    start_services
}

# View logs
show_logs() {
    cd "$PROJECT_ROOT"
    "$DOCKER_COMPOSE_BIN" $DOCKER_COMPOSE_SUBCMD -f docker-compose.dev.yml logs -f
}

# View status
show_status() {
    cd "$PROJECT_ROOT"
    "$DOCKER_COMPOSE_BIN" $DOCKER_COMPOSE_SUBCMD -f docker-compose.dev.yml ps
}

# Start backend application (local)
start_app() {
    log_info "Starting backend application (local development mode)..."
    
    cd "$PROJECT_ROOT"
    
    # Check if Go is installed
    if ! command -v go &> /dev/null; then
        log_error "Go is not installed"
        return 1
    fi
    
    # Load environment variables (use set -a to ensure all variables are exported)
    if [ -f ".env" ]; then
        log_info "Loading .env file..."
        set -a
        source .env
        set +a
    else
        log_error ".env file does not exist, please create configuration file first"
        return 1
    fi
    
    # Set local development environment variables (override Docker container addresses)
    export DB_HOST=localhost
    export DOCREADER_ADDR=localhost:50051
    export MINIO_ENDPOINT=localhost:9000
    export REDIS_ADDR=localhost:6379
    export OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317
    export NEO4J_URI=bolt://localhost:7687
    export QDRANT_HOST=localhost
    
    # Ensure required environment variables are set
    if [ -z "$DB_DRIVER" ]; then
        log_error "DB_DRIVER environment variable is not set, please check .env file"
        return 1
    fi
    
    log_info "Environment variables set, starting application..."
    log_info "Database address: $DB_HOST:${DB_PORT:-5432}"
    
    # Check if Air (hot reload tool) is installed
    if command -v air &> /dev/null; then
        log_success "Air detected, starting in hot reload mode..."
        log_info "Go code changes will automatically recompile and restart"
        air
    else
        log_info "Air not detected, starting in normal mode"
        log_warning "Tip: Install Air to enable automatic restart on code changes"
        log_info "Install command: go install github.com/air-verse/air@latest"
        # Run application
        go run cmd/server/main.go
    fi
}

# Start frontend (local)
start_frontend() {
    log_info "Starting frontend development server..."
    
    cd "$PROJECT_ROOT/frontend"
    
    # Check if npm is installed
    if ! command -v npm &> /dev/null; then
        log_error "npm is not installed"
        return 1
    fi
    
    # Check if dependencies are installed
    if [ ! -d "node_modules" ]; then
        log_warning "node_modules does not exist, installing dependencies..."
        npm install
    fi
    
    log_info "Starting Vite development server..."
    log_info "Frontend will run on http://localhost:5173"
    
    # Run development server
    npm run dev
}

# Parse command
CMD="${1:-help}"
case "$CMD" in
    start)
        start_services "$@"
        ;;
    stop)
        stop_services
        ;;
    restart)
        restart_services
        ;;
    logs)
        show_logs
        ;;
    status)
        show_status
        ;;
    app)
        start_app
        ;;
    frontend)
        start_frontend
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        log_error "Unknown command: $CMD"
        show_help
        exit 1
        ;;
esac

exit 0

