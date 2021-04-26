FROM golang

WORKDIR /Server2

COPY . /Server2

RUN go get -u github.com/segmentio/kafka-go
RUN go get -u github.com/tidwall/sjson

RUN export GOOGLE_APPLICATION_CREDENTIALS='clave.json'

CMD [ "go", "run", "async_pull.go"]
