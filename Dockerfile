FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /usr/app

RUN apk --no-cache add ca-certificates

# Copy the Go module files and download dependencies (optional)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container's working directory
COPY . .

# Build the GoLang Gin service
RUN go build -o main ./cmd

FROM scratch

WORKDIR /usr/app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/app/app.env app.env
COPY --from=builder /usr/app/main .

EXPOSE 8080

# Command to run the application
CMD ["./main"]
