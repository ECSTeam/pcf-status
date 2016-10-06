/*main is the entry point package for this application.
 *
 */
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rs/cors"
)

// port and portFlag handle the port for this application. The default is
// arbitrarily set to 8080.
var port int
var host string
var (
	portFlag = flag.Int("port", 8080, "The port to use.")
	hostFlag = flag.String("host", "", "The host name.")
)

// init the program.
func init() {
	var err error
	log.Print("Initializing the application")
	flag.Parse()

	portStr := os.Getenv("PORT")
	if port, err = strconv.Atoi(portStr); err != nil {
		port = *portFlag
		err = nil
	}

	host = os.Getenv("HOST")
	if len(strings.TrimSpace(host)) <= 0 {
		host = *hostFlag
	}
}

// reportError will report an error result.
func reportError(err error, format string, resp http.ResponseWriter, status int) {
	log.Printf("Unable to get status: %s", err.Error())
	resp.WriteHeader(status)
	resp.Write([]byte(err.Error()))
}

// handlePanic will capture any panics and return a message to the output.
func handlePanic(resp http.ResponseWriter, status int) {
	if p := recover(); p != nil {

		messageFmt := "Unhandled panic: %s"
		var err error

		switch p.(type) {
		case nil:
			// normal case, just ignore.
		case string:
			messageFmt = p.(string)
			err = errors.New(messageFmt)
		case error:
			err = p.(error)
		default:
			err = errors.New(fmt.Sprint(p))
		}

		if err != nil {
			reportError(err, messageFmt, resp, status)
		}
	}
}

// httpHandler handles the http requests.
func httpHandler(resp http.ResponseWriter, req *http.Request) {
	defer handlePanic(resp, http.StatusInternalServerError)

	var err error
	switch req.URL.Path {
	case "/versions":
		{
			switch req.Method {
			case "GET":
				{
					var status *Status
					if status, err = NewStatus(); err == nil {
						var bytes []byte
						if bytes, err = json.Marshal(status); err == nil {
							// First, add the headers as the Write will start streaming right away.
							resp.Header().Set("Content-Type", "application/json")
							resp.Header().Set("Cache-Control", "no-cache")

							// TODO: we may need to remove the \u0001 and \u0002 from the
							//       result because this was a hack to sort the labels
							//       correctly. The SVG though doesn't mind.

							_, err = resp.Write(bytes)
						}
					}
				}
			}
		}
	default:
		err = errors.New("Unknown route.")
	}

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
	}
}

// main entry point.
func main() {
	log.Print("Starting application")

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET"},
	})

	handler := http.HandlerFunc(httpHandler)

	endpoint := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(endpoint, c.Handler(handler)); err != nil {
		log.Fatalf("Failed to listen on endpoint '%s': %s", endpoint, err.Error())
	} else {
		log.Printf("Started application on endpoint: '%s'", endpoint)
	}
}
