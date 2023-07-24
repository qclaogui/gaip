package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/qclaogui/golang-api-server/pkg/version"

	"log"
)

var port = "8080"

var sourceLink = "https://github.com/qclaogui/golang-api-server"

func handleHello(w http.ResponseWriter, _ *http.Request) {
	var ver = fmt.Sprintf("Build on %s [%s]", version.GoVersion, version.GetVersion())

	var name, _ = os.Hostname()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<br/><center><h1>Happy Coding </h1><br/><code>%s</code><p><a href=%q target=_blank>source code</a></p></center><hr><br/>"+
		"<center>this request was processed by host: %s</center>", ver, sourceLink, name)
}

// handleHealthz is a liveness probe.
func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Printf("Starting the service...[%s]", version.GetVersion())

	http.HandleFunc("/", handleHello)
	http.HandleFunc("/healthz", handleHealthz)

	// get port env var
	portEnv := os.Getenv("APP_PORT")
	if len(portEnv) > 0 {
		port = portEnv
	}
	log.Println("Listening on port:" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
