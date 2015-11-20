package testrail

import (
	"bytes"
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

func (c *Client) NewClient(url, username, password string) {
	c.username = username
	c.password = password

	c.url = url
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

	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("response: status=%q", resp.Status)
	}

	jsonCnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading: %s", err)
	}

	if v != nil {
		err = json.Unmarshal(jsonCnt, v)
		if err != nil {
			return fmt.Errorf("unmarshaling response: %s", err)
		}
	}

	return nil
}
