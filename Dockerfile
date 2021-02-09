FROM golang:alpine AS builder

RUN apk --no-cache add make

WORKDIR /go/src/github.com/jesseobrien/jesseobrien.dev

COPY . .

RUN make install
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/github.com/jesseobrien/jesseobrien.dev/dist/jesseobrien .
CMD ["./jesseobrien", "--ip", "0.0.0.0"]
