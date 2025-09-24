FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /docker-go-app ./cmd/main.go

EXPOSE 8080

CMD ["/docker-go-app"]
