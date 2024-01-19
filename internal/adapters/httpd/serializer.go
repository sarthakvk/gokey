package httpd

import (
	"encoding/json"
	"io"
	"net/http"

	key_store "github.com/sarthakvk/gokey/internal/key_store"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	resp := ErrResponse{
		Code:    code,
		Message: message,
	}

	data, err := json.Marshal(resp)

	if err != nil {
		logger.Error(err.Error())
		return
	}
	w.Write(data)
}

type KeyStoreCommandResponse struct {
	// Used only incase of GET_OR_CREATE
	Created bool `json:"created,omitempty"`

	// will be omited incase of SET & DELETE
	Value string `json:"value,omitempty"`
}

func SendKeyStoreCommandResponse(w http.ResponseWriter, value string, created ...bool) {
	resp := KeyStoreCommandResponse{}

	if len(created) > 0 {
		resp.Created = created[0]
	}

	if len(value) > 0 {
		resp.Value = value
	}

	data, err := json.Marshal(resp)

	if err != nil {
		if err != nil {
			logger.Error(err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ValidateKeyStoreCommand(body io.ReadCloser) (*key_store.Command, error) {
	var cmd key_store.Command

	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}

	return &cmd, nil
}
