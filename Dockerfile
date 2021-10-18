FROM golang:1.17-alpine AS build

WORKDIR /app
COPY . /app
RUN go mod vendor \
    && CGO_ENABLED=0 go build -ldflags='-w -s -extldflags "-static"' -a ./cmd/go-ddns

FROM gcr.io/distroless/static

USER nonroot
COPY --from=build /app/go-ddns /go-ddns
ENTRYPOINT [ "/go-ddns" ]
