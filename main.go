package main

import (
	sslhandshake "github.com/tuladhar/ssl-handshake/internal"
)

func main() {
	sslhandshake := sslhandshake.SSLHandshake{
		Metadata: &sslhandshake.SSLHandshakeMetadata{
			Name:        "ssl-handshake",
			Description: "A debugging tool to check the latency of SSL handshake.",
			Version:     "1.0",
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
