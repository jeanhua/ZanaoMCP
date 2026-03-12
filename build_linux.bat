@echo off
setlocal
set GOOS=linux
set GOARCH=amd64
go build -o build/zanao-mcp
endlocal