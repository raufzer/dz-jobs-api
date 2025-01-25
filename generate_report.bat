@echo off

:: Navigate to the root directory
cd /d "%~dp0"

:: Install gocov if not already installed
where gocov >nul 2>&1
if %errorlevel% neq 0 (
  echo Installing gocov...
  go install github.com/axw/gocov/gocov@latest
)

:: Install go-cover-treemap if not already installed
where go-cover-treemap >nul 2>&1
if %errorlevel% neq 0 (
  echo Installing go-cover-treemap...
  go install github.com/nikolaydubina/go-cover-treemap@latest
)

:: Run tests and generate coverage profile
go test -coverprofile=coverage.out ./...

:: Generate treemap visualization
go-cover-treemap -coverprofile coverage.out > coverage.html

echo Colorful coverage report generated: coverage.html