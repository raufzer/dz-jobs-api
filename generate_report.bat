@echo off

cd /d "%~dp0"

where gocov >nul 2>&1
if %errorlevel% neq 0 (
  echo Installing gocov...
  go install github.com/axw/gocov/gocov@latest
)
where go-cover-treemap >nul 2>&1
if %errorlevel% neq 0 (
  echo Installing go-cover-treemap...
  go install github.com/nikolaydubina/go-cover-treemap@latest
)

go test -coverprofile=coverage.out ./...


go-cover-treemap -coverprofile coverage.out > coverage.html

echo Colorful coverage report generated: coverage.html