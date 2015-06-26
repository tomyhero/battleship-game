package utils

import (
	"encoding/json"
	"net/http"
)

// RenderJSON method alow you to render json format output to browser.
// Also take care of setting contenct-type.
func RenderJSON(w http.ResponseWriter, stash map[string]interface{}) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	encoder := json.NewEncoder(w)
	encoder.Encode(stash)
}
