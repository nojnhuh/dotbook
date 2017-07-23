# Build stage
FROM golang:1.8-alpine AS build-env

RUN apk add --no-cache git
RUN mkdir -p /go/src/github.com/nojnhuh/dotbook
WORKDIR /go/src/github.com/nojnhuh/dotbook
EXPOSE 8080
ADD . .
RUN go get -v && go install -v ./...
ENTRYPOINT dotbook



# # Final stage
# FROM alpine

# COPY --from=build-env /go/bin/dotbook /
# EXPOSE 8080
# ENTRYPOINT /dotbook
