FROM golang:1.18.2-alpine as builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /main .

EXPOSE $BACKEND_PORT

CMD ["/main"]