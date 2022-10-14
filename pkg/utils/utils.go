package utils

import (
	"encoding/json"
	"net/http"
)

// ParseBody adalah fungsi untuk mengubah body request menjadi struct.
// Pake NewDecoder dan Decode karena lebih efisien daripada Unmarshal (sudah ada io.read dan unmarshal)
func ParseBody(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	return err
}

// RespondWithJson adalah fungsi untuk mengirimkan response dengan format json
// sekaligus mengatur header content-type jadi json dan menaruh status code
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError untuk menaruh pesan error dan status code ke response via RespondWithJson
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}
