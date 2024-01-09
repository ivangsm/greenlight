package main

import (
	"net/http"
)

// healthcheckHandler handles the healthcheck request for the application.
//
// w: an http.ResponseWriter used to write the response.
// r: a pointer to an http.Request that contains the request information.
// It returns no values.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
