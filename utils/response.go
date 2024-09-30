package utils

import (
	"encoding/json"
	"net/http"
)

// AppError adalah struktur yang digunakan untuk menampilkan pesan error dalam format JSON
type AppError struct {
	StatusCode int    `json:"status"`  // Status HTTP yang akan dikirimkan (misalnya, 400, 404, 500)
	Code       string `json:"code" `   // Kode error yang unik (misalnya, "INVALID_REQUEST")
	Message    string `json:"message"` // Pesan error yang dapat dimengerti oleh klien
	Err        error  `json:"-" `      // Error asli yang tidak dikirim ke klien
}

// NewAppError membuat instance baru dari AppError
func NewAppError(statusCode int, code, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

// RespondWithError mengirimkan response error dalam format JSON
func RespondWithError(w http.ResponseWriter, appErr *AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.StatusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": appErr.Message,
		"code":  appErr.Code,
	})
}

// RespondJSON mengirimkan response sukses dalam format JSON
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
