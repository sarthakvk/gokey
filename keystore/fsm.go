package keystore

import (
	"encoding/json"
	"io"

	raft_lib "github.com/hashicorp/raft"
)

// KeyStoreFSM implements the raft.FSM interface for the KeyStore
//
// It is basically a KeyStore, but it's FSM implementation is extracted out for maintainablity purpose.
type KeyStoreFSM KeyStore

// Apply is called once a log entry is committed by a majority of the cluster.
//
// Apply should apply the log to the FSM. Apply must be deterministic and
// produce the same result on all peers in the cluster.
//
// The returned value is returned to the client as the ApplyFuture.Response.
func (store *KeyStoreFSM) Apply(log *raft_lib.Log) interface{} {

	cmd, err := GetCommand(log.Data)

	if err != nil {
		return err
	}

	switch cmd.Operation {
	case SET:
		store.applySet(cmd.Key, cmd.Value)
	case DELETE:
		store.applyDel(cmd.Key)
	}

	return nil
}

// Implements Snapshot, this helps raft to not have unbounded logs,
// And data can be recovered later on
func (store *KeyStoreFSM) Snapshot() (raft_lib.FSMSnapshot, error) {
	store.rw_lock.RLock()
	defer store.rw_lock.RUnlock()

	snap, err := newKeyStoreSnapshot(store.data)

	if err != nil {
		logger.Error("Failed to create snapshot")
		return nil, err
	}
	return snap, nil
}

// Implements Restore, It will restore the snapshopt back to the KeyStore
func (store *KeyStoreFSM) Restore(snapshot io.ReadCloser) error {
	store.rw_lock.Lock()
	defer store.rw_lock.Unlock()

	var (
		snap_data     []byte
		restored_data map[string]string
	)

	snapshot.Read(snap_data)

	err := json.Unmarshal(snap_data, &restored_data)
	if err != nil {
		logger.Error("Encountered erros while restoring snapshot data")
		return err
	}

	store.data = restored_data

	return nil
}

func (store *KeyStoreFSM) applySet(key, value string) {
	store.rw_lock.Lock()
	defer store.rw_lock.Unlock()

	store.data[key] = value
}

func (store *KeyStoreFSM) applyDel(key string) {
	store.rw_lock.Lock()
	defer store.rw_lock.Unlock()

	delete(store.data, key)
}
