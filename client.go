package TestRailGo

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type Client struct {
	url      string
	username string
	password string
}

func (c *Client) NewClient(inUrl, inUsername, inPassword string) {
	c.username = inUsername
	c.password = inPassword
	c.url = inUrl
	if !strings.HasSuffix(c.url, "/") {
		c.url += "/"
	}
	c.url += "index.php?/api/v2/"
}

func (c *Client) sendRequest(method, uri string, data interface{}) (interface{}, error) {
	req, err := http.NewRequest(method, c.url+uri, nil)
	if err != nil {
		return nil, err
	}

	if method == "POST" && data != nil {
		req.Body = "" // TODO make it work
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
	header = map[string][]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + auth,
	}

	req.Header = header

	resp, err := http.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
