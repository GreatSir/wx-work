package httpclient

import (
	"log"
	"net/http"
)

type Client struct {
	http.Client
	header map[string]string
}

func (c *Client) SetHeader(k, v string) *Client {
	c.header[k] = v
	return c
}
func (c *Client) Get() {

}
func (c *Client) PostJson() {

}
func (c *Client) PostRemoteFile() {

}
func (c *Client) PostFile() {

}
func (c *Client) Request(method, url string, body interface{}) {
	var req *http.Request
	var err error
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(c.header) > 0 {
		for k, v := range c.header {
			req.Header.Set(k, v)
		}
	}
	if method == http.MethodPost {

	}
	c.Do(req)
}
