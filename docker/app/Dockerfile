FROM golang:latest

RUN mkdir -p /go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

ADD . /go/src/github.com/josephniel/go-api
WORKDIR /go/src/github.com/josephniel/go-api

RUN go get -u github.com/golang/dep/cmd/dep
RUN cd /go/src/github.com/golang/dep/cmd/dep && go build
RUN dep init

RUN go build -v app/main.go

CMD ./main -port=${PORT} -env=${ENV}
