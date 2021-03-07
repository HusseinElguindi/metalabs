# metalabs
A [MetaLabs](https://metalabs.io/) Golang API wrapper. Work in progress.

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
