package httpd

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sarthakvk/gokey/keystore"
)

func ValidateKeyStoreCommand(request *http.Request) (*keystore.Command, error) {
	if request.Method != http.MethodPost {
		err := errors.New("bad request")
		return nil, err
	}

	var cmd keystore.Command
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
