FROM golang

RUN go get gopkg.in/mgo.v2 \
    && git clone https://github.com/ryanhartje/gopot /tmp/gopot \
    && cd /tmp/gopot \
    && go build wp.go

EXPOSE "80"

ENTRYPOINT ["/tmp/gopot/wp"]
