FROM golang

COPY wp.go /tmp/

RUN go get gopkg.in/mgo.v2 \
    && go build /tmp/wp.go

EXPOSE "80"

ENTRYPOINT ["/tmp/wp"]
