package pkg

import (
	"fmt"
	"net/http"
)

// Response function create a error response with a error message
func ErrCheck(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": { "message": "` + err.Error() + `"}}`))
	}
	return
}
