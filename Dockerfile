FROM golang:1.9

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go get -v github.com/linyuy/editor

WORKDIR $GOPATH/src/github.com/linyuy/editor

RUN go build

CMD editor -p 80
