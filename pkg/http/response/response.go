package response

import (
	"net/http"
)

type HTTPResponse struct {
	Data any `json:"data"`
}

// Flush checks if it is stream response and sends any buffered data to the client.
func Flush(w http.ResponseWriter) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
