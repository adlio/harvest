package harvest

import (
	"fmt"
)

const HARVEST_DOMAIN = "harvestapp.com"

type API struct {
	BaseURL   string
	SubDomain string
	User      string
	Password  string
}

func NewBasicAuthAPI(subdomain, user, password string) API {
	a := API{}
	a.SubDomain = subdomain
	a.User = user
	a.Password = password
	a.BaseURL = fmt.Sprintf("https://%s.%s", subdomain, HARVEST_DOMAIN)
	return a
}
