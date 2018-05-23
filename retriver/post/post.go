package post

import "net/http"

type Poster struct {
	UserAgent    string
	LastContents string
}

func (post Poster) Post(url string) string {
	http.Post()
}
