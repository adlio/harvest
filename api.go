package harvest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const CLIENT_VERSION = "1.0"
const HARVEST_DOMAIN = "api.harvestapp.com"
const HARVEST_API_VERSION = "v2"

type API struct {
	client       *http.Client
	Logger       logger
	BaseURL      string
	AccountID    string
	AccessToken  string
	RefreshToken string
	UserAgent    string
}

type logger interface {
	Debugf(string, ...interface{})
}

func NewTokenAPI(accountID string, accessToken string) *API {
	a := API{}
	a.client = http.DefaultClient
	a.BaseURL = "https://" + HARVEST_DOMAIN + "/" + HARVEST_API_VERSION
	a.AccountID = accountID
	a.AccessToken = accessToken
	return &a
}

func (a *API) GetPaginated(path string, args Arguments, target Pageable, afterFetch func()) error {
	page := 1
	args["page"] = fmt.Sprintf("%d", page)
	err := a.Get(path, args, target)
	if err != nil {
		return err
	}

	afterFetch()

	for target.HasNextPage() {
		page++
		args["page"] = fmt.Sprintf("%d", page)
		err = a.Get(path, args, target)
		if err != nil {
			return err
		}
		afterFetch()
	}
	return nil
}

func (a *API) Get(path string, args Arguments, target interface{}) error {
	url := fmt.Sprintf("%s%s", a.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, args.ToURLValues().Encode())

	req, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		return errors.Wrapf(err, "Invalid GET request %s", url)
	}
	a.AddHeaders(req)

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
	if err != nil {
		return errors.Wrapf(err, "Invalid PUT request %s", url)
	}
	a.AddHeaders(req)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

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
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(target)
		if err != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			return errors.Wrapf(err, "JSON decode failed on %s: %s", url, string(body))
		}
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
	if err != nil {
		return errors.Wrapf(err, "Invalid POST request %s", url)
	}
	a.AddHeaders(req)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

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
	if err != nil {
		return errors.Wrapf(err, "Invalid DELETE request %s", url)
	}
	a.AddHeaders(req)

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

// Applies relevant User-Agent, Accept, Authorization headers
func (a *API) AddHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")

	if a.UserAgent != "" {
		req.Header.Set("User-Agent", a.UserAgent)
	} else {
		req.Header.Set("User-Agent", defaultUserAgent())
	}

	if a.AccountID != "" {
		req.Header.Set("Harvest-Account-Id", a.AccountID)
	}

	if a.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+a.AccessToken)
	}
}

func (a *API) log(format string, args ...interface{}) {
	if a.Logger != nil {
		a.Logger.Debugf(format, args)
	}
}

func defaultUserAgent() string {
	return "github.com/adlio/harvest v" + CLIENT_VERSION
}
