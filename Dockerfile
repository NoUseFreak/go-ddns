FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . /app
RUN go mod vendor \
    && CGO_ENABLED=0 go build -ldflags='-w -s -extldflags "-static"' -a ./cmd/go-ddns

FROM gcr.io/distroless/static

COPY --from=build /app/go-ddns /go-ddns
USER 65532
ENTRYPOINT [ "/go-ddns" ]
