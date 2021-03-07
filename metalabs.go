package metalabs

import (
	"errors"
	"fmt"
	"net/http"
)

// Represents global API errors
var (
	ErrUnauthorized = errors.New("licence not found")
	ErrUnknown      = errors.New("an unknown error occurred")
)

// API represents the authentication object used to call the MetaLabs API
type API struct {
	Client *http.Client
	Key    string
}

// NewAPI creates an object to consume the MetaLabs API, with authorization headers set
func NewAPI(APIKey string, client *http.Client) *API {
	// Create client if no custom client is passed
	if client == nil {
		client = &http.Client{}
	}

	// Return API object
	return &API{
		Client: client,
		Key:    APIKey,
	}
}

func (api *API) authorizeRequest(req *http.Request) {
	// Overwrite or add the authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.Key))
}
