package keystore

import (
	raft_lib "github.com/hashicorp/raft"
	"github.com/sarthakvk/gokey/internal/logging"
)

var (
	logger = logging.GetLogger()
)

type KeyStoreFSM struct {
	raft_lib.FSM
}

func (fsm *KeyStoreFSM) Applt(log *raft_lib.Log) {
	logger.Debug("Append entry request")
	logger.Debug(log.Type.String())
}
