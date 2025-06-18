FROM golang:1.25rc1-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

EXPOSE ${PORT}

CMD ["./server"]
