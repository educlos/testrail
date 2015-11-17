package testrail

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	url        string
	username   string
	password   string
	httpClient *http.Client
}

func (c *Client) NewClient(inUrl, inUsername, inPassword string) {
	c.username = inUsername
	c.password = inPassword
	c.url = inUrl
	if !strings.HasSuffix(c.url, "/") {
		c.url += "/"
	}
	c.url += "index.php?/api/v2/"

	c.httpClient = &http.Client{}
}

func (c *Client) sendRequest(method, uri string, data, v interface{}) error {
	var body io.Reader
	if data != nil {
		jsonReq, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("marshaling data: %s", err)
		}

		body = bytes.NewBuffer(jsonReq)
	}

	req, err := http.NewRequest(method, c.url+uri, body)
	if err != nil {
		return err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	jsonCnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if v != nil {
		err = json.Unmarshal(jsonCnt, v)
		if err != nil {
			return fmt.Errorf("unmarshaling data: %s", err)
		}
	}

	return nil
}
