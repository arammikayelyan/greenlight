package main

import (
	"encoding/json"
	"net/http"
)

// Handler for writing application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string.
	// json := `{"status": "available", "environment": %q, "version": %q}`
	// json = fmt.Sprintf(json, app.config.env, version)

	// Create a map which holds the information
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Set the "Content-Type: application/json"
	w.Header().Set("Content-Type", "application/json")

	// Send the []byte slice containing the JSON as the response body.
	w.Write(jsonData)
}
