FROM golang:latest
RUN go version
COPY . "/go/src/github.com/chengguangyu/comodoca-helloworld"
RUN mkdir /helloworld
WORKDIR "/go/src/github.com/chengguangyu/comodoca-helloworld"
RUN go get -u -v github.com/kardianos/govendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /helloworld
CMD ["/helloworld"]
EXPOSE 8090