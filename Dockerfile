FROM golang:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app
#Add go mod tidy command before
RUN go mod tidy
RUN go build -o main .
EXPOSE 9001
CMD ["/app/main"]