package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Config struct {
	Port    int      `json:"port"`
	Targets []string `json:"targets"`
}

var config Config

func main() {
	loadConfig()

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/intermediate", handleIntermediate)
	http.HandleFunc("/reload-config", handleReloadConfig)

	serverAddress := fmt.Sprintf(":%d", config.Port)
	println("Server is running at http://localhost" + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

func loadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}

func handleReloadConfig(w http.ResponseWriter, r *http.Request) {
	loadConfig()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Configuration reloaded successfully."))
}

// handleMain is the handler for the root URL ("/"). It selects a random target URL
// from the configured list of URLs and initiates a redirect to the "/intermediate"
// endpoint, passing the selected target URL as a query parameter. This function
// serves as the initial point of contact for incoming requests, deciding randomly
// which target URL to redirect to next.
func handleMain(w http.ResponseWriter, r *http.Request) {
	target := config.Targets[rand.Intn(len(config.Targets))]

	http.Redirect(w, r, "/intermediate?target="+target, http.StatusFound)
}

// handleIntermediate is the handler for the "/intermediate" URL. It is called after
// the initial redirection from handleMain. This function reads the target URL from
// the query parameters, sets the "Referrer-Policy" header to "no-referrer" to enhance
// privacy by not sending the referrer header in the next redirect, and finally redirects
// the client to the actual target URL. This intermediate step allows for additional processing
// like logging, analytics, or manipulation of HTTP headers before the final redirection.
func handleIntermediate(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")

	w.Header().Set("Referrer-Policy", "no-referrer")

	http.Redirect(w, r, target, http.StatusFound)
}
