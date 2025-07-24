FROM golang:1.24

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /main ./cmd/server

CMD ["/main"]