package main

import (
	"flag"

	"github.com/sarthakvk/gokey/httpd"
	keystore "github.com/sarthakvk/gokey/keystore"
)

func RunHttpDaemon(store *keystore.KeyStore, httpPort int) {
	httpd.RunServer(store, httpPort)
}

var (
	bootstrap = flag.Bool("bootstrap", false, "Bootstrap Cluster")
	addr      = flag.String("address", "", "address of the node")
	raftID    = flag.String("node-id", "", "unique id for node")
	httpPort  = flag.Int("http-port", 9000, "port to run http service")
)

func main() {
	flag.Parse()

	store := keystore.New(*raftID, *addr, *bootstrap)

	RunHttpDaemon(store, *httpPort)
}
