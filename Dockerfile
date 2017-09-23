FROM golang:1.9.0-alpine3.6

RUN apk add --no-cache git
RUN apk add --no-cache openssl
RUN apk add --no-cache protobuf

RUN go get gopkg.in/kniren/gota.v0/dataframe
RUN go get github.com/golang/protobuf/protoc-gen-go

WORKDIR /go
COPY . .

RUN protoc --go_out=plugins=grpc:src/ proto/*.proto

# RUN go build -o ./account ./src/account
RUN go build -o ./web ./src/web
RUN go build -o ./event_log ./src/event_log

FROM alpine:latest
#scratch
WORKDIR /root/
COPY --from=0 /go .
# CMD ["./app"]
