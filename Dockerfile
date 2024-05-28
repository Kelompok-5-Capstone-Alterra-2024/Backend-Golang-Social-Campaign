# Use the official Golang image to create a build artifact.
FROM golang:1.22 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code and the .env file to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use the official Alpine image for a lean production container.
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp .
# COPY --from=builder /app/.env .

# Optionally, if you want to include the .env file:
# COPY --from=builder /app/.env .

# Set default environment variables for the database connection
# These can be overridden by docker run -e
ENV DB_HOST=35.198.212.107
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=password
ENV DB_NAME=capstone5

ARG SMTP_HOST
ARG SMTP_PORT
ARG SMTP_USER
ARG SMTP_PASS

ENV SMTP_HOST=${SMTP_HOST}
ENV SMTP_PORT=${SMTP_PORT}
ENV SMTP_USER=${SMTP_USER}
ENV SMTP_PASS=${SMTP_PASS}

# Run the web service on container startup.
CMD ["./myapp"]