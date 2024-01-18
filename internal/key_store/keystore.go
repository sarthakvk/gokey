package keystore

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sarthakvk/gokey/internal/logging"
	raft "github.com/sarthakvk/gokey/internal/raft_interface"
)

var (
	logger = logging.GetLogger()
)

type KeyStore struct {
	// Consesus mechanism
	raft *raft.Raft

	// Data storage for the key-value store
	data map[string]string

	rw_lock sync.RWMutex
}

func New(nodeID, address string, bootstrap_cluster bool) *KeyStore {
	store := &KeyStore{}
	fsm := (*KeyStoreFSM)(store)
	store.data = make(map[string]string)

	raft := raft.NewRaft(nodeID, address, fsm)
	store.raft = raft

	if bootstrap_cluster {
		go raft.BootstrapCluster()
	}

	return store
}

func (store *KeyStore) AddVoterNodes(nodeID, address string) {
	go store.raft.AddVoter(nodeID, address)
}

func (store *KeyStore) Get(key string) (string, bool) {
	store.rw_lock.Lock()
	defer store.rw_lock.Unlock()

	value, ok := store.data[key]

	return value, ok
}

func (store *KeyStore) Delete(key string) error {
	if !store.raft.IsLeader() {
		logger.Error("key deletion request denied, not a leader")
		return errors.New("request denied, Not a leader")
	}

	var cmd Command

	cmd.Operation = DELETE
	cmd.Key = key

	raw_cmd, err := getRawCommand(cmd)

	if err != nil {
		logger.Error("failed to marshal the delete request command")
		return errors.New("failed to marshal the delete request command")
	}

	fut := store.raft.Apply(raw_cmd)

	if err := fut.Error(); err != nil {
		logger.Error(err.Error())
		return err
	}

	resp := fut.Response()

	switch resp := resp.(type) {
	case error:
		return resp
	default:
		return nil
	}

}

func (store *KeyStore) Set(key, value string) error {
	if !store.raft.IsLeader() {
		logger.Error("key deletion request denied, Not a leader!")
		return errors.New("request denied, not a leader")
	}

	var cmd Command

	cmd.Operation = SET
	cmd.Key = key
	cmd.Value = value

	raw_cmd, err := getRawCommand(cmd)

	if err != nil {
		logger.Error("failed to marshal the delete request command")
		return errors.New("failed to marshal the delete request command")
	}

	fut := store.raft.Apply(raw_cmd)

	if err := fut.Error(); err != nil {
		logger.Error(err.Error())
		return err
	}

	resp := fut.Response()

	switch resp := resp.(type) {
	case error:
		return resp
	default:
		return nil

	}
}

func (store *KeyStore) GetOrCreate(key, value string) (bool, string, error) {
	created := false
	val, ok := store.Get(key)

	if ok {
		return created, val, nil
	}
	err := store.Set(key, value)

	if err == nil {
		created = true
	}

	return created, value, err

}

// TODO make interface to interact with store
// This is just for testing
func (store *KeyStore) Run() {
	print("\033[H\033[2J")
	print("\033[H\033[2J")
	println("Now it's time to execute commands")
	var cmd Operation
	var key, value string

	for {
		fmt.Scan(&cmd, &key, &value)
		switch cmd {
		case GET:
			out, _ := store.Get(key)
			println(out)
		case DELETE:
			out := store.Delete(key)
			println(out)

		case SET:
			out := store.Set(key, value)
			println(out)

		case GET_OR_CREATE:
			_, out, _ := store.GetOrCreate(key, value)
			println(out)

		}
	}
}
