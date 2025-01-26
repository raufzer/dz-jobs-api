# DZ Jobs API

DZ Jobs API is a backend service built with Golang using the Gin framework. It provides robust functionality for job posting and recruitment, integrating various external services and implementing modern software practices such as layered architecture, caching, and CI/CD pipelines.

## Table of Contents
- [DZ Jobs API](#dz-jobs-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Setup and Installation](#setup-and-installation)
    - [Prerequisites](#prerequisites)
    - [Installation Steps](#installation-steps)
  - [Environment Variables](#environment-variables)
  - [Running the Application](#running-the-application)
  - [Docker Setup and Usage](#docker-setup-and-usage)
    - [Prerequisites](#prerequisites-1)
    - [Building the Docker Image](#building-the-docker-image)
    - [Running the Application with Docker](#running-the-application-with-docker)
  - [Swagger Documentation Generation](#swagger-documentation-generation)
    - [Install Swag CLI](#install-swag-cli)
    - [Generate Swagger Docs](#generate-swagger-docs)
    - [Accessing Swagger UI](#accessing-swagger-ui)
  - [Testing](#testing)
  - [CI/CD Pipeline](#cicd-pipeline)
  - [API Documentation](#api-documentation)
    - [Swagger Documentation](#swagger-documentation)
    - [Postman Workspace](#postman-workspace)
  - [Project Structure](#project-structure)
  - [Troubleshooting](#troubleshooting)
  - [Contributing](#contributing)
  - [License](#license)

## Features
- RESTful API built with Gin framework
- PostgreSQL database integration
- Redis for caching
- External services:
  - **SendGrid**: Email notifications
  - **Google OAuth**: Authentication
  - **Cloudinary**: File uploads (profile pictures, resumes)
- Unit testing with `rapport`
- CI/CD pipeline with automated testing, building, and deployment

## Technologies Used
- **Programming Language**: Golang
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Caching**: Redis
- **External Services**: SendGrid, Google OAuth, Cloudinary
- **Containerization**: Docker
- **CI/CD**: GitHub Actions

## Setup and Installation

### Prerequisites
1. Golang (1.20 or above)
2. PostgreSQL
3. Redis
4. Docker
5. Git

### Installation Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/raufzer/dz-jobs-api.git
   cd dz-jobs-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory with required variables

## Environment Variables
Create a `.env` file with the following configuration:
```plaintext
# Server Configuration
BACK_END_DOMAIN=your-backend-domain.com
FRONT_END_DOMAIN=your-frontend-domain.com
SERVER_PORT=8080

# Database Configuration
DATABASE_URI=postgres://username:password@localhost:5432/dbname

# Redis Configuration
REDIS_URI=redis://localhost:6379
REDIS_PASSWORD=your-redis-password

# Security & Authentication
ACCESS_TOKEN_SECRET=your-access-token-secret
REFRESH_TOKEN_SECRET=your-refresh-token-secret
RESET_PASSWORD_TOKEN_SECRET=your-reset-password-token-secret
ACCESS_TOKEN_MAX_AGE=24h
REFRESH_TOKEN_MAX_AGE=168h
RESET_PASSWORD_TOKEN_MAX_AGE=1h

# External Services
SENDGRID_API_KEY=your-sendgrid-api-key
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_REDIRECT_URL=https://your-backend-domain.com/oauth/google/callback
CLOUDINARY_CLOUD_NAME=your-cloudinary-cloud-name
CLOUDINARY_API_KEY=your-cloudinary-api-key
CLOUDINARY_API_SECRET=your-cloudinary-api-secret

# Defaults
DEFAULT_PROFILE_PICTURE=https://your-cloudinary-url.com/default-profile-picture.jpg
DEFAULT_RESUME=https://your-cloudinary-url.com/default-resume.pdf

# Application Metadata
BUILD_VERSION=1.0.0
COMMIT_HASH=your-commit-hash
ENVIRONMENT=production
DOC_URL=https://your-backend-domain.com/docs
LAST_MIGRATION=timestamp-of-last-migration
HEALTH_URL=https://your-backend-domain.com/health
VERSION_URL=https://your-backend-domain.com/version
METRICS_URL=https://your-backend-domain.com/metrics

# Service Email
SERVICE_EMAIL=your-service-email@example.com
```

## Running the Application
Start the server:
```bash
go run cmd/server/main.go
```
The server will be accessible at http://localhost:9090.

## Docker Setup and Usage
- **Link**: [Docker Hub Repo](https://hub.docker.com/repository/docker/raufzer/dz-jobs-api-docker/)

### Prerequisites
- Docker installed
- Docker Compose (optional but recommended)

### Building the Docker Image
```bash
# Build the Docker image
docker build -t dz-jobs-api .

# Or using Docker Compose
docker-compose build
```

### Running the Application with Docker
```bash
# Run the container
docker run -p 9090:9090 --env-file .env dz-jobs-api

# Or using Docker Compose
docker-compose up
```

## Swagger Documentation Generation

### Install Swag CLI
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Generate Swagger Docs
```bash
# Generate Swagger documentation
swag init -g cmd/server/main.go
```

### Accessing Swagger UI
- Local URL: `http://localhost:9090/docs/index.html`
- Ensure Swagger annotations are added to route handlers

## Testing
Run unit tests:
```bash
go test ./...
```
Coverage and reporting managed with the `rapport` testing tool.

## CI/CD Pipeline
GitHub Actions workflow:
- Runs tests and linter
- Builds Docker images
- Pushes to container registry
- Deploys application

## API Documentation

### Swagger Documentation
- **Endpoint**: `/docs/index.html`
- **Features**: 
  - Interactive endpoint details
  - Request/response schemas
  - Direct API testing

### Postman Workspace
- **Link**: [DZ Jobs API Postman Workspace](reach-me-out-for-the-link)
- **Collections**: 
  - Authentication endpoints
  - Job posting endpoints
  - User management endpoints

## Project Structure
```
dz-jobs-api/
├── cmd/server
│   └── main.go          # Application entry point
├── internal/
│   ├── bootstrap/       # Application bootstrapping
│   ├── controllers/     # API request handlers  
│   ├── dto/             # Data transfer objects
│   ├── integrations/    # External service integrations
│   ├── middlewares/     # Middleware logic
│   ├── models/          # Domain models
│   ├── repositories/    # Data access layer
│   └── services/        # Business logic
├── docs/                # Swagger documentation
├── Dockerfile           # Docker image configuration
└── docker-compose.yml   # Docker Compose setup
```

## Troubleshooting
- Verify environment variables
- Check Docker and Go versions
- Ensure network ports are available
- Review logs for specific errors

## Contributing
1. Fork the repository
2. Create a feature branch
3. Commit changes
4. Push branch
5. Open pull request

## License
Abd Raouf Zerkhef
zerkhefraouf90@gmail.com