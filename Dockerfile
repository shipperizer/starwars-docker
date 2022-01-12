# BUILDER
FROM --platform=$BUILDPLATFORM golang:1.16.5 AS builder

ARG SKAFFOLD_GO_GCFLAGS
ARG TARGETOS
ARG TARGETARCH

ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN go env

WORKDIR /var/app

COPY . .

RUN go build -a -installsuffix nocgo -o /go/bin/app main.go

LABEL org.opencontainers.image.source=https://github.com/shipperizer/starwars-docker

# RUNTIME
FROM gcr.io/distroless/static:nonroot

LABEL org.opencontainers.image.source=https://github.com/shipperizer/starwars-docker
LABEL maintainer "Thomas Graf <tgraf@tgraf.ch>"

COPY --from=builder /go/bin/app /

EXPOSE 80

ENTRYPOINT ["/app", "--port", "80", "--host", "0.0.0.0"]
