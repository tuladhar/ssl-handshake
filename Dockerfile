FROM golang:1.17.5-alpine as builder
LABEL maintainer "Puru Tuladhar <ptuladhar3@gmail.com>"
WORKDIR /go/src/github.com/tuladhar/ssl-handshake
ENV GOPATH=/go
COPY . .
RUN go mod download && go install

FROM alpine
COPY --from=builder /go/bin/ssl-handshake /bin/ssl-handshake
USER 1001
ENTRYPOINT ["/bin/ssl-handshake"]
