package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const HARVEST_DOMAIN = "harvestapp.com"

type API struct {
	client    *http.Client
	BaseURL   string
	SubDomain string
	User      string
	Password  string
}

func NewBasicAuthAPI(subdomain, user, password string) *API {
	a := API{}
	a.SubDomain = subdomain
	a.User = user
	a.Password = password
	a.BaseURL = fmt.Sprintf("https://%s.%s", subdomain, HARVEST_DOMAIN)
	return &a
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	url := fmt.Sprintf("%s/%s", a.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		return errors.Errorf("HTTP request failure on %s: %s %s", url, string(body), err)
	}

	return nil
}
