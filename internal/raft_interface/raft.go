package raft

import (
	"time"

	raft_lib "github.com/hashicorp/raft"

	"github.com/sarthakvk/gokey/internal/key_store"
	"github.com/sarthakvk/gokey/internal/logging"
)

const (
	ElectionWaitDuration    = time.Duration(time.Millisecond * 10)
	AddVoterTimeoutDuration = time.Duration(time.Millisecond * 500)
)

var logger = logging.GetLogger()

type Raft struct {
	raftID    raft_lib.ServerID
	node      *raft_lib.Raft
	transport *raft_lib.NetworkTransport
}

func (r *Raft) IsLeader() bool {
	raftFuture := r.node.VerifyLeader()

	return raftFuture.Error() == nil
}

// GetLeader will return the Address of the current leader
// If there is no leader, it will wait till new leader is elected.
func (r *Raft) GetLeader() raft_lib.ServerAddress {
	leader := raft_lib.ServerAddress("")
	for {
		leader, _ = r.node.LeaderWithID()
		if leader == "" {
			time.Sleep(ElectionWaitDuration)
		} else {
			break
		}
	}
	return leader
}

// Creates a new raft node
// TODO : Currently using default config, provide a way to configure the Raft node
func NewRaft(raftID, address string) *Raft {
	config := raft_lib.DefaultConfig()
	config.LocalID = raft_lib.ServerID(raftID)

	fsm := keystore.KeyStoreFSM{}
	transport := NewTransport(address)

	node, err := raft_lib.NewRaft(config, fsm, LogStore, StableStore, SnapshotStore, transport)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return &Raft{raftID: raft_lib.ServerID(raftID), node: node, transport: transport}
}

// Get latest configuration, panic if error
func (r *Raft) getLatestConfiguration() raft_lib.Configuration {
	configFuture := r.node.GetConfiguration()
	err := configFuture.Error()

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	return configFuture.Configuration()
}

// Bootstrap Cluster
func (r *Raft) BootstrapCluster() {
	config := raft_lib.Configuration{
		Servers: []raft_lib.Server{{
			Suffrage: raft_lib.Voter,
			Address:  r.transport.LocalAddr(),
			ID:       r.raftID,
		}},
	}

	fut := r.node.BootstrapCluster(config)
	err := fut.Error()

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}

func (r *Raft) AddVoter(raftID, address string) {
	serverID := raft_lib.ServerID(raftID)
	serverAddress := raft_lib.ServerAddress(address)
	prevIndex := uint64(0)

	future := r.node.AddVoter(serverID, serverAddress, prevIndex, AddVoterTimeoutDuration)
	err := future.Error()

	if err != nil {
		logger.Debug("Failed to add voter raft node!")
		logger.Error(err.Error())
		return
	}

}

func (r *Raft) Apply(cmd)
