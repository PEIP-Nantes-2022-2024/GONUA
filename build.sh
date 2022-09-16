# Build for Linux (my default system)
go build -o calendar-linux
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o calendar-windows.exe
# Build for Mac
GOOS=darwin GOARCH=amd64 go build -o calendar-mac
