package httpd

import (
	"fmt"
	"net/http"

	key_store "github.com/sarthakvk/gokey/internal/key_store"
	"github.com/sarthakvk/gokey/internal/logging"
)

var (
	Keystore *key_store.KeyStore
	logger   = logging.GetHttpdLogger()
)

func RunServer(store *key_store.KeyStore, port int) {
	Keystore = store
	for _, url := range urls {
		http.HandleFunc(url.Pattern, url.Handler)
	}
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}
