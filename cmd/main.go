package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/qclaogui/golang-api-server/pkg/version"
	"golang.org/x/exp/slog"
)

var port = "8080"

var sourceLink = "https://github.com/qclaogui/golang-api-server"

func handleHello(w http.ResponseWriter, r *http.Request) {
	slog.Info("new request", "method", r.Method, "uri", r.URL.String(), "userAgent", r.Header.Get("User-Agent"))

	var ver = fmt.Sprintf("Build on %s [%s]", version.GoVersion, version.GetVersion())
	var name, _ = os.Hostname()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprintf(w, "<br/><center><h1>Happy Coding </h1><br/><code>%s</code><p><a href=%q target=_blank>source code</a></p></center><hr><br/>"+
		"<center>this request was processed by host: %s</center>", ver, sourceLink, name)
}

// handleHealthz is a liveness probe.
func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	lvl := slog.LevelInfo
	// LOG_LEVEL is set, let's default to the desired level
	if lvlEnv, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if err := lvl.UnmarshalText([]byte(lvlEnv)); err != nil {
			slog.Error("unknown log level specified, choises are [DEBUG, INFO, WARN, ERROR]", errors.New(lvlEnv))
			os.Exit(-1)
		}
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: lvl})))

	slog.Info("Starting the service...", "version", version.GetVersion())

	http.HandleFunc("/", handleHello)
	http.HandleFunc("/healthz", handleHealthz)

	// get port env var
	portEnv := os.Getenv("APP_PORT")
	if len(portEnv) > 0 {
		port = portEnv
	}

	slog.Info("Listening on", "port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("ListenAndServe failed", "error", err)
		os.Exit(-1)
	}
}
