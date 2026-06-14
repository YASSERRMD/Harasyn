FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway ./cmd/gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o /policy-worker ./cmd/policy-worker
RUN CGO_ENABLED=0 GOOS=linux go build -o /session-worker ./cmd/session-worker
RUN CGO_ENABLED=0 GOOS=linux go build -o /audit-worker ./cmd/audit-worker

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /api /app/api
COPY --from=builder /gateway /app/gateway
COPY --from=builder /policy-worker /app/policy-worker
COPY --from=builder /session-worker /app/session-worker
COPY --from=builder /audit-worker /app/audit-worker

EXPOSE 8080 8081 8082 8083 8084

CMD ["/app/api"]
