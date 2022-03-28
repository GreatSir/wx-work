package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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
		return nil, err
	}
	return c.Request(http.MethodPost, url, bytes.NewReader(body))
}
func (c *Client) PostRemoteFile() {

}
func (c *Client) PostFile(fieldname, filename, url string, params map[string]string) ([]byte, error) {
	pr, pw := io.Pipe()
	defer pr.Close()
	defer pw.Close()
	bodyWriter := multipart.NewWriter(pw)
	defer bodyWriter.Close()
	fileWriter, err := bodyWriter.CreateFormFile(fieldname, filename)
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	for k, v := range params {
		bodyWriter.WriteField(k, v)
	}
	c.SetHeader("Content-Type", bodyWriter.FormDataContentType())
	c.SetHeader("Transfer-Encoding", "chunked")
	body := io.NopCloser(pr)
	return c.Request(http.MethodPost, url, body)
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
