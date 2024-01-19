package httpd

import (
	"net/http"

	store "github.com/sarthakvk/gokey/internal/key_store"
)

func HandleKeyStoreCommand(w http.ResponseWriter, req *http.Request) {
	cmd, err := ValidateKeyStoreCommand(req.Body)

	if err != nil {
		SendError(w, 400, err.Error())
		return
	}

	switch cmd.Operation {
	case store.GET:
		value, ok := Keystore.Get(cmd.Key)

		if !ok {
			logger.Debug(err.Error())
			SendError(w, http.StatusNotFound, "Not found")
		} else {
			SendKeyStoreCommandResponse(w, value)
		}

	case store.DELETE:
		err := Keystore.Delete(cmd.Key)

		if err != nil {
			logger.Debug(err.Error())
			SendError(w, http.StatusUnprocessableEntity, err.Error())
		} else {
			SendKeyStoreCommandResponse(w, "")
		}

	case store.SET:
		err := Keystore.Set(cmd.Key, cmd.Value)

		if err != nil {
			logger.Debug(err.Error())
			SendError(w, 401, err.Error())
		} else {
			SendKeyStoreCommandResponse(w, "", true)
		}

	case store.GET_OR_CREATE:
		created, value, err := Keystore.GetOrCreate(cmd.Key, cmd.Value)

		if err != nil {
			logger.Debug(err.Error())
			SendError(w, http.StatusUnprocessableEntity, err.Error())
		} else {
			SendKeyStoreCommandResponse(w, value, created)
		}
	}
}
