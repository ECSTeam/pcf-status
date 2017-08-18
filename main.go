package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ECSTeam/pcf-status/helpers"
	"github.com/ECSTeam/pcf-status/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	defaultPort         = 8080
	defaultReadTimeout  = 15
	defaultWriteTimeout = 15
)

var (
	routes = []helpers.RouteDefinition{
		models.ProductCollectionRoutesDefinition,
		models.ProductRoutesDefinition,
		models.VMCollectionRoutesDefinition,
		models.BuildpacksCollectionRoutesDefinition,
		models.InfoRoutesDefinition,
	}
)

// port and portFlag handle the port for this application. The default is
// arbitrarily set to 8080.
var (
	portFlag = flag.Int("port", defaultPort, "The port to use.")
	hostFlag = flag.String("host", "", "The host name.")

	opsManAddrFlag  = flag.String("opsman", "", "The OpsMan address.")
	opsUaaAddrFlag  = flag.String("opsmanuaa", "", "The OPsMan UAA address.")
	opsUserFlag     = flag.String("opsmanuser", "", "The OpsMan user.")
	opsPasswordFlag = flag.String("opsmanpassword", "", "The OpsMan password.")

	appsManAddrFlag  = flag.String("appsman", "", "The AppsMan address.")
	appsUaaAddrFlag  = flag.String("appsmanuaa", "", "The AppsMan UAA address.")
	appsUserFlag     = flag.String("appsmanuser", "", "The AppsMan user.")
	appsPasswordFlag = flag.String("appsmanpassword", "", "The AppsMan password.")
)

// APIs defines the api collection
type APIs map[helpers.APIType]helpers.API

// getParam from the env or cmd line. Env takes proirity.
func getParam(env string, cmd *string) string {
	param := os.Getenv(env)
	if len(strings.TrimSpace(param)) <= 0 {
		param = *cmd
	}

	log.Printf("[%s]: %s", env, param)
	return param
}

// createAPIs will generate the collection of apis.
func createAPIs() (apis APIs, err error) {

	opsManUaa := getParam("OPSMAN_UAA_ADDRESS", opsUaaAddrFlag)
	opsManAddr := getParam("OPSMAN_ADDRESS", opsManAddrFlag)
	opsUser := getParam("OPSMAN_USER", opsUserFlag)
	opsPassword := getParam("OPSMAN_PASSWORD", opsPasswordFlag)

	var api helpers.API
	if api, err = helpers.NewOpsManAPI(opsManUaa, opsManAddr, opsUser, opsPassword); err == nil {

		apis = APIs{}
		apis[helpers.OpsMan] = api

		appsManUaa := getParam("APPSMAN_UAA_ADDRESS", appsUaaAddrFlag)
		appsManAddr := getParam("APPSMAN_ADDRESS", appsManAddrFlag)
		appsUser := getParam("APPSMAN_USER", appsUserFlag)
		appsPassword := getParam("APPSMAN_PASSWORD", appsPasswordFlag)

		if api, err = helpers.NewAppsManAPI(appsManUaa, appsManAddr, appsUser, appsPassword); err == nil {
			apis[helpers.AppsMan] = api
		}
	}

	return apis, err
}

// main entry point.
func main() {

	log.Print("Initializing the application")
	flag.Parse()

	var apis APIs
	var err error
	if apis, err = createAPIs(); err == nil {
		r := mux.NewRouter()

		// CORS definition.
		orig := handlers.AllowedOrigins([]string{"*"})
		corsHandler := handlers.CORS(orig)(r)

		port := defaultPort
		if portStr := os.Getenv("PORT"); len(portStr) > 0 {
			if port, err = strconv.Atoi(portStr); err != nil {
				port = *portFlag
			}
		}

		host := os.Getenv("HOST")
		if len(strings.TrimSpace(host)) <= 0 {
			host = *hostFlag
		}

		addr := fmt.Sprintf("%s:%d", host, port)
		log.Printf("Starting application: %s", addr)

		srv := &http.Server{
			Handler:      corsHandler,
			Addr:         addr,
			WriteTimeout: defaultWriteTimeout * time.Second,
			ReadTimeout:  defaultReadTimeout * time.Second,
		}

		for _, route := range routes {
			log.Printf("Route: [%s] %s", route.Method, route.Path)
			r.Methods(route.Method).Path(route.Path).Handler(route.Handler(apis[route.APIType]))
		}

		err = srv.ListenAndServe()
	}

	log.Fatalf("Failed to start: %v", err)
}
