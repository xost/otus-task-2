FROM golang:1.21

ADD . /app

WORKDIR /app

RUN GOOS=linux GOARG=amd64 go build -o otus-task-2

CMD ["/app/otus-task-2"]
