FROM golang

WORKDIR /Server1

COPY . /Server1

RUN go get -u github.com/segmentio/kafka-go
RUN go get -u github.com/tidwall/sjson

EXPOSE 80

CMD [ "go", "run", "Publisher.go"]
