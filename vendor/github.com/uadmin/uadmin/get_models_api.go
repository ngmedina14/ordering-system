package uadmin

import (
	"net/http"
)

// GetModelsAPI returns a list of models
func GetModelsAPI(w http.ResponseWriter, r *http.Request, session *Session) {
	response := []string{}
	for _, v := range modelList {
		response = append(response, getModelName(v))
	}
	ReturnJSON(w, r, response)
}
