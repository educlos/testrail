// testrail provides a go api for testrail
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

// A Client stores the client informations
// and implement all the api functions
// to communicate with testrail
type Client struct {
	url        string
	username   string
	password   string
	HTTPClient *http.Client
}

// NewClient returns a new client
// with the given credential
// for the given testrail domain
func NewClient(url, username, password string) (c *Client) {
	c = &Client{}
	c.username = username
	c.password = password

	c.url = url
	if !strings.HasSuffix(c.url, "/") {
		c.url += "/"
	}
	c.url += "index.php?/api/v2/"

	c.HTTPClient = http.DefaultClient

	return
}

// sendRequest sends a request of type "method"
// to the url "client.url+uri" and with optional data "data"
// Returns an error if any and the optional data "v"
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

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
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
