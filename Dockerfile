FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go install github.com/cespare/reflex@latest

EXPOSE 4000

CMD go test --cover -v ./...