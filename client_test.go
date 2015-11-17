package testrail

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

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

func TestSendRequest(t *testing.T) {
	c := &Client{}
	c.NewClient("http://example.com", "testUsername", "testPassword")
	c.httpClient = NewTestClient(newResponse(`{ "status_id": 1 }`), nil)

	var v struct {
		StatusID int `json:"status_id"`
	}
	err := c.sendRequest("GET", "test", nil, &v)

	if err != nil {
		t.Fatal("Expected no error but got ", err)
	}

	if v.StatusID != 1 {
		t.Fatal("Expected StatusId to be 1, was ", v.StatusID)
	}
}
