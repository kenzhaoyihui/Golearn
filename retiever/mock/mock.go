package mock

import "fmt"

type M struct {
	Contents string
}

func (r *M) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
func (r *M) Get(url string) string {
	return r.Contents
}

func (r *M) String() string {
	return fmt.Sprintf(
		"Retiever: {Contents=%s}", r.Contents,
	)
}
