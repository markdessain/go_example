FROM golang:1.9.0-alpine3.6

RUN apk add --no-cache git
RUN apk add --no-cache openssl
RUN apk add --no-cache protobuf

RUN go get gopkg.in/kniren/gota.v0/dataframe
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/lib/pq

WORKDIR /go
COPY . .

RUN protoc --go_out=plugins=grpc:src/ proto/*.proto

RUN go build -o app ./src

FROM alpine:latest
#scratch
WORKDIR /root/
COPY --from=0 /go .
CMD ["./app"]
