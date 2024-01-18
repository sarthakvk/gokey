package main

import (
	"flag"
	"fmt"

	raft "github.com/sarthakvk/gokey/internal/raft_interface"
)

var (
	bootstrap = flag.Bool("bootstrap", false, "Bootstrap Cluster")
	addr      = flag.String("address", "", "address of the node")
	raftID    = flag.String("node-id", "", "unique id for node")
	peerCount = flag.Int("peer-count", 0, "Number of peers to add; this only makes sense when bootstraping a cluster")
)

func main() {
	flag.Parse()

	node := raft.NewRaft(*raftID, *addr)

	if *bootstrap {
		node.BootstrapCluster()
		addPeers(node, *peerCount)
	}
	node.Run()
}

func addPeers(r *raft.Raft, count int) {
	for i := 0; i < count; i++ {
		var raftID, address string
		println("Enter Peer: ", i+1)
		fmt.Scan(&raftID, &address)
		r.AddVoter(raftID, address)
	}
}
