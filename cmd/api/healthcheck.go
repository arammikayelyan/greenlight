package main

import (
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

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(
			w,
			"The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
		return
	}
}
