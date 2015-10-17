package client

import (
	"github.com/andygrunwald/go-gerrit"
	"net/http"
	"time"
)

type Client struct {
	Gerrit *gerrit.Client
}

const (
	DefaultChangeSetQueryLimit = 500
	AuthModeBasic              = "basic"
	AuthModeCookie             = "cookie"
)

func NewGerritClient(instance string, seconds int) (*Client, error) {
	timeout := time.Duration(time.Duration(seconds) * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
	}

	g, err := gerrit.NewClient(instance, httpClient)
	c := &Client{
		Gerrit: g,
	}

	return c, err
}

func (c *Client) Authentication(m, u, p string) {
	switch m {
	case AuthModeBasic:
		c.Gerrit.Authentication.SetBasicAuth(u, p)
	case AuthModeCookie:
		c.Gerrit.Authentication.SetCookieAuth(u, p)
	default:
	}
}

func (c *Client) GetQueryLimit() int {
	if !c.Gerrit.Authentication.HasAuth() {
		return DefaultChangeSetQueryLimit
	}

	opt := &gerrit.CapabilityOptions{
		Filter: []string{"queryLimit"},
	}
	capability, _, err := c.Gerrit.Accounts.ListAccountCapabilities("self", opt)
	if err != nil {
		return DefaultChangeSetQueryLimit
	}

	return capability.QueryLimit.Max
}
