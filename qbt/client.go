package qbt

import (
	"bytes"
	"crypto/tls"
	"fmt"
	wrapper "github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// NewClient creates a connection to qbittorrent and performs requests
func NewClient(base string) *Client {
	c := &Client{}

	if strings.HasSuffix(base, "/") {
		c.URL = base[:len(base)-1]
	} else {
		c.URL = base
	}

	c.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	proxy, _ := url.Parse("http://localhost:8080")
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.http = &http.Client{
		Jar:       c.Jar,
		Transport: tr,
	}

	c.Authenticated = false
	return c
}

// Get will perform a GET request
func (client *Client) Get(endpoint string, opts map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(URLPattern, client.URL, endpoint),
		nil,
	)

	if err != nil {
		return nil, wrapper.Wrap(err, "failed to build request")
	}

	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "qBitTorrent-API "+Version)

	// add optional parameters that the user wants
	if opts != nil {
		query := req.URL.Query()
		for k, v := range opts {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, wrapper.Wrap(err, "failed to perform request")
	}

	return resp, nil
}

func (client *Client) GetResponseText(endpoint string, opts map[string]string) ([]byte, error) {
	resp, err := client.Get(endpoint, opts)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, ErrBadResponse
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post will perform a POST request with no content-type specified
func (client *Client) Post(endpoint string, opts map[string]string, headers map[string]string) (*http.Response, error) {
	// add optional parameters that the user wants
	form := url.Values{}
	if opts != nil {
		for k, v := range opts {
			form.Add(k, v)
		}
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(URLPattern, client.URL, endpoint),
		bytes.NewBuffer([]byte(form.Encode())),
	)

	if err != nil {
		return nil, wrapper.Wrap(err, "failed to build request")
	}

	// add the content-type so qbittorrent knows what to expect
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "qBitTorrent-API "+Version)

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, wrapper.Wrap(err, "failed to perform request")
	}

	return resp, nil
}
