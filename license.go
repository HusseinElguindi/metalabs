package metalabs

import (
	"errors"
	"time"
)

// Represents the errors with retrieving licenses
var (
	ErrLicenseNotFound = errors.New("license not found")
)

// License represents a MetaLabs license's details
type License struct {
	Email         string      `json:"email"`
	Key           string      `json:"key"`
	Unlocked      bool        `json:"unlocked"`
	Status        string      `json:"status"`
	CancelAt      interface{} `json:"cancel_at"`
	TrialEnd      time.Time   `json:"trial_end"`
	Created       int64       `json:"created"`
	Account       string      `json:"account"`
	Customer      string      `json:"customer"`
	Subscription  string      `json:"subscription"`
	PaymentMethod string      `json:"payment_method"`
	Plan          struct {
		Account        string   `json:"account"`
		Active         bool     `json:"active"`
		Product        string   `json:"product"`
		Price          string   `json:"price"`
		Name           string   `json:"name"`
		AllowUnbinding bool     `json:"allow_unbinding"`
		AllowReselling bool     `json:"allow_reselling"`
		Amount         int      `json:"amount"`
		Created        int64    `json:"created"`
		Currency       string   `json:"currency"`
		Roles          []string `json:"roles"`
		Recurring      struct {
			Interval      string `json:"interval"`
			IntervalCount int    `json:"interval_count"`
		} `json:"recurring"`
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"plan"`
	Release  string                 `json:"release"`
	Metadata map[string]interface{} `json:"metadata"`
	User     struct {
		Avatar        string `json:"avatar"`
		ID            string `json:"id"`
		Username      string `json:"username"`
		Discriminator string `json:"discriminator"`
		AccessToken   string `json:"access_token"`
		RefreshToken  string `json:"refresh_token"`
	} `json:"user"`
	ID string `json:"id"`
}
