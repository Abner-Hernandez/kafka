FROM golang

WORKDIR /Server2

COPY . /Server2

RUN go get -u github.com/segmentio/kafka-go
RUN go get -u github.com/tidwall/sjson

CMD [ "go", "run", "Subscriber.go"]
