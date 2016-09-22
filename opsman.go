package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//	"github.com/cloudfoundry-community/gogobosh"

// OpsManClient is the client for the ops manager.
type OpsManClient struct {
	address string
	token   string
}

// DiagnosticReport is the result from /api/v0/diagnostic_report
type DiagnosticReport struct {
	Versions struct {
		Schema  string `json:"installation_schema_version"`
		Meta    string `json:"metadata_version"`
		Release string `json:"release_version"`
	} `json:"versions"`

	Products struct {
		Deployed []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"deployed"`
	} `json:"added_products"`
}

// OAuthPayload is the wrapper for the oauth token.
type OAuthPayload struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

// HeaderValue is the authorization header value.
func (payload OAuthPayload) HeaderValue() (output string) {
	output = fmt.Sprintf("%s %s", payload.TokenType, payload.AccessToken)
	return
}

// NewOpsManClient will create a new ops manager client.
func NewOpsManClient() (opsManClient *OpsManClient, err error) {
	opsManClient = &OpsManClient{
		address: os.Getenv("UAA_ADDRESS"),
	}

	user := os.Getenv("OPSMAN_USER")
	log.Printf("Logging in as: %s", user)

	path := fmt.Sprintf("/uaa/oauth/token?grant_type=password&username=%s&password=%s", user, os.Getenv("OPSMAN_PASSWORD"))

	err = opsManClient.callURL(path, func(data []byte) (e error) {
		var token OAuthPayload
		if e = json.Unmarshal(data, &token); e == nil {
			opsManClient.token = token.HeaderValue()
		}
		return
	})

	return
}

// GetInfo will return the info.
func (opsManClient *OpsManClient) GetInfo(report *DiagnosticReport) (err error) {
	err = opsManClient.callURL("/api/v0/diagnostic_report", func(data []byte) (e error) {
		e = json.Unmarshal(data, report)
		return
	})

	return
}

// callUrl will call the url and add the appropriate headers.
func (opsManClient *OpsManClient) callURL(path string, operation func([]byte) error) (err error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url := opsManClient.address + path
	var req *http.Request

	if req, err = http.NewRequest("GET", url, nil); err == nil {
		if len(opsManClient.token) != 0 {
			req.Header.Add("Authorization", opsManClient.token)
		} else {
			// shortcut for opsman login.
			req.SetBasicAuth("opsman", "")
		}

		var resp *http.Response
		if resp, err = client.Do(req); err == nil {
			defer resp.Body.Close()

			if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
				var resBody []byte
				if resBody, err = ioutil.ReadAll(resp.Body); err == nil {
					err = operation(resBody)
				}
			} else {
				err = errors.New(resp.Status)
			}
		}
	}
	return
}
