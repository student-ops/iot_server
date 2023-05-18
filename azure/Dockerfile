# Use the official Microsoft Azure CLI Docker image
FROM mcr.microsoft.com/azure-cli

# Set the working directory
WORKDIR /app

# Copy all files from the current directory on the host to the working directory in the container
COPY ../ ./

# Install Docker CLI
RUN apk add --no-cache docker-cli
