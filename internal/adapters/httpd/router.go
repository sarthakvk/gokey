package httpd

import "net/http"

type Url struct {
	Pattern string
	Handler func(http.ResponseWriter, *http.Request)
}

var (
	urls = []Url{
		{
			Pattern: "/key-store",
			Handler: HandleKeyStoreCommand,
		},
	}
)
