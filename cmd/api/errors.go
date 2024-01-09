package main

import (
	"fmt"
	"net/http"
)

// logError logs the given error using the application's logger.
//
// r: a pointer to an http.Request.
// err: the error to be logged.
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// errorResponse writes an error message to the ResponseWriter and logs the error.
//
// w: http.ResponseWriter - the ResponseWriter to write to.
// r: *http.Request - the incoming request.
// status: int - the HTTP status code to write.
// message: interface{} - the error message to write.
// Returns nothing.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := Envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// serverErrorResponse handles HTTP errors on the server side.
//
// It receives a http.ResponseWriter, a *http.Request, and an error and logs
// the error using the application's logError function. Then, it sets a message
// indicating that the server encountered a problem and could not process the
// request, and sends an HTTP 500 internal server error response using the
// errorResponse function.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse handles HTTP requests when the requested resource
// cannot be found.
//
// Takes in a http.ResponseWriter and a http.Request as parameters.
// Returns nothing.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse handles HTTP requests with unsupported methods,
// returning an error response with a message indicating the unsupported method.
//
// Parameters:
// - w: http.ResponseWriter
// - r: *http.Request
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// badRequestResponse is a function that handles HTTP bad requests by sending an error response.
//
// It takes in a http.ResponseWriter w, a http.Request r, and an error err. It returns no values.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

// failedValidationResponse handles the response when a validation fails.
//
// It takes in three parameters:
//   - w http.ResponseWriter: the http.ResponseWriter used to write the response.
//   - r *http.Request: the http.Request received.
//   - errors map[string]string: a map containing the validation errors.
//
// It does not return anything.
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

// editConflictResponse handles an edit conflict error when updating a record.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// It does not return anything.
func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, r, http.StatusConflict, message)
}
