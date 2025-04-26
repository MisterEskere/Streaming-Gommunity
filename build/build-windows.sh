echo "Building for Windows..."
cd "$(dirname "$0")/.." # Navigate to the project root
GOOS=windows GOARCH=amd64 go build -o build/Streaming-Gommunity.exe ./src/main.go
echo "Build completed: build/Streaming-Gommunity.exe"