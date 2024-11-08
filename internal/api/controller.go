package api

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

type GetHealthResponse struct {
	Message string `json:"message"`
}

func (c *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, GetHealthResponse{Message: "Welcome"})
}

// writeResponse sets appropriate headers and json encodes the response
func writeResponse(w http.ResponseWriter, code int, result interface{}) {
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
