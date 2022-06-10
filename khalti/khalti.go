package khalti

import "net/http"

const (
	defaultBaseURL = "https://khalti.com/"
)

type KhaltiService struct {
	ClientID  string
	SecretKey string
	BaseUrl   string
	Client    *http.Client
}

// NewClient returns a new Khalti API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewKhaltiClient(clientID string, secret string, client *http.Client) (*KhaltiService, error) {
	khaltiServer := &KhaltiService{
		ClientID:  clientID,
		SecretKey: secret,
		BaseUrl:   defaultBaseURL,
		Client:    &http.Client{},
	}
	return khaltiServer, nil
}
