FROM golang:1.8

RUN mkdir -p /go/src/github.com/nojnhuh/dotbook
WORKDIR /go/src/github.com/nojnhuh/dotbook
ADD . .
RUN go get -v
RUN go install -v ./...

CMD ["dotbook"]

EXPOSE 8080
