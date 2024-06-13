FROM golang:1.22.3-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache git \
    && go mod download \
    && go build -o main .

EXPOSE 8080

CMD ["./main"]
