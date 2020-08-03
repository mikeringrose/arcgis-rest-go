package arcgis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type api struct {
	client *Client
}

// Client TODO
type Client struct {
	baseURL    string
	httpClient *http.Client
	Portal     PortalAPI
	Community  CommunityAPI
}

// Context TODO
type Context struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient todo
func NewClient(httpClient *http.Client, portalURL string) *Client {
	client := &Client{
		baseURL:    portalURL,
		httpClient: httpClient,
	}

	client.Portal = PortalAPI{client}
	client.Community = CommunityAPI{client}

	return client
}

// NewRequest TODO
func (c *Client) NewRequest(method string, path string, v interface{}) (*http.Request, error) {
	var bs []byte

	if v != nil {
		b, err := json.Marshal(v)

		if err != nil {
			return nil, err
		}

		bs = b
	}

	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewReader(bs))

	if err != nil {
		return nil, err
	}

	return req, err
}

// Do TODO
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != err {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		errUnmarshall := json.Unmarshal(body, v)

		if errUnmarshall != nil {
			return nil, errUnmarshall
		}
	}

	return resp, err
}
