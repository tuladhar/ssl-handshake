# ssl-handshake
A command-line tool for testing SSL/TLS handshake latency, written in [Go](https://go.dev/).

## Features
* TCP handshake latency
* SSL/TLS handshake latency
* TLS version used during the handshake
* Display handshake statistics
* Configurable endpoint port, handshake interval, timeout and count

## 
<p align="center">
  <img width="600" src="https://github.com/tuladhar/ssl-handshake/blob/main/demo/ssl-handshake-1.6.2.svg">
</p>


## What is an SSL/TLS Handshake?
An SSL/TLS handshake is the process that kicks off a communication session between client and server that uses [TLS encryption](https://en.wikipedia.org/wiki/Transport_Layer_Security). During a TLS handshake, the two communicating sides exchange messages to acknowledge each other, verify each other, establish the encryption algorithms they will use, and agree on session keys. TLS handshakes are a foundational part of how HTTPS works and it is defined in [RFC 8446 (for TLS 1.3)](https://tools.ietf.org/html/rfc8446) or in [RFC 5246 (for TLS 1.2)](https://datatracker.ietf.org/doc/html/rfc5246).

TLS handshakes occur after a TCP connection has been opened via a TCP handshake. 

![image](https://user-images.githubusercontent.com/5674762/147405109-319358ff-8506-4383-8778-6961e99c4e99.png)

TLS handshake packets captured with [Wireshark](https://www.wireshark.org/).

![image](https://user-images.githubusercontent.com/5674762/147404043-7e6d983a-e9c5-4477-a8e2-3e054d4f861d.png)

## Docker Image

Docker image is publicly available at DockerHub:
* https://hub.docker.com/r/ptuladhar/ssl-handshake

Run `ssl-handshake` as Docker container:
```
docker run --rm ptuladhar/ssl-handshake -c 5 tuladhar.github.io:443
```
You can also alias `ssl-handshake`, for ease of use:
```
alias ssl-handshake="docker run --rm ptuladhar/ssl-handshake"
ssl-handshake tuladhar.github.com:443
```
## Install binary
Binary is available for Linux, Windows and Mac OS (amd64 and arm64). Download the binary for your respective platform from the [releases page](https://github.com/tuladhar/ssl-handshake/releases).

Linux:
```
curl -sSLO https://github.com/tuladhar/ssl-handshake/releases/download/v1.6.1/ssl-handshake-v1.6.1-linux-amd64.tar.gz
```
```
tar zxf ssl-handshake-v1.6.1-linux-amd64.tar.gz
```
```
sudo install -m 0755 ssl-handshake /usr/local/bin/ssl-handshake
```

macOS (Intel):
```
curl -sSLO https://github.com/tuladhar/ssl-handshake/releases/download/v1.6.1/ssl-handshake-v1.6.1-darwin-amd64.tar.gz
```
```
tar zxf ssl-handshake-v1.6.1-darwin-amd64.tar.gz
```
```
sudo install -m 0755 ssl-handshake /usr/local/bin/ssl-handshake
```

macOS (Apple Silicon):
```
curl -sSLO https://github.com/tuladhar/ssl-handshake/releases/download/v1.6.1/ssl-handshake-v1.6.1-darwin-arm64.tar.gz
```
```
tar zxf ssl-handshake-v1.6.1-darwin-arm64.tar.gz
```
```
sudo install -m 0755 ssl-handshake /usr/local/bin/ssl-handshake
```

Windows:
```
curl -sSLO https://github.com/tuladhar/ssl-handshake/releases/download/v1.6.1/ssl-handshake-v1.6.1-windows-amd64.zip
```
```
unzip ssl-handshake-v1.6.1-windows-amd64.zip
```

## Development

If you wish to contribute or compile from source code, you'll first need Go installed on your machine. Go version 1.17+ is required. Currently, there are no dependencies on third-party modules. 

```
git clone https://github.com/tuladhar/ssl-handshake
cd ssl-handshake 
go build
```

## Contributors
- [Puru Tuladhar](https://tuladhar.github.io)
