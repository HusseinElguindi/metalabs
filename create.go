package metalabs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// CreateLicense creates a new license with the passed details
func (api *API) CreateLicense(planID, email string) (*License, error) {
	body := map[string]interface{}{
		"plan":  planID,
		"email": email,
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, EndpointLicenses, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
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

	createdLicense := &License{}
	err = json.NewDecoder(resp.Body).Decode(createdLicense)
	if err != nil {
		return nil, err
	}

	return createdLicense, nil
}
