FROM golang:1.9.0-alpine3.6

RUN apk add --no-cache git

RUN go get -d -v github.com/kniren/gota/dataframe

WORKDIR /go
COPY ./src .

RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go .
CMD ["./app"]
