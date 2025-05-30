package core

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	MailTo     string
}

type Option func(*Client)

func PolitePool(email string) Option {
	return func(c *Client) {
		c.MailTo = email
	}
}

func New(opts ...Option) *Client {
	c := &Client{
		BaseURL:    "https://api.openalex.org",
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) Get(path string, out any) error {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}

	rel, err := url.Parse(path)
	if err != nil {
		return err
	}

	u := base.ResolveReference(rel)

	q := u.Query()
	if c.MailTo != "" {
		q.Set("mailto", c.MailTo)
	}
	u.RawQuery = q.Encode()

	resp, err := c.HTTPClient.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}
