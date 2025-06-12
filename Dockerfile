FROM golang:1.24.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

EXPOSE 8080

RUN go build -o user-api ./main.go

CMD ["./user-api"]
