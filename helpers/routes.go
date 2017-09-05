package helpers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

// APIType defines the type of API to use.
type APIType string

// NavItemType defines the type of nav item.
type NavItemType int

const (

	// OpsMan is the indicator of the OpsMan API.
	OpsMan = APIType("OpsMan")

	// AppsMan is the indicator of the AppsMan API.
	AppsMan = APIType("AppsMan")

	// None is the no api response
	None = APIType("None")

	// NavItemDefault is the default value.
	NavItemDefault = NavItemType(0)

	// NavItemCurrent is the current value.
	NavItemCurrent = NavItemType(1)

	// NavItemGroup is the group value.
	NavItemGroup = NavItemType(2)
)

// navItems holds the collection of navigation items.
var navItems []NavItem

// NavItem defines the navigation items.
type NavItem struct {
	Title    string
	Type     NavItemType
	Link     string
	SubLinks []struct {
		Title string
		Link  string
	}
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

	title := name
	group := ""
	if len(name) > 0 {
		switch parts := strings.SplitN(name, "/", 2); len(parts) {
		case 1:
			navItems = append(navItems, NavItem{
				Title: name,
				Link:  path,
				Type:  NavItemDefault,
			})
		case 2:

			title = parts[1]
			group = parts[0]

			var found bool
			for index, item := range navItems {
				if item.Title == group {
					found = true
					navItems[index].SubLinks = append(item.SubLinks, struct {
						Title string
						Link  string
					}{
						Title: title,
						Link:  path,
					})
				}
			}

			if !found {
				navItems = append(navItems, NavItem{
					Title: group,
					Type:  NavItemGroup,
					SubLinks: []struct {
						Title string
						Link  string
					}{
						{
							Title: title,
							Link:  path,
						},
					},
				})
			}
		}
	}

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
						ID       string
						Title    string
						NavItems []NavItem
					}{
						Title:    title,
						NavItems: make([]NavItem, len(navItems)),
					}

					// Set the variable if it exists.
					if variables := mux.Vars(r); len(variables) > 0 {
						data.ID, _ = variables["id"]
					}

					copy(data.NavItems, navItems)
					for index, item := range data.NavItems {
						if item.Title == title || item.Title == group {
							data.NavItems[index].Type |= NavItemCurrent
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
