package parser

import (
	"net/http"
	"net/url"
	"torParser/client"
	"torParser/settings"
)

type Settings struct {
	Domain    string
	TagName   string
	Attribute string
}

func SetSettings(d, t, a string) *Settings {
	return &Settings{
		Domain:    d,
		TagName:   t,
		Attribute: a,
	}
}

func (s *Settings) ParseUrl() (htmlPage string) {
	c := client.CreateClient()

	urlString := "https://" + s.Domain

	u, err := url.Parse(urlString)
	if err != nil {
		settings.ErrorExit(err)
	}

	req := &http.Request{
		Method: "GET",
		URL:    u,
		Header: http.Header{},
	}

	req.Header.Set("Accept", "text/html")
	req.Header.Set("User-Agent", settings.RandAgent())

	resp, err := c.Do(req)
	if err != nil {
		settings.ErrorExit(err)
	}
	defer resp.Body.Close()

	var buff = make([]byte, settings.BUFFER)

	for {
		lenth, err := resp.Body.Read(buff)
		htmlPage += string(buff[:lenth])
		if err != nil || lenth == 0 {
			break
		}
	}

	return htmlPage
}
