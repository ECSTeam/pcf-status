package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ECSTeam/pcf-status/helpers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	defaultPort         = 8080
	defaultReadTimeout  = 15
	defaultWriteTimeout = 15
)

// port and portFlag handle the port for this application. The default is
// arbitrarily set to 8080.
var (
	portFlag = flag.Int("port", defaultPort, "The port to use.")
	hostFlag = flag.String("host", "", "The host name.")
)

// APIs defines the api collection
type APIs map[helpers.APIType]helpers.API

// createAPIs will generate the collection of apis.
func createAPIs() (apis APIs, err error) {

	param := os.Getenv("OPSMAN")

	var config helpers.OpsManConfig
	if err = json.Unmarshal([]byte(param), &config); err == nil {

		// First find the opsman API
		var opsMan *helpers.OpsManAPI
		if opsMan, err = helpers.NewOpsManAPI(config); err == nil {
			apis = APIs{
				helpers.None:   nil,
				helpers.OpsMan: opsMan,
			}

			var api helpers.API
			if api, err = helpers.NewAppsManAPI(opsMan); err == nil {
				apis[helpers.AppsMan] = api
			}
		}
	}

	if err != nil {
		log.Printf("API Error: %s", err)
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

			var handler http.Handler
			custom := "raw"
			if handler = route.RawHandler; handler == nil {
				usedAPI := apis[route.APIType]
				handler = route.Handler(usedAPI)
				custom = "gen"
			}

			r.Methods(route.Method).Path(route.Path).Handler(handler)
			log.Printf("Route (%s): [%s] %s", custom, route.Method, route.Path)
		}

		err = srv.ListenAndServe()
	}

	log.Fatalf("Failed to start: %v", err)
}
