# metalabs
A [MetaLabs](https://metalabs.io/) Golang API wrapper.

## Installation

```bash
go get github.com/HusseinElguindi/metalabs
```

```golang
import "github.com/husseinelguindi/metalabs"
const metalabsAPIKey = "<metalabs api key>"

api := metalabs.NewAPI(metalabsAPIKey, nil)

license := "<license key>"
licenseData, _ := api.RetrieveLicense(license)
...
```
