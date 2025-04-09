package ont

import "net/http"

type Session struct {
	*http.Client
	Endpoint string
}
