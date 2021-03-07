# metalabs
A [MetaLabs](https://metalabs.io/) Golang API wrapper.

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
licenseData, err := api.RetrieveLicense(license)
if err != nil {
    switch err {
    case metalabs.ErrLicenseNotFound:
        break
    case metalabs.ErrUnauthorized:
        break
    default:
        ...
    }
}
...
```
