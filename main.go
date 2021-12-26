package main

import (
	sslhandshake "github.com/tuladhar/ssl-handshake/internal"
)

func main() {
	sslhandshake := sslhandshake.SSLHandshake{
		Metadata: &sslhandshake.SSLHandshakeMetadata{
			Name:        "ssl-handshake",
			Description: "A command-line tool for testing SSL/TLS handshake latency.",
			Version:     "1.5.2",
			URL:         "https://tuladhar.github.com/ssl-handshake",
		},
		Config: &sslhandshake.SSLHandshakeConfig{
			Endpoint:  "",
			Interval:  1000,
			Timeout:   5000,
			StopCount: 0,
		},
		State: &sslhandshake.SSLHandshakeState{},
	}
	sslhandshake.InitCLI()
	sslhandshake.NotifySignal()
	sslhandshake.Start()
}
