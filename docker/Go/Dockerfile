FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]