package helpers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// APIType defines the type of API to use.
type APIType string

const (

	// OpsMan is the indicator of the OpsMan API.
	OpsMan = APIType("OpsMan")

	// AppsMan is the indicator of the AppsMan API.
	AppsMan = APIType("AppsMan")

	// None is the no api response
	None = APIType("None")
)

// navItems holds the collection of navigation items.
var navItems []NavItem

// NavItem defines the navigation items.
type NavItem struct {
	Title   string
	Current bool
	Link    string
}

// RouteDefinition stores information about the route.
type RouteDefinition struct {
	Method     string
	Path       string
	APIType    APIType
	Handler    func(api API) http.HandlerFunc
	RawHandler http.Handler
}

// StaticFiles creates a static site.
func StaticFiles(parts ...string) (route RouteDefinition) {

	sep := "/"
	path := strings.Join(parts, sep) + sep

	return RouteDefinition{
		Method:     http.MethodGet,
		Path:       sep + path + "{rest}",
		APIType:    None,
		RawHandler: http.StripPrefix(sep+path, http.FileServer(http.Dir(path))),
	}
}

// TemplateRoute defines the route
func TemplateRoute(name string, path string, definition string) (route RouteDefinition) {

	navItems = append(navItems, NavItem{
		Title: name,
		Link:  path,
	})

	return RouteDefinition{
		Method:  http.MethodGet,
		Path:    path,
		APIType: OpsMan,
		Handler: func(api API) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {

				paths := []string{
					filepath.Join("static", definition),
					filepath.Join("static", "layout.html"),
				}

				var err error
				var tmpl *template.Template
				if tmpl, err = template.ParseFiles(paths...); err == nil {

					data := struct {
						Title    string
						NavItems []NavItem
					}{
						Title:    name,
						NavItems: make([]NavItem, len(navItems)),
					}

					copy(data.NavItems, navItems)
					for index, item := range data.NavItems {

						if item.Title == name {
							data.NavItems[index].Current = true
							break
						}
					}

					err = tmpl.ExecuteTemplate(w, "layout", data)
				}

				if err != nil {
					status := http.StatusNotFound
					log.Println(err.Error())
					http.Error(w, http.StatusText(status), status)
				}
			}
		},
	}
}
