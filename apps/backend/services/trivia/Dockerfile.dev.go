FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3002

CMD ["go", "run", "cmd/trivia/main.go"]