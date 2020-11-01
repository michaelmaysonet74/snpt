FROM golang:1.15.3

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY ./cmd/snpt/main.go .
COPY ./go.mod .

RUN go get -d -v
RUN go install -v

EXPOSE 9090

CMD ["app"]
