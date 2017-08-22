package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ECSTeam/pcf-status/helpers"
	"github.com/ECSTeam/pcf-status/models"
)

var (
	coreFragments = []string{"nav.html", "scripts.html", "layout.html"}
)

// staticFiles creates a static site.
func staticFiles(parts ...string) (route helpers.RouteDefinition) {

	sep := "/"
	path := strings.Join(parts, sep) + sep

	return helpers.RouteDefinition{
		Method:     http.MethodGet,
		Path:       sep + path + "{rest}",
		APIType:    helpers.None,
		RawHandler: http.StripPrefix(sep+path, http.FileServer(http.Dir(path))),
	}
}

// templateRoute defines the route
func templateRoute(path string, parts ...string) (route helpers.RouteDefinition) {
	return helpers.RouteDefinition{
		Method:  http.MethodGet,
		Path:    path,
		APIType: helpers.OpsMan,
		Handler: func(api helpers.API) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {

				var paths []string
				for _, part := range parts {
					paths = append(paths, filepath.Join("static", part))
				}

				for _, part := range coreFragments {
					paths = append(paths, filepath.Join("static", part))
				}

				if tmpl, err := template.ParseFiles(paths...); err == nil {

					temp, data := helpers.MakeData(api.(*helpers.OpsManAPI))
					if err = tmpl.ExecuteTemplate(w, temp, data); err != nil {
						log.Println(err.Error())
						http.Error(w, http.StatusText(500), 500)
					}

				} else {
					// Log the detailed error
					log.Println(err.Error())
					// Return a generic "Internal Server Error" message
					http.Error(w, http.StatusText(500), 500)
					return
				}
			}
		},
	}
}

var (
	routes = []helpers.RouteDefinition{
		models.OpsManProductCollectionRoute,
		models.OpsManProductRoute,
		models.OpsManVMTypesRoute,
		models.OpsManVMInstances,
		models.AppsManReleasesRoute,
		models.AppsManBuildpacksRoute,
		models.AppsManStemcellsRoute,
		models.AppsManInfoRoute,
		staticFiles("static", "js"),
		staticFiles("static", "css"),
		templateRoute("/", "default.html"),
		templateRoute("/releases", "releases.html"),
	}
)
