FROM golang:1.14.2

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /go/src

RUN apt-get update \
  && go get -u "github.com/labstack/echo/..."

RUN mkdir -p /home/root/go
RUN cp -R /go/src /home/root/go

RUN mkdir -p /go/app
WORKDIR /go/app
CMD ["go", "run", "main.go"]
