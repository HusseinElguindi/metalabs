package metalabs

import (
	"encoding/json"
	"net/http"
)

// List represents a list of Licenses and other data
type List struct {
	Data    []License `json:"data"`
	Total   int       `json:"total"`
	Page    int       `json:"page"`
	HasMore bool      `json:"has_more"`
}

// ListLicenses lists the licenes for the passed API instance
func (api *API) ListLicenses() (*List, error) {
	req, err := http.NewRequest(http.MethodGet, EndpointLicenses, nil)
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

	licenesList := &List{}
	err = json.NewDecoder(resp.Body).Decode(licenesList)
	if err != nil {
		return nil, err
	}

	return licenesList, nil
}
