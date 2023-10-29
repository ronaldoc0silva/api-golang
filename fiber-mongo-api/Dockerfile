FROM golang:alpine

WORKDIR /app
COPY . /app

RUN go install 
CMD ["/go/bin/fiber-mongo-api"]
EXPOSE 8080