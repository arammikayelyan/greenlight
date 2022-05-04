package main

import (
	"fmt"
	"net/http"
)

// Handler for writing application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string.
	json := `{"status": "available", "environment": %q, "version": %q}`
	json = fmt.Sprintf(json, app.config.env, version)

	// Set the "Content-Type: application/json"
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON as the HTTP response body.
	w.Write([]byte(json))
}
