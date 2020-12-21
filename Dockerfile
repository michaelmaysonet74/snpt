FROM golang:1.15.3

WORKDIR $GOPATH/src/github.com/mmaysonet74/snpt

# Manage dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy src code from the host and compile it
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /snpt cmd/snpt/main.go

EXPOSE 9090

CMD ["/snpt"]
