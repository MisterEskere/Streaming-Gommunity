echo "Building for Linux..."
cd "$(dirname "$0")/.." # Navigate to the project root
GOOS=linux GOARCH=amd64 go build -o build/Streaming-Gommunity ./src/main.go
echo "Build completed: build/Streaming-Gommunity"