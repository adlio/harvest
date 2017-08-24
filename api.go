package harvest

import (
	"bytes"
	"encoding/json"
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
	a.client = http.DefaultClient
	a.SubDomain = subdomain
	a.User = user
	a.Password = password
	a.BaseURL = fmt.Sprintf("https://%s.%s", subdomain, HARVEST_DOMAIN)
	return &a
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	url := fmt.Sprintf("%s%s", a.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	req.Header.Set("Accept", "application/json")
	if a.User != "" && a.Password != "" {
		req.SetBasicAuth(a.User, a.Password)
	}
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		return errors.Errorf("HTTP request failure on %s: %s %s", url, string(body), err)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	if err != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.Wrapf(err, "JSON decode failed on %s: %s", url, string(body))
	}

	return nil
}

func (a *API) Put(path string, args Arguments, postData interface{}, target interface{}) error {
	url := fmt.Sprintf("%s%s", a.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	buffer := new(bytes.Buffer)
	if postData != nil {
		json.NewEncoder(buffer).Encode(postData)
	}

	req, err := http.NewRequest("PUT", urlWithParams, buffer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/adlio/harvest")
	if a.User != "" && a.Password != "" {
		req.SetBasicAuth(a.User, a.Password)
	}
	if err != nil {
		return errors.Wrapf(err, "Invalid PUT request %s", url)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		return errors.Wrapf(err, "HTTP request failure on %s: %s %s", url, string(body), err)
	}

	// Harvest V1 API returns an empty response, with a Location header including the
	// URI of the created object (e.g. /projects/254454)
	redirectDestination := resp.Header.Get("Location")
	if redirectDestination != "" {
		return a.Get(redirectDestination, args, target)
	} else {
		return errors.Errorf("PUT to %s failed to return a Location header. This means we couldn't fetch the new state of the record.", url)
	}

	return nil
}

func (a *API) Post(path string, args Arguments, postData interface{}, target interface{}) error {
	url := fmt.Sprintf("%s%s", a.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	buffer := new(bytes.Buffer)
	if postData != nil {
		json.NewEncoder(buffer).Encode(postData)
	}

	req, err := http.NewRequest("POST", urlWithParams, buffer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "github.com/adlio/harvest")
	if a.User != "" && a.Password != "" {
		req.SetBasicAuth(a.User, a.Password)
	}
	if err != nil {
		return errors.Wrapf(err, "Invalid POST request %s", url)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		return errors.Wrapf(err, "HTTP request failure on %s: %s %s", url, string(body), err)
	}

	// Harvest V1 API returns an empty response, with a Location header including the
	// URI of the created object (e.g. /projects/254454)
	redirectDestination := resp.Header.Get("Location")
	if redirectDestination != "" {
		return a.Get(redirectDestination, args, target)
	} else {
		return errors.Errorf("POST to %s failed to return a Location header. This means we couldn't fetch the new state of the record.", url)
	}

	return nil
}

func (a *API) Delete(path string, args Arguments) error {
	url := fmt.Sprintf("%s%s", a.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	req, err := http.NewRequest("DELETE", urlWithParams, nil)
	req.Header.Set("Accept", "application/json")
	if a.User != "" && a.Password != "" {
		req.SetBasicAuth(a.User, a.Password)
	}
	if err != nil {
		return errors.Wrapf(err, "Invalid DELETE request %s", url)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "HTTP request failure on %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		return errors.Wrapf(err, "HTTP request failure on %s: %s %s", url, string(body), err)
	}

	return nil
}
