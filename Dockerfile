# Build stage
FROM golang:1.10 AS build-env

ARG RUN_TESTS

# Get dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/github.com/nojnhuh/dotbook
WORKDIR /go/src/github.com/nojnhuh/dotbook

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

COPY . .
RUN if [ ! -z "$RUN_TESTS" ]; then go test -v ./...; fi
RUN CGO_ENABLED=0 go build -o /dotbook


# Final stage
FROM alpine

EXPOSE 5050
COPY --from=build-env /dotbook /
ENTRYPOINT /dotbook
