FROM golang
COPY . /go/src/app
WORKDIR /go/src/app
RUN go build wp.go
EXPOSE "80"
ENTRYPOINT ["/go/src/app/wp"]
