# Build for Linux (my default system)
go build -o build/calendar-linux
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o build/calendar-windows.exe
# Build for Mac
GOOS=darwin GOARCH=amd64 go build -o build/calendar-mac
