package svixclient

import (
	"net/url"
	"os"

	svix "github.com/svix/svix-webhooks/go"
)

func New(u URL, token AuthToken) (*svix.Svix, error) {
	serverUrl, err := url.Parse(string(u))
	if err != nil {
		return nil, err
	}
	return svix.New(string(token), &svix.SvixOptions{
		ServerUrl: serverUrl,
		Debug:     true,
	}), nil
}

type URL string

var DefaultURL URL = ""

func ProvideURL() URL {
	if url := URL(os.Getenv("SVIX_URL")); url != "" {
		return url
	}
	return DefaultURL
}

type AuthToken string

var DefaultAuthToken AuthToken = ""

func ProvideAuthToken() AuthToken {
	if token := AuthToken(os.Getenv("SVIX_AUTH_TOKEN")); token != "" {
		return token
	}
	return DefaultAuthToken
}

var Client *svix.Svix

func Init() (err error) {
	if Client == nil {
		Client, err = New(
			ProvideURL(),
			ProvideAuthToken(),
		)
	}
	return err
}
