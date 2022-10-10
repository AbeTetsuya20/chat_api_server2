package entity

import (
	"golang.org/x/xerrors"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// errorResponse エラーレスポンスを表す。
type errorResponse struct {
	// Message エラーメッセージ
	Message string `json:"message"`
}

// WriteErrorResponse w にエラーレスポンスを書き込む。
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	resp := &errorResponse{
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Printf("[ERROR] error response encoding failed: %+v\n", err)
	}
}

// writeHTTPError w にエラーレスポンスを書き込む。
// エラーメッセージは HTTP ステータスコードに対応した文字列になる。
func WriteHTTPError(w http.ResponseWriter, code int) {
	WriteErrorResponse(w, code, http.StatusText(code))
}

func NewError(err error) string {
	newError := xerrors.Errorf("error: %v", err)
	return fmt.Sprintf("%+v\n", newError)
}
