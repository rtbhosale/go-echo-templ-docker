#For Production
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

COPY ./src .

RUN go build -C cmd -o /tmp/app

FROM alpine:3.16 AS production

COPY --from=builder /tmp/app .

CMD ./app

EXPOSE 1313