package elasticemail

import (
	"net/http"
)

const (
	ApiBase = "https://api.elasticemail.com/v2"
)

type ElasticEmail interface {
	APIBase() string
	APIKey() string
	Client() *http.Client
	SetClient(client *http.Client)
	GetEmailStatusParams(params GetEmailStatusParams)
}

type ElasticEmailImpl struct {
	apiBase string
	apiKey  string
	client  *http.Client
}

func NewElasticEmail(apiKey string) *ElasticEmailImpl {
	return &ElasticEmailImpl{
		apiBase: ApiBase,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

// Client returns the HTTP client configured for this client.
func (m *ElasticEmailImpl) Client() *http.Client {
	return m.client
}

// SetClient updates the HTTP client for this client.
func (m *ElasticEmailImpl) SetClient(c *http.Client) {
	m.client = c
}

// ApiBase returns the API Base URL configured for this client.
func (m *ElasticEmailImpl) APIBase() string {
	return m.apiBase
}

// SetAPIBase updates the API Base URL for this client.
func (m *ElasticEmailImpl) SetAPIBase(address string) {
	m.apiBase = address
}
