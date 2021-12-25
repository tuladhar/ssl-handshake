# ssl-handshake
A command-line tool for testing SSL handshake latency, written in [Go](https://go.dev/).

## What is SSL Handshake?
> SSL handshake enables client and server to establish the secret keys with which they communicate securely.

## Installation
Binary is available for Linux, Windows and Mac OS (amd64 and arm64). Download the binary for your respective platform from the [releases page](https://github.com/tuladhar/ssl-handshake/releases).

Linux:
```
wget https://github.com/tuladhar/ssl-handshake/releases/download/v1.5/ssl-handshake-v1.5-linux-amd64.tar.gz
```
```
tar zxf ssl-handshake-v1.5-linux-amd64.tar.gz
```
```
sudo install -m 0755 ssl-handshake /usr/local/bin/ssl-handshake
```

macOS (Intel):
```
wget https://github.com/tuladhar/ssl-handshake/releases/download/v1.5/ssl-handshake-v1.5-darwin-amd64.tar.gz
```
```
tar zxf ssl-handshake-v1.5-darwin-amd64.tar.gz
```
```
sudo install -m 0755 ssl-handshake /usr/local/bin/ssl-handshake
```

macOS (Apple Silicon):
```
wget https://github.com/tuladhar/ssl-handshake/releases/download/v1.5/ssl-handshake-v1.5-darwin-arm64.tar.gz
```
```
tar zxf ssl-handshake-v1.5-darwin-arm64.tar.gz
```
```
sudo install -m 0755 ssl-handshake /usr/local/bin/ssl-handshake
```

## Demo
<p align="center">
  <img width="600" src="https://github.com/tuladhar/ssl-handshake/blob/main/demo/ssl-handshake.svg">
</p>

## Development

If you wish to contribute or compile from source code, you'll first need Go installed on your machine. Go version 1.17+ is required. Currently, there's no dependencies on third-party modules. 

```
git clone https://github.com/tuladhar/ssl-handshake
cd ssl-handshake 
go build
```

## Contributors
- [Puru Tuladhar](https://tuladhar.github.io)
