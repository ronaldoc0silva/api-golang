FROM golang:latest as builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/fiber-mongo-api ./

EXPOSE 8080

CMD ["./fiber-mongo-api"]