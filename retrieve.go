package metalabs

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Represents the errors with retrieving licenses
var (
	ErrLicenseNotFound = errors.New("license not found")
)

// RetrieveLicenseResponse represents the response object for v4/licenses
type RetrieveLicenseResponse struct {
	Email         string      `json:"email"`
	Key           string      `json:"key"`
	Unlocked      bool        `json:"unlocked"`
	Status        string      `json:"status"`
	CancelAt      interface{} `json:"cancel_at"`
	TrialEnd      time.Time   `json:"trial_end"`
	Created       int64       `json:"created"`
	Account       string      `json:"account"`
	Customer      string      `json:"customer"`
	Subscription  string      `json:"subscription"`
	PaymentMethod interface{} `json:"payment_method"`
	Plan          struct {
		Account        string   `json:"account"`
		Active         bool     `json:"active"`
		Product        string   `json:"product"`
		Price          string   `json:"price"`
		Name           string   `json:"name"`
		AllowUnbinding bool     `json:"allow_unbinding"`
		AllowReselling bool     `json:"allow_reselling"`
		Amount         int      `json:"amount"`
		Created        int64    `json:"created"`
		Currency       string   `json:"currency"`
		Roles          []string `json:"roles"`
		Recurring      struct {
			Interval      string `json:"interval"`
			IntervalCount int    `json:"interval_count"`
		} `json:"recurring"`
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"plan"`
	Release  string                  `json:"release"`
	Metadata *map[string]interface{} `json:"metadata"`
	User     struct {
		Avatar        string `json:"avatar"`
		ID            string `json:"id"`
		Username      string `json:"username"`
		Discriminator string `json:"discriminator"`
		AccessToken   string `json:"access_token"`
		RefreshToken  string `json:"refresh_token"`
	} `json:"user"`
	ID string `json:"id"`
}

// RetrieveLicense retrieves the passed license from MetaLabs using the API object
func (api *API) RetrieveLicense(license string) (*RetrieveLicenseResponse, error) {
	req, err := http.NewRequest("GET", EndpointLicense(license), nil)
	if err != nil {
		return nil, err
	}
	api.authorizeRequest(req)

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		break
	case 404:
		return nil, ErrLicenseNotFound
	case 401:
		return nil, ErrUnauthorized
	default:
		return nil, ErrUnknown
	}

	retrieveLicenseObj := &RetrieveLicenseResponse{}
	err = json.NewDecoder(resp.Body).Decode(retrieveLicenseObj)
	if err != nil {
		return nil, err
	}

	return retrieveLicenseObj, nil
}
