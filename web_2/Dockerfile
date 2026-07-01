FROM golang:1.24

WORKDIR /app


COPY . .

EXPOSE 8080

RUN go build -o server .

CMD ["./server"]