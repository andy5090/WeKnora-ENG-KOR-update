#!/bin/bash
# Unified version information retrieval script
# Supports local build and CI build environments

# Set default values
VERSION="unknown"
COMMIT_ID="unknown"
BUILD_TIME="unknown"
GO_VERSION="unknown"

# Get version number
if [ -f "VERSION" ]; then
    VERSION=$(cat VERSION | tr -d '\n\r')
fi

# Get commit ID
if [ -n "$GITHUB_SHA" ]; then
    # GitHub Actions environment
    COMMIT_ID="${GITHUB_SHA:0:7}"
elif command -v git >/dev/null 2>&1; then
    # Local environment
    COMMIT_ID=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
fi

# Get build time
if [ -n "$GITHUB_ACTIONS" ]; then
    # GitHub Actions environment, use standard time format
    BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S UTC')
else
    # Local environment
    BUILD_TIME=$(date -u '+%Y-%m-%d %H:%M:%S UTC')
fi

# Get Go version
if command -v go >/dev/null 2>&1; then
    GO_VERSION=$(go version 2>/dev/null || echo "unknown")
fi

# Output different formats based on parameter
case "${1:-env}" in
    "env")
        # Output environment variable format, escape values containing spaces
        echo "VERSION=$VERSION"
        echo "COMMIT_ID=$COMMIT_ID"
        echo "BUILD_TIME=\"$BUILD_TIME\""
        echo "GO_VERSION=\"$GO_VERSION\""
        ;;
    "json")
        # Output JSON format
        cat << EOF
{
  "version": "$VERSION",
  "commit_id": "$COMMIT_ID",
  "build_time": "$BUILD_TIME",
  "go_version": "$GO_VERSION"
}
EOF
        ;;
    "docker-args")
        # Output Docker build argument format
        echo "--build-arg VERSION_ARG=$VERSION"
        echo "--build-arg COMMIT_ID_ARG=$COMMIT_ID"
        echo "--build-arg BUILD_TIME_ARG=$BUILD_TIME"
        echo "--build-arg GO_VERSION_ARG=$GO_VERSION"
        ;;
    "ldflags")
        # Output Go ldflags format
        echo "-X 'github.com/Tencent/WeKnora/internal/handler.Version=$VERSION' -X 'github.com/Tencent/WeKnora/internal/handler.CommitID=$COMMIT_ID' -X 'github.com/Tencent/WeKnora/internal/handler.BuildTime=$BUILD_TIME' -X 'github.com/Tencent/WeKnora/internal/handler.GoVersion=$GO_VERSION'"
        ;;
    "info")
        # Output information format
        echo "Version: $VERSION"
        echo "Commit ID: $COMMIT_ID"
        echo "Build Time: $BUILD_TIME"
        echo "Go Version: $GO_VERSION"
        ;;
    *)
        echo "Usage: $0 [env|json|docker-args|ldflags|info]"
        echo "  env        - Output environment variable format (default)"
        echo "  json       - Output JSON format"
        echo "  docker-args - Output Docker build argument format"
        echo "  ldflags    - Output Go ldflags format"
        echo "  info       - Output information format"
        exit 1
        ;;
esac
