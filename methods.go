package webmania

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseUrl = "https://webmaniabr.com/api/1/nfe/"

// Map extra request headers
type Headers map[string]string

// Make request and return the response
func (a *Api) execute(method string, path string, params interface{}, headers Headers, model interface{}) error {

	// init vars
	var url = baseUrl + path

	// init an empty payload
	payload := strings.NewReader("")

	// check for params
	if params != nil {

		// marshal params
		b, err := json.Marshal(params)
		if err != nil {
			return err
		}

		// set payload with params
		payload = strings.NewReader(string(b))

	}

	// set request
	request, _ := http.NewRequest(method, url, payload)
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Consumer-Key", a.ConsumerKey)
	request.Header.Add("X-Consumer-Secret", a.ConsumerSecret)
	request.Header.Add("X-Access-Token", a.AccessToken)
	request.Header.Add("X-Access-Token-Secret", a.AccessSecret)

	// add extra headers
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// init error response
	erm := &ErrorMessage{}

	// check for error message
	if err = json.Unmarshal(data, erm); err == nil && erm.Err != "" {
		return erm
	}

	// parse data
	return json.Unmarshal(data, model)

}

// Execute GET requests
func (a *Api) Get(path string, params interface{}, headers Headers, model interface{}) error {
	return a.execute("GET", path, params, headers, model)
}

// Execute POST requests
func (a *Api) Post(path string, params interface{}, headers Headers, model interface{}) error {
	return a.execute("POST", path, params, headers, model)
}

// Execute PUT requests
func (a *Api) Put(path string, params interface{}, headers Headers, model interface{}) error {
	return a.execute("PUT", path, params, headers, model)
}

// Execute DELETE requests
func (a *Api) Delete(path string, params interface{}, headers Headers, model interface{}) error {
	return a.execute("DELETE", path, params, headers, model)
}
