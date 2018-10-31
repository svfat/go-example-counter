FROM golang:alpine
ADD . /go/src/github.com/svfat/go-example-counter
RUN go install github.com/svfat/go-example-counter
ENTRYPOINT /go/bin/go-example-counter
EXPOSE 8080
