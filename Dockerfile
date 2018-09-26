FROM golang:latest
RUN mkdir /app
EXPOSE 8080
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]