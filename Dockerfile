FROM golang:1.21.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# ENV GIN_MODE=release

COPY . .

RUN go build -o main ./cmd/go-backend-banco

EXPOSE 8080

CMD ["./main"]
