package client

import (
	"context"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"time"
	"torParser/settings"
)

func CreateClient() *http.Client {
	var err error

	torUrl, err := url.Parse(settings.TorProxy)
	settings.ErrorExit(err)

	torDialer, err := proxy.FromURL(torUrl, proxy.Direct)
	settings.ErrorExit(err)

	torTransport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return torDialer.Dial(network, addr)
		},
	}

	return &http.Client{
		Transport: torTransport,
		Timeout:   7 * time.Second,
	}
}
