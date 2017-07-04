# Build stage
FROM golang:1.8-alpine AS build-env

RUN mkdir -p /go/src/github.com/nojnhuh/dotbook
WORKDIR /go/src/github.com/nojnhuh/dotbook
ADD . .
RUN apk add --no-cache git && go get -v
RUN go install -v ./...

# Final stage
FROM alpine

COPY --from=build-env /go/bin/dotbook /
EXPOSE 8080
ENTRYPOINT /dotbook
