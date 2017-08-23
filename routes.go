package main

import (
	"github.com/ECSTeam/pcf-status/helpers"
	"github.com/ECSTeam/pcf-status/models"
)

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
		helpers.StaticFiles("static", "fonts"),
		helpers.StaticFiles("static", "js"),
		helpers.StaticFiles("static", "css"),
		helpers.TemplateRoute("Home", "/", "default.html"),
		helpers.TemplateRoute("Releases", "/releases", "releases.html"),
		helpers.TemplateRoute("Stemcells", "/stemcells", "stemcells.html"),
		helpers.TemplateRoute("VMs", "/vms", "vms.html"),
	}
)
