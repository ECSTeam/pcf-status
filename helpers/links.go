package helpers

import (
	"fmt"
	"net/http"
	"strings"
)

// Links is a collection of links
type Links struct {
	Items []Link `json:"links,omitempty"`
}

// Link is a reference to another document
type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// Append a link to the collection of links
func (links *Links) Append(rel string, href string) {
	links.Items = append(links.Items, Link{
		Href: href,
		Rel:  rel,
	})
}

// MakeURL creates a url based off the request and the items.
// NOTE: Only the base is kept from the reqest.
func MakeURL(req *http.Request, items ...string) string {

	schema := "http"
	if s := req.URL.Scheme; len(s) > 0 {
		schema = s
	}

	host := req.Host

	port := ""
	if p := req.URL.Port(); len(p) > 0 {
		port = ":" + p
	}

	url := fmt.Sprintf("%s://%s%s/%s", schema, host, port, strings.Join(items, "/"))

	return url
}
