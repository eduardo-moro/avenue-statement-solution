FROM golang:1.24-alpine

RUN apk add --no-cache curl

WORKDIR /app

# Copy everything
COPY . .

# Download dependencies
RUN go work sync

# Build the application
RUN go build -o /docker-go-app ./services/api-write/cmd/main.go

EXPOSE 8080

CMD ["/docker-go-app"]
