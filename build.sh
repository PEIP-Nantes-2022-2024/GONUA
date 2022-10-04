# Build for Linux (my default system)
echo "Building for Linux..."
go build -o build/calendar-linux
echo "Built for Linux"
# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o build/calendar-windows.exe
echo "Built for Windows"
# Build for Mac
echo "Building for Mac..."
GOOS=darwin GOARCH=amd64 go build -o build/calendar-mac
echo "Built for Mac"
