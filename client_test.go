package testrail

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// NewTestClient returns a mocked http.Client
func NewTestClient(replyResp *http.Response, err error) *http.Client {
	client := &http.Client{}
	client.Transport = &MockTransport{
		resp: replyResp,
		err:  err,
	}
	return client
}

type MockTransport struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (b *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	b.req = req
	return b.resp, b.err
}

func newResponse(body string) *http.Response {
	return &http.Response{Body: ioutil.NopCloser(bytes.NewBuffer([]byte(body)))}
}

// TestSendRequest tests all the client functionalities
func TestSendRequest(t *testing.T) {
	testClient(t)

	c := NewCustomClient("http://example.com", "testUsername", "testPassword", NewTestClient(newResponse(`{ "status_id": 1 }`), nil))

	testValidGetRequest(t, c)
	testInvalidGetRequest(t, c)
	testValidPostRequest(t, c)
}

// testClient tests the NewClient method
func testClient(t *testing.T) {
	c1 := NewClient("http://example.com", "testUsername", "testPassword")

	if c1.url != "http://example.com/index.php?/api/v2/" {
		t.Fatal("Expected valid url but got ", c1.url)
	}

	c2 := NewClient("http://example.com/", "testUsername", "testPassword")

	if c2.url != "http://example.com/index.php?/api/v2/" {
		t.Fatal("Expected valid url but got ", c2.url)
	}
}

// testValidGetRequest tests the sendRequest method for a GET
func testValidGetRequest(t *testing.T, c *Client) {
	var v struct {
		StatusID int `json:"status_id"`
	}

	err := c.sendRequest("GET", "test", nil, &v)

	if err != nil {
		t.Fatal("Expected no error but got ", err)
	}

	if v.StatusID != 1 {
		t.Fatal("Expected StatusID to be 1, was ", v.StatusID)
	}
}

func testValidPostRequest(t *testing.T, c *Client) {
	var v struct {
		Title string `json:"title"`
	}

	v.Title = "test"
	err := c.sendRequest("POST", "test", v, nil)

	if err != nil {
		t.Fatal("Expected no error but got ", err)
	}
}

func testInvalidGetRequest(t *testing.T, c *Client) {
	var v struct {
		Status int `json:"status"`
	}

	err := c.sendRequest("GET", "test", nil, &v)

	if err == nil {
		t.Fatal("Expected error but got none")
	}
}
