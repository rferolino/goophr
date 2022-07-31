FROM golang:1.9.1 

ADD . /go/src/github.com/rferolino/goophr/librarian 

WORKDIR /go/src/github.com/rferolino/goophr/librarian 

RUN go install github.com/rferolino/goophr/librarian 

ENTRYPOINT /go/bin/librarian 

EXPOSE 9000 