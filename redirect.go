package redirect

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/nbari/violetear"
)

type Redirect struct {
	db *bolt.DB
}

func New(c, d string) (*Redirect, error) {
	config, err := os.Open(c)
	if err != nil {
		return nil, fmt.Errorf("Could not open config file %q: %s", c, err)
	}

	db, err := bolt.Open(d, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not open db file %q: %s", d, err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("302"))
		if err != nil {
			return fmt.Errorf("Error creating db bucket: %s", err)
		}
		return nil
	})

	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)

	var lineNum int = 1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Bad format line in config %q line %d", c, lineNum)
		}
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("302"))
			err := b.Put(
				[]byte(strings.TrimSpace(parts[0])),
				[]byte(strings.TrimSpace(parts[1])),
			)
			return err
		})
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading from config file %q: %s", c, err)
	}

	return &Redirect{
		db,
	}, nil
}

func (r *Redirect) Start(p int) error {
	router := violetear.New()
	router.LogRequests = true

	router.HandleFunc("*", r.catchAll)
	return http.ListenAndServe(fmt.Sprintf(":%d", p), router)
}
