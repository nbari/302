package redirect

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
)

func (r *Redirect) catchAll(w http.ResponseWriter, req *http.Request) {
	host := strings.Split(req.Host, ":")[0]
	if req.RequestURI == "/" {
		req.RequestURI = ""
	}
	query := fmt.Sprintf("%s%s", host, req.RequestURI)
	r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("302"))
		v := b.Get([]byte(query))
		if v != nil {
			http.Redirect(w, req, string(v), 302)
		}
		return nil
	})
}
