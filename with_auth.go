package animelayer

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// Credentials - Animelayer data for loging request
type Credentials struct {
	Login    string
	Password string
}

// DefaultClientWithAuth - http.Client with auth to AnimeLayer
func DefaultClientWithAuth(cred Credentials) (*http.Client, error) {
	jarc, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jarc,
	}

	hostURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	loginURL, _ := hostURL.Parse("/auth/login/")
	login := loginURL.String()
	_, err = client.PostForm(
		login,
		url.Values{
			"login":    {cred.Login},
			"password": {cred.Password},
		},
	)

	if len(client.Jar.Cookies(hostURL)) < 3 {
		return nil, fmt.Errorf("error on login to '%s'", hostURL.Host)
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
