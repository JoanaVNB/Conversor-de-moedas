FROM golang:latest

WORKDIR /exchangeapp
ADD . .
COPY . /exchangeapp
EXPOSE 8000

RUN go build main.go
ENTRYPOINT ["./main"]

