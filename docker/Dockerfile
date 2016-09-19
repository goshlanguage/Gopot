FROM golang

RUN go get gopkg.in/mgo.v2 \
    && git clone https://github.com/ryanhartje/gopot /tmp/gopot

EXPOSE "80"

ENTRYPOINT ["/tmp/gopot/wp"]
