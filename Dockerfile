FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init -g internal/app/app.go -o docs

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/db/migrations ./db/migrations

RUN adduser -D -g '' appuser
RUN chown -R appuser:appuser /app
USER appuser

CMD ["./app"] 