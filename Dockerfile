FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ./web-service-echo

EXPOSE 1323

CMD ["./web-service-echo"]
