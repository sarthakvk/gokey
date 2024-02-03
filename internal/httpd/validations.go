package httpd

import (
	"encoding/json"
	"errors"
	"net/http"

	key_store "github.com/sarthakvk/gokey/internal/key_store"
)

func ValidateKeyStoreCommand(request *http.Request) (*key_store.Command, error) {
	if request.Method != http.MethodPost {
		err := errors.New("bad request")
		return nil, err
	}

	var cmd key_store.Command
	body := request.Body

	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}

	return &cmd, nil
}

func ValidateReplicationRequest(request *http.Request) (*ReplicationRequest, error) {
	if request.Method != http.MethodPost {
		err := errors.New("bad request")
		return nil, err
	}

	var req ReplicationRequest
	body := request.Body

	err := json.NewDecoder(body).Decode(&req)

	if err != nil {
		logger.Debug(err.Error())
		return nil, err
	}

	return &req, nil
}
