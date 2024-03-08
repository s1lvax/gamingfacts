From golang:1.22

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]