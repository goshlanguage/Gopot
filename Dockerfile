FROM golang

COPY wp.go /go/src/app/wp.go
WORKDIR /go/src/app

RUN go build wp.go

EXPOSE "80"

ENTRYPOINT ["/go/src/app/wp"]
