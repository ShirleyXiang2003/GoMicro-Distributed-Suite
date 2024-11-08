package main

import (
	"encoding/json"
	"net/http"
)

// jsonResponse represents the structure of the JSON response sent by the Broker handler.
type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// readJSON reads the request body and decodes it into the provided data structure.
func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// writeJSON writes a JSON response with the provided status code, data, and optional headers.
func (app *Config) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJSON generates a JSON error response with the provided error message and status code.
func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := jsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return app.writeJSON(w, statusCode, payload)
}
