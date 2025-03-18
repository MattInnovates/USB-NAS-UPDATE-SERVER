@echo off
setlocal enabledelayedexpansion

echo Building usb-nas-update-server.exe...

:: Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% neq 0 (
    echo ERROR: Go is not installed or not in PATH.
    echo Please install Go from https://go.dev/dl/ and try again.
    pause
    exit /b 1
)

:: Build the executable
go build -o usb-nas-update-server.exe server.go

:: Check if the build succeeded
if %ERRORLEVEL% neq 0 (
    echo Build failed! Check for errors above.
    pause
    exit /b 1
) else (
    echo Build succeeded! usb-nas-update-server.exe is ready.
    pause
)

exit /b 0
