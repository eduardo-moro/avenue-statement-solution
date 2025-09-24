FROM golang:1.24-alpine

WORKDIR /app

# Copy everything
COPY . .

# Download dependencies
RUN go work sync

# Build the application
RUN go build -o /docker-go-app ./services/client-tui/cmd/main.go

EXPOSE 8080

CMD ["/docker-go-app"]
