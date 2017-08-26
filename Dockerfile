# Build stage
FROM golang:1.8-alpine AS build-env

RUN apk add --no-cache git
RUN mkdir -p /go/src/github.com/nojnhuh/dotbook
WORKDIR /go/src/github.com/nojnhuh/dotbook
COPY . .
RUN go get -v
RUN go install ./...
ENTRYPOINT dotbook



# Final stage
FROM alpine

EXPOSE 443
COPY --from=build-env /go/src/github.com/nojnhuh/dotbook/server.* /
COPY --from=build-env /go/bin/dotbook /
ENTRYPOINT /dotbook
