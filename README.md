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
const metalabsAPIKey = "<metalabs api key>"

api := metalabs.NewAPI(metalabsAPIKey, nil)

license := "<license key>"
licenseObj, err := api.RetrieveLicense(license)
if err != nil && err == metalabs.ErrLicenseNotFound {
    ...
}
...
```

## Features
- [ ] Create license
- [x] Retrieve license
- [x] Update license metadata
- [ ] Revoke license
- [ ] List licenses

All unchecked features are being actively worked on.