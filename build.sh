#!/bin/bash

# Clear existing build directories
rm -rf builds
mkdir builds

GOOS=linux GOARCH=amd64 go build -o builds/mediaStorageServer_linux_amd64 cmd/mediaStorageServer/main.go
echo "Linux amd64 built..."
GOOS=linux GOARCH=386 go build -o builds/mediaStorageServer_linux_386 cmd/mediaStorageServer/main.go
echo "Linux amd32 built..."
GOOS=linux GOARCH=arm64 go build -o builds/mediaStorageServer_linux_arm64 cmd/mediaStorageServer/main.go
echo "Linux arm64 built..."
GOOS=linux GOARCH=arm go build -o builds/mediaStorageServer_linux_arm cmd/mediaStorageServer/main.go
echo "Linux arm32 built..."

GOOS=windows GOARCH=amd64 go build -o builds/mediaStorageServer_windows_amd64 cmd/mediaStorageServer/main.go
echo "Windows amd64 built..."
GOOS=windows GOARCH=386 go build -o builds/mediaStorageServer_windows_386 cmd/mediaStorageServer/main.go
echo "Windows amd32 built..."
GOOS=windows GOARCH=arm64 go build -o builds/mediaStorageServer_windows_arm64 cmd/mediaStorageServer/main.go
echo "Windows arm64 built..."
GOOS=windows GOARCH=arm go build -o builds/mediaStorageServer_windows_arm cmd/mediaStorageServer/main.go
echo "Windows arm32 built..."

GOOS=darwin GOARCH=amd64 go build -o builds/mediaStorageServer_darwin_amd64 cmd/mediaStorageServer/main.go
echo "MacOS amd64 built..."
GOOS=darwin GOARCH=arm64 go build -o builds/mediaStorageServer_darwin_arm64 cmd/mediaStorageServer/main.go
echo "MacOS arm64 built..."

# Create compressed source file
sourceDir="."
output="builds/source.zip"
tarOutput="builds/source.tar.gz"
exclude=("builds" ".git")

filesToZip=$(find "$sourceDir" -type f \( ! -path "*${exclude[0]}*" -a ! -path "*${exclude[1]}*" \))

zip -r "$output" $filesToZip
tar -czf "$tarOutput" -C "$sourceDir" $filesToZip

echo "Zip archives created..."

# Create compressed archives
# cd builds
# tar -czf mediaStorageServer_linux_amd64.tar.gz mediaStorageServer_linux_amd64
# tar -czf mediaStorageServer_linux_386.tar.gz mediaStorageServer_linux_386
# zip mediaStorageServer_windows_amd64.zip mediaStorageServer_windows_amd64.exe
# zip mediaStorageServer_windows_386.zip mediaStorageServer_windows_386.exe
# tar -czf mediaStorageServer_darwin_amd64.tar.gz mediaStorageServer_darwin_amd64
# tar -czf mediaStorageServer_darwin_386.tar.gz mediaStorageServer_darwin_386
# cd ..

echo "Builds completed and compressed archives created in the 'builds' directory."

