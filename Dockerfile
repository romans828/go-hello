FROM golang:1.10.1
MAINTAINER romans828 <romans0828@gmail.com>

ADD ./ .
RUN go build main.go

ENTRYPOINT ["./main"]
