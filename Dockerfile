FROM golang:latest
MAINTAINER Chenzhengwei 980752307@qq.com
RUN go get -u github.com/dubbogo/gost/bytes
RUN go get -u github.com/pkg/errors
COPY ./ $GOPATH/src/github.com/apache/dubbo-go-hessian2/
COPY ./run $GOPATH/src/github.com/apache/dubbo-go-hessian2/run
WORKDIR $GOPATH/src/github.com/apache/dubbo-go-hessian2/run
EXPOSE 2001
ENTRYPOINT ["go","run","Http.go"]
