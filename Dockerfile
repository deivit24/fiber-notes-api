FROM golang:1.18
WORKDIR /go/src/github.com/deivit24/api-db
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]