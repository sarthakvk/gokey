// This file defines the global level objects for the httpd service
// It is also intended to store configurations for the httpd service

package httpd

import (
	"fmt"
	"net/http"

	"github.com/sarthakvk/gokey/keystore"
	"github.com/sarthakvk/gokey/logging"
)

var (
	Keystore *keystore.KeyStore
	logger   = logging.GetHttpdLogger()
)

func RunServer(store *keystore.KeyStore, port int) {
	Keystore = store
	for _, url := range urls {
		http.HandleFunc(url.Pattern, url.Handler)
	}
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}
