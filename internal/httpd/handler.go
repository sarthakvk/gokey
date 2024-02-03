package httpd

import (
	"net/http"

	store "github.com/sarthakvk/gokey/internal/key_store"
)

// Handler for the KeyStore Commands, this is used to perform
// Operations on our key store
func HandleKeyStoreCommand(w http.ResponseWriter, req *http.Request) {
	cmd, err := ValidateKeyStoreCommand(req)

	if err != nil {
		SendResponse(w, 400, err.Error())
		return
	}

	switch cmd.Operation {
	case store.GET:
		value, ok := Keystore.Get(cmd.Key)

		if !ok {
			SendResponse(w, http.StatusNotFound, "Not found")
		} else {
			SendKeyStoreCommandResponse(w, value)
		}

	case store.DELETE:
		err := Keystore.Delete(cmd.Key)

		if err != nil {
			logger.Debug(err.Error())
			SendResponse(w, http.StatusUnprocessableEntity, err.Error())
		} else {
			SendKeyStoreCommandResponse(w, "")
		}

	case store.SET:
		err := Keystore.Set(cmd.Key, cmd.Value)

		if err != nil {
			logger.Debug(err.Error())
			SendResponse(w, 401, err.Error())
		} else {
			SendKeyStoreCommandResponse(w, "", true)
		}

	case store.GET_OR_CREATE:
		created, value, err := Keystore.GetOrCreate(cmd.Key, cmd.Value)

		if err != nil {
			logger.Debug(err.Error())
			SendResponse(w, http.StatusUnprocessableEntity, err.Error())
		} else {
			SendKeyStoreCommandResponse(w, value, created)
		}
	default:
		SendResponse(w, http.StatusBadRequest, "Invalid command")
	}
}

// Handler requests for adding more nodes to the cluster
func AddReplicaHandler(w http.ResponseWriter, req *http.Request) {
	validated_data, err := ValidateReplicationRequest(req)

	if err != nil {
		SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	nodeID, addr := validated_data.NodeID, validated_data.Address
	Keystore.Replicate(nodeID, addr)

	SendResponse(w, http.StatusOK, "Replication started")
}
