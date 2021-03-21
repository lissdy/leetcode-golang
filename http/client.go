package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	//resp, err := client.Get("")
	client.Get("")
}
