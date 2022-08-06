
FROM golang:latest 

WORKDIR /go/src/bobble
COPY . .
RUN go get -u github.com/lib/pq

RUN go build -o main .

CMD ["./main"]