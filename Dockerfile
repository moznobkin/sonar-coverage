FROM golang:alpine3.16 AS builder

RUN mkdir -p /go/src
COPY . /go/src
WORKDIR /go/src/


ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o target .



FROM golang:alpine3.16 AS runtime

COPY --from=builder /go/src/target /bin/


COPY ./.certs/certs.tar /tmp/certs.tar
RUN mkdir -p /usr/local/share/ca-certificates/ && \
    cd /usr/local/share/ca-certificates/ && \
    tar xf /tmp/certs.tar && \
    update-ca-certificates

EXPOSE 8080/tcp
ENTRYPOINT ["/bin/target"]
