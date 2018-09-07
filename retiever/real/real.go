package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type R struct {
	UserAgent string
	Timeout   time.Duration
}

func (r *R) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(result)
}
