# ssl-handshake
A command-line tool for testing SSL handshake latency, written in [Go](https://go.dev/).

## Installation
Binary is available for Linux, Windows and Mac OS (amd64 and arm64). Download the binary for your respective platform from the [releases page](https://github.com/tuladhar/ssl-handshake/releases).

Linux (amd64):
```
wget https://github.com/tuladhar/ssl-handshake/releases/download/v1.3/ssl-handshake-v1.3-linux-amd64.tar.gz
tar zxf ssl-handshake-v1.3-linux-amd64.tar.gz
./ssl-handshake
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
