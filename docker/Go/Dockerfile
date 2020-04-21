FROM golang:1.14.2

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /go/src

RUN apt-get update \
  && go get "github.com/go-sql-driver/mysql" \
  && go get "github.com/beego/bee" \
  && go get "github.com/astaxie/beego"

RUN bee new app
WORKDIR /go/src/app
CMD ["bee", "run"]