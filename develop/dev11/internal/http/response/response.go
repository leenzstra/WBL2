package response

import (
	"encoding/json"
	"net/http"
)

// База отмета от сервера
type response[T any] struct {
	// Сообщение об ошибке
	Error  string `json:"error,omitempty"`
	// Результат/данные
	Result T      `json:"result,omitempty"`
}

// Создание ответа - ошибки
func err[T any](errMsg string) *response[T] {
	return &response[T]{
		Error:  errMsg,
		Result: *new(T),
	}
}

// Создание ответа с данными
func ok[T any](data T) *response[T] {
	return &response[T]{
		Result: data,
	}
}

// Отправка ответа как JSON
func sendjson(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Отправка ответа с ошибкой как JSON
func SendJSONErr(w http.ResponseWriter, r *http.Request, e error, statusCode int) {
	sendjson(w, r, err[any](e.Error()), statusCode)
}

// Отправка ответа с данными как JSON
func SendJSON(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	sendjson(w, r, ok(data), statusCode)
}
