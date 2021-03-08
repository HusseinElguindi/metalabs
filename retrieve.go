package metalabs

import (
	"encoding/json"
	"net/http"
)

// RetrieveLicense retrieves the passed license from MetaLabs using the API object
func (api *API) RetrieveLicense(license string) (*License, error) {
	req, err := http.NewRequest(http.MethodGet, EndpointLicense(license), nil)
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

	retrievedLicense := &License{}
	err = json.NewDecoder(resp.Body).Decode(retrievedLicense)
	if err != nil {
		return nil, err
	}

	return retrievedLicense, nil
}
