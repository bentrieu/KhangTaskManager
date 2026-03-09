FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /taskManager ./cmd/

FROM alpine:3.20
WORKDIR /
COPY --from=builder /taskManager /taskManager
COPY ./changelog /changelog
EXPOSE 8080
CMD ["/taskManager"]
