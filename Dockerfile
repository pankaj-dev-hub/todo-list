# Use the official Golang image as the base image for building
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the container's workspace
COPY . .

# Build both microservices
RUN go build -o todo ./cmd/todo 
RUN go build -o user ./cmd/user 

# Expose the necessary ports
EXPOSE 8080
EXPOSE 8081

# Command to run the services (you can customize the command based on your needs)

# Entrypoint script to run both services
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

