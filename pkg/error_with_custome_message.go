package pkg

import (
	"net/http"
)

func ErrWithCusMsg(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(`{"error": { "message": "` + msg + `" }}`))
}
