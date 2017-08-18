package helpers

import (
	"net/http"
)

// APIType defines the type of API to use.
type APIType string

const (

	// OpsMan is the indicator of the OpsMan API.
	OpsMan = APIType("OpsMan")

	// AppsMan is the indicator of the AppsMan API.
	AppsMan = APIType("AppsMan")
)

// RouteDefinition stores information about the route.
type RouteDefinition struct {
	Method  string
	Path    string
	APIType APIType
	Handler func(api API) http.HandlerFunc
}
