package metalabs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// UpdateLicenseMetadata updates and retrieves the passed license from MetaLabs using the API object
func (api *API) UpdateLicenseMetadata(license string, metadata *map[string]interface{}) (*RetrieveLicenseResponse, error) {
	body := map[string]interface{}{"metadata": metadata}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, EndpointLicense(license), buf)
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

	retrieveLicenseObj := &RetrieveLicenseResponse{}
	err = json.NewDecoder(resp.Body).Decode(retrieveLicenseObj)
	if err != nil {
		return nil, err
	}

	return retrieveLicenseObj, nil
}
