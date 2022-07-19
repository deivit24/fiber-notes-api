FROM golang:1.18
WORKDIR /go/src/github.com/deivit24/fiber-notes-api
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]