FROM golang:alpine AS build

WORKDIR /go/src/github.com/NoUseFreak/go-ddns/
RUN apk add --no-cache dep git

COPY ./Gopkg.* /go/src/github.com/NoUseFreak/go-ddns/
COPY ./cmd /go/src/github.com/NoUseFreak/go-ddns/cmd
COPY ./internal /go/src/github.com/NoUseFreak/go-ddns/internal

RUN cd /go/src/github.com/NoUseFreak/go-ddns/ \
    && dep ensure \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-ddns ./cmd/go-ddns

FROM alpine

RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=build /go/src/github.com/NoUseFreak/go-ddns/go-ddns .

ENTRYPOINT ["./go-ddns"]
CMD [ "/config.yml" ]
