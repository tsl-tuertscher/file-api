FROM golang:1.14
COPY . /usr/src/app/
WORKDIR /usr/src/app/

COPY vendor ./

RUN go mod vendor
RUN go test -v

COPY . .

RUN go build -o main

CMD ["./main"]
