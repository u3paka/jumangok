package jmg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
}

const version = "0.01"

var userAgent = fmt.Sprintf("XXXGoClient/%s (%s)", version, runtime.Version())

func NewClient(urlStr string) *Client {
	return &Client{
		URL: &url.URL{
			Scheme: "http",
			Path:   urlStr,
		},
		HTTPClient: http.DefaultClient,
	}
}

// func (c *Client) JumanppRaw(ctx context.Context, input string) (ret string, err error) {
// 	const spath = ""
// 	req, err := c.newRequest(ctx, "POST", spath, strings.NewReader(input))
// 	if err != nil {
// 		return
// 	}

// 	res, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return
// 	}

// 	res.Body
// 	return
// }

func (c *Client) Jumanpp(ctx context.Context, input string) (ws []*Word, err error) {
	const spath = "json"
	req, err := c.newRequest(ctx, "POST", spath, strings.NewReader(input))
	if err != nil {
		return
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&ws)
	return
}

// 共通部 common
func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	// req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}
