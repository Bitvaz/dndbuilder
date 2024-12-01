FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM alpine:latest
WORKDIR /app
COPY --from=builder /docker-gs-ping /docker-gs-ping
COPY templates/ ./templates/
COPY static/ ./static/
COPY calc/ ./calc/
COPY data/ ./data/
ENV PORT=8080
EXPOSE 8080
CMD ["/docker-gs-ping"]