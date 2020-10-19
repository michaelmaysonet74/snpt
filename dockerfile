FROM golang:1.15.3

WORKDIR /go/src/app
COPY ./cmd/snpt/main.go .

# RUN go get -d -v
RUN go install -v

# EXPOSE <port> 

CMD ["app"]
