#  Build Image
FROM golang:1.13 AS builder
WORKDIR /go/src/github.com/app_greet
RUN go get -d -v github.com/urfave/cli
COPY main.go .
RUN GOOS=linux go build -a -o greet .

# ------------------------------

#  Production Image
FROM busybox
WORKDIR /opt/greet/bin
COPY --from=builder /go/src/github.com/app_greet/ .
ENTRYPOINT ["./greet"]

