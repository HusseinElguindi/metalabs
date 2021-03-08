# MetaLabs API Wrapper
A [MetaLabs](https://metalabs.io/) Golang API wrapper.
Read the official [MetaLabs docs](https://docs.metalabs.io/reference).

## Installation
```bash
go get github.com/HusseinElguindi/metalabs
```

## Usage
```golang
import "github.com/husseinelguindi/metalabs"

// Create a new authenticated api instance
const metalabsAPIKey = "<metalabs api key>"
api := metalabs.NewAPI(metalabsAPIKey, nil)

// The license to perform actions on
license := "<license key>"

// Create a license
createdLicense, _ := api.CreateLicense("<plan id>", "<email>")

// Retrieve license data
retrievedLicense, _ := api.RetrieveLicense(license)

// Update a license's metadata
newMetadata := map[string]interface{}{"hwid": hwid}
updatedLicense, _ := api.UpdateLicenseMetadata(license, &newMetadata)

// Revoke a license
_ = api.RevokeLicense(license)

// List licenses
licensesList, _ := api.ListLicenses()

// Error handling
_, err := api.RetrieveLicense(license)
switch err {
case metalabs.ErrLicenseNotFound:
    ...
case metalabs.ErrUnauthorized:
    ...
case metalabs.ErrUnknown:
    ...
default:
    ...
}
```

## Coverage
- [x] Create license
- [x] Retrieve license
- [x] Update license metadata
- [x] Revoke license
- [x] List licenses

## License
- [BSD 3-Clause License](https://github.com/HusseinElguindi/metalabs/blob/master/LICENSE)
