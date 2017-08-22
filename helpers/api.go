package helpers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (

	// MimeTypePlainText defines the plain text mime type
	MimeTypePlainText = "text/plain"

	// MimeTypeJSON defines the json mime type
	MimeTypeJSON = "application/json"

	// MimeTypeHTML defines the html mime type
	MimeTypeHTML = "text/html"
)

// callHandler will handle calls from a url.
type callHandler func(int, []byte) error

// API is the interface to the OpsManApi
type API interface {
	Get(string, interface{}) error
	MakeAPIURL(items ...string) string
}

// opsManAPI is the implementation of the OpsManApi interface
type baseAPI struct {
	clientID     string
	clientSecret string
	uaaAddress   string
	address      string
	user         string
	password     string
	token        string
	urlCreator   func(items ...string) string
}

// NewBaseAPI from the child invoking it.
func (api *baseAPI) bindURLMethod(urlCreator func(items ...string) string) (err error) {
	api.urlCreator = urlCreator
	return err
}

// MakeApiURL will create the url
func (api *baseAPI) MakeAPIURL(items ...string) string {
	return api.urlCreator(items...)
}

// GetSlice will return a slice from the path.
func (api *baseAPI) Get(path string, value interface{}) (err error) {
	return api.callURLWithToken(http.MethodGet, api.MakeAPIURL(path), func(status int, body []byte) (e error) {
		return json.Unmarshal(body, value)
	})
}

// RequestWrapper will wrap the request.
func RequestWrapper(handler func(*http.Request) (int, string, []byte, error)) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		var contentType string
		var status int
		var body []byte
		var err error

		if status, contentType, body, err = handler(req); err != nil {
			status = http.StatusNotAcceptable
			contentType = MimeTypePlainText
			body = []byte(err.Error())
		}

		writer.Header().Add("Content-Type", contentType)
		writer.WriteHeader(status)
		writer.Write(body)
	}
}

// CreateHandler creates an http handler
func CreateHandler(api API, handler func(*http.Request, API) (interface{}, error)) http.HandlerFunc {
	return RequestWrapper(func(req *http.Request) (status int, contentType string, body []byte, err error) {
		status = http.StatusNotImplemented
		contentType = MimeTypePlainText

		var data interface{}
		if data, err = handler(req, api); err == nil {
			if body, err = json.Marshal(data); err == nil {
				status = http.StatusOK
				contentType = MimeTypeJSON
			}
		}

		return status, contentType, body, err
	})
}

// getToken return the auth token.
func (api *baseAPI) getToken() (token string, err error) {

	if token = api.token; len(token) == 0 {

		log.Printf("Getting new auth token: %s", api.user)

		var client *http.Client
		if client, err = newClient(); err == nil {

			// using this: https://docs.cloudfoundry.org/api/uaa/#password-grant
			// but the documentation does not include the auth header.
			data := url.Values{}
			data.Add("grant_type", "password")
			data.Add("response_type", "token")
			data.Add("username", api.user)
			data.Add("password", api.password)

			body := data.Encode()
			addr := fmt.Sprintf("%s/oauth/token", api.uaaAddress)

			var req *http.Request
			if req, err = http.NewRequest("POST", addr, bytes.NewBufferString(body)); err == nil {

				// We must set the opsman client id and client secret.
				req.SetBasicAuth(api.clientID, api.clientSecret)
				req.Header.Add("Accept", "application/json")
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				var resp *http.Response

				log.Printf("Getting token from: %s", addr)
				if resp, err = client.Do(req); err == nil {
					if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {

						var data struct {
							TokenType   string `json:"token_type"`
							AccessToken string `json:"access_token"`

							// TODO: Include the following fields...
							// expires_in
							// refresh_token
						}

						var resBody []byte
						if resBody, err = ioutil.ReadAll(resp.Body); err == nil {
							if err = json.Unmarshal(resBody, &data); err == nil {
								token = fmt.Sprintf("%s %s", strings.Title(data.TokenType), data.AccessToken)
								log.Printf("Got token type: %s", data.TokenType)
								api.token = token
							}
						}
					} else {

						details := ""
						var resBody []byte
						if resBody, err = ioutil.ReadAll(resp.Body); err == nil {
							log.Printf("Error Body [%s]: %s", resp.Status, string(resBody))

							var errorDetails struct {
								Description string `json:"error_description"`
							}

							if err = json.Unmarshal(resBody, &errorDetails); err == nil {
								details = fmt.Sprintf("\nDetails: %s", errorDetails.Description)
							}
						}

						err = NewErrorf("UAA authorization failure: [%d] %s%s", resp.StatusCode, resp.Status, details)
					}
				}
			}
		}
	} else {
		// https://docs.cloudfoundry.org/api/uaa/#refresh-token
		// TODO: We need to check the token!
	}

	return token, err
}

// callUrl will call the url and add the appropriate headers.
func (api *baseAPI) callURLWithToken(method string, url string, operation callHandler) (err error) {
	var token string
	if token, err = api.getToken(); err == nil {
		callURL(method, url, map[string]string{
			"Authorization": token,
		}, operation)
	}

	return err
}

// callURL will call a URL
func callURL(method string, url string, headers map[string]string, operation callHandler) (err error) {
	var client *http.Client
	if client, err = newClient(); err == nil {
		var req *http.Request
		if req, err = http.NewRequest(method, url, nil); err == nil {

			req.Header.Add("Accept", "application/json") // Make sure the output is a json.
			if headers != nil {
				for name, value := range headers {
					if len(req.Header.Get(name)) == 0 {
						req.Header.Add(name, value)
					} else {
						req.Header.Set(name, value)
					}
				}
			}

			var resp *http.Response

			log.Printf("Calling: %s", req.URL.String())
			if resp, err = client.Do(req); err == nil {
				defer resp.Body.Close()

				var body []byte
				code := resp.StatusCode
				if code >= http.StatusOK && code < http.StatusBadRequest {
					if body, err = ioutil.ReadAll(resp.Body); err == nil {
						err = operation(code, body)
					}
				}
			}
		}
	}

	return err
}

// newClient creates a new client.
func newClient() (client *http.Client, err error) {
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return client, err
}
