package metalabs

// APIVersion represents the version used to consume MetaLabs API
var APIVersion = "4"

// Represents the known MetaLabs endpoints
var (
	EndpointAPI      = "https://api.metalabs.io/v" + APIVersion + "/"
	EndpointLicenses = EndpointAPI + "licenses/"
	EndpointLicense  = func(license string) string { return EndpointLicenses + license }
)
