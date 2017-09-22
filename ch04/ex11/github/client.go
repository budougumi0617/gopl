// Copyright 2017 budougumi0617 All Rights Reserved.

package github

import (
	"net/http"
	"net/url"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Username, Password string
}

func NewClient(url, username, password string) (*Client, err) {
	return nil, _
}
