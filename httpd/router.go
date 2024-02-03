// This file defines the routes in a slice for the applications, the design is inspired from Django Framework.

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
		{
			Pattern: "/add-replica",
			Handler: AddReplicaHandler,
		},
	}
)
