echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o build/Streaming-Gommunity.exe ./src/main.go
echo "Build completed: build/Streaming-Gommunity.exe"