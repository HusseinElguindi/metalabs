package metalabs

import "net/http"

// RevokeLicense revokes/deletes a license from MetaLabs
func (api *API) RevokeLicense(license string) error {
	req, err := http.NewRequest(http.MethodDelete, EndpointLicense(license), nil)
	if err != nil {
		return err
	}
	api.authorizeRequest(req)

	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		break
	case 404:
		return ErrLicenseNotFound
	case 401:
		return ErrUnauthorized
	default:
		return ErrUnknown
	}

	return nil
}
