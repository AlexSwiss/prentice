package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	// Shared common fields.
	Code    int               `json:"-"`
	Headers map[string]string `json:"-"`
	Action  string            `json:"action,omitempty"`

	// Success specific fields.
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`

	// Failure specific fields.
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func Write(w http.ResponseWriter, r Response) error {
	if r.Code == 0 {
		return errors.New("0 is not a valid code")
	}

	for k, v := range r.Headers {
		w.Header().Add(k, v)
	}

	if !isBodyAllowed(r.Code) {
		w.WriteHeader(r.Code)
		return nil
	}

	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.Code)

	if _, err := w.Write(body); err != nil {
		return err
	}

	return nil
}

// See RFC 7230, section 3.3.
func isBodyAllowed(status int) bool {
	if (status >= 100 && status <= 199) || status == 204 || status == 304 {
		return false
	}

	return true
}
