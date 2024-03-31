FROM golang:1.21.8 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o api .

FROM golang:1.21.8 AS runner

COPY --from=builder /app/api /

ENTRYPOINT ["/api"]
