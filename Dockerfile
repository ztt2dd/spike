FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/ztt2dd/spike
COPY . $GOPATH/src/github.com/ztt2dd/spike
RUN go build .

EXPOSE 8888
ENTRYPOINT ["./spike"]
