package error

import (
	"net/http"
)

// Response function create a error response with a error message
func Response(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
		return
	}
}
