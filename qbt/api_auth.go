package qbt

import (
	"github.com/huj13k4n9/qbittorrent-api/consts"
	wrapper "github.com/pkg/errors"
	"net/http"
	"net/url"
)

// Login perform login request to server
func (client *Client) Login(username, password string) (success bool, err error) {
	resp, err := client.Post(
		consts.LoginEndpoint,
		map[string]string{"username": username, "password": password},
		map[string]string{
			"Origin": client.URL, "Referer": client.URL,
		},
	)
	if err != nil {
		return false, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if cookies := resp.Cookies(); len(cookies) > 0 {
			cookieURL, _ := url.Parse(client.URL)
			client.Jar.SetCookies(cookieURL, cookies)
		} else {
			return false, wrapper.Wrap(ErrBadResponse, "login failed: no cookie returned")
		}
		client.http.Jar = client.Jar
		client.Authenticated = true
		return true, nil
	case http.StatusForbidden:
		return false, wrapper.Wrap(ErrBadResponse, "user's IP is banned for too many failed login attempts")
	default:
		return false, wrapper.Wrap(ErrBadResponse, "login failed")
	}
}

// Logout perform logout request to server
func (client *Client) Logout() error {
	_, err := client.RequestAndHandleError(
		"POST", consts.LogoutEndpoint, nil, nil,
		map[string]string{"!200": "logout failed"})

	if err != nil {
		return err
	}

	client.Authenticated = false
	return nil
}
