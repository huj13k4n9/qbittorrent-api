package qbt

import (
	"bytes"
	"fmt"
	wrapper "github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

// NewClient creates a new client and is used to perform future requests.
func NewClient(base string) *Client {
	c := &Client{}

	if strings.HasSuffix(base, "/") {
		c.URL = base[:len(base)-1]
	} else {
		c.URL = base
	}

	c.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	c.http = &http.Client{
		Jar: c.Jar,
	}

	c.Authenticated = false
	return c
}

// Get will perform a GET request, with parameters.
func (client *Client) Get(endpoint string, opts map[string]string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(URLPattern, client.URL, endpoint),
		nil,
	)

	if err != nil {
		return nil, wrapper.Wrap(err, "failed to build request")
	}

	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "qBittorrent-API "+Version)

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

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

// GetResponseBody will perform a GET request with parameters,
// and directly returns the body of response.
func (client *Client) GetResponseBody(endpoint string, opts map[string]string, headers map[string]string) ([]byte, error) {
	resp, err := client.Get(endpoint, opts, headers)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadResponse
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post will perform a POST request with application/x-www-form-urlencoded parameters
// and custom HTTP headers.
func (client *Client) Post(endpoint string, opts any, headers map[string]string, contentType string) (*http.Response, error) {
	var postData *bytes.Buffer
	typeAsserted := false
	if opts != nil {
		if params, ok := opts.(map[string]string); ok {
			form := url.Values{}
			for k, v := range params {
				form.Add(k, v)
			}
			postData = bytes.NewBufferString(form.Encode())
			typeAsserted = true
		}
		if params, ok := opts.(*bytes.Buffer); ok {
			postData = params
			typeAsserted = true
		}

		if !typeAsserted {
			return nil, wrapper.Wrap(ErrUnknownType, "post data type unknown")
		}
	} else {
		postData = nil
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(URLPattern, client.URL, endpoint),
		postData,
	)

	if err != nil {
		return nil, wrapper.Wrap(err, "failed to build request")
	}

	// add the content-type so qbittorrent knows what to expect
	req.Header.Set("Content-Type", contentType)
	// add user-agent header to allow qbittorrent to identify us
	req.Header.Set("User-Agent", "qBittorrent-API "+Version)

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

func (client *Client) PostWithParams(endpoint string, opts map[string]string, headers map[string]string) (*http.Response, error) {
	return client.Post(endpoint, opts, headers, "application/x-www-form-urlencoded")
}

func (client *Client) PostMultipart(endpoint string, data *bytes.Buffer, contentType string) (*http.Response, error) {
	return client.Post(endpoint, data, nil, contentType)
}

func (client *Client) RequestAndHandleError(
	method string,
	endpoint string,
	opts map[string]string,
	headers map[string]string,
	errorMsgs map[string]string,
) (*http.Response, error) {
	if !client.Authenticated {
		return nil, ErrUnauthenticated
	}

	var resp *http.Response
	var err error

	switch method {
	case "GET":
		resp, err = client.Get(endpoint, opts, headers)
		break
	case "POST":
		resp, err = client.PostWithParams(endpoint, opts, headers)
		break
	default:
		return nil, wrapper.Wrap(ErrBadResponse, "Unknown method "+method)
	}

	if err != nil {
		return nil, err
	}

	for k, v := range errorMsgs {
		var code int
		if k[0] == '!' {
			code, _ = strconv.Atoi(k[1:])
			if resp.StatusCode != code {
				return nil, wrapper.Wrap(ErrBadResponse, v)
			}
		} else {
			code, _ = strconv.Atoi(k)
			if resp.StatusCode == code {
				return nil, wrapper.Wrap(ErrBadResponse, v)
			}
		}
	}
	return resp, nil
}
