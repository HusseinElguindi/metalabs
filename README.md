# MetaLabs API Wrapper
A [MetaLabs](https://metalabs.io/) Golang API wrapper.
Read the official [MetaLabs docs](https://docs.metalabs.io/reference).
Work in progress.

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

// Retrieve license data
licenseObj, err := api.RetrieveLicense(license)
if err != nil && err == metalabs.ErrLicenseNotFound {
    ...
}

// Update a license's metadata 
newMetadata := map[string]interface{}{"hwid": hwid}
licenseObj, err = api.UpdateLicenseMetadata(license, &newMetadata)
if err != nil && err == metalabs.ErrLicenseNotFound {
    ...
}
```

## Coverage
- [ ] Create license
- [x] Retrieve license
- [x] Update license metadata
- [ ] Revoke license
- [ ] List licenses

All unchecked features are quick to implement and being actively worked on.
