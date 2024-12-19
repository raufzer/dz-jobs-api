# Set the base image to create the image for Go app
FROM golang:1.23.2-alpine

# Create a user with permissions to run the app
# -S -> create a system user
# -G -> add the user to a group
# This is done to avoid running the app as root
# If the app is run as root, any vulnerability in the app can be exploited to gain access to the host system
# It's a good practice to run the app as a non-root user
RUN addgroup app && adduser -S -G app app

# Set the user to run the app
USER app

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum to the working directory
# This is done before copying the rest of the files to take advantage of Docker’s cache
# If the go.mod and go.sum files haven’t changed, Docker will use the cached dependencies
COPY go.mod go.sum ./

# Sometimes the ownership of the files in the working directory is changed to root
# and thus the app can't access the files and throws an error -> EACCES: permission denied
# To avoid this, change the ownership of the files to the root user
USER root

# Change the ownership of the /app directory to the app user
# chown -R <user>:<group> <directory>
# chown command changes the user and/or group ownership of for given file.
RUN chown -R app:app .

# Change the user back to the app user
USER app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to tell Docker that the container listens on the specified network ports at runtime
EXPOSE 9090

# Command to run the app
CMD ["./main"]