package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
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
func (c *Client) Get(url string) ([]byte, error) {
	return c.Request(http.MethodGet, url, nil)
}
func (c *Client) PostJson(url string, params map[string]interface{}) ([]byte, error) {
	body, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
	}
	return c.Request(http.MethodPost, url, bytes.NewReader(body))
}
func (c *Client) PostRemoteFile() {

}
func (c *Client) PostFile() {

}
func (c *Client) Request(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if len(c.header) > 0 {
		for k, v := range c.header {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil

}
