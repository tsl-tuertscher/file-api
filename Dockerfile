FROM golang:1.14

WORKDIR /usr/src/app

COPY go.mod go.sum ./
COPY vendor ./

RUN go mod vendor
RUN go test -v

COPY . .

RUN go build -o main

CMD ["./main"]
