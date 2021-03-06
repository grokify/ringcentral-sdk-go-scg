package lite

import (
	"net/http"
	"strings"

	"github.com/grokify/simplego/net/httputilmore"
	"github.com/grokify/simplego/net/urlutil"
	"github.com/pkg/errors"
)

type ClientLite struct {
	ServerURL  string
	Token      string
	AccountID  string
	HTTPClient *http.Client
}

func NewClientLite(serverURL, accountID, token string) ClientLite {
	client := ClientLite{
		ServerURL: strings.TrimSpace(serverURL),
		AccountID: strings.TrimSpace(accountID),
		Token:     strings.TrimSpace(token)}
	client.LoadHTTPClient(client.Token)
	return client
}

func (lite *ClientLite) LoadHTTPClient(token string) {
	if len(strings.TrimSpace(token)) == 0 {
		token = lite.Token
	}
	lite.HTTPClient = NewHTTPClient(token)
}

func (lite *ClientLite) Tokens() ([]string, error) {
	return ListTokens(lite.ServerURL, lite.Token)
}

func NewHTTPClient(token string) *http.Client {
	header := http.Header{}
	header.Add(EngageVoiceTokenHeader, token)

	client := &http.Client{}
	client.Transport = httputilmore.TransportWithHeaders{
		Transport: client.Transport,
		Header:    header}
	return client
}

func APIInfo(serverURL, urlPath, authOrApiToken string) (http.Header, string, error) {
	authOrApiToken = strings.TrimSpace(authOrApiToken)
	if len(authOrApiToken) == 0 {
		return http.Header{}, "", errors.New("E_NO_TOKEN")
	}
	serverURL = strings.TrimSpace(serverURL)
	if len(serverURL) == 0 {
		serverURL = EngageVoiceServerURL
	}
	apiURL := urlutil.JoinAbsolute(serverURL, urlPath)

	return http.Header(map[string][]string{
		EngageVoiceTokenHeader: {authOrApiToken},
	}), apiURL, nil
}
