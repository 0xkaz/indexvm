// ref: github.com/cockroachdb/pebble

package main

import (
	"log"

	"github.com/ava-labs/hypersdk/pebble"
)

// import 	"github.com/ava-labs/hypersdk/chain"
// import 	"github.com/ava-labs/avalanchego/database"

func init() {

}

// var metaDB database.Database
func main() {
	cfg := pebble.NewDefaultConfig()
	blockPath := "/tmp/test-block"
	db, err := pebble.New(blockPath, cfg)
	if err != nil {
		log.Printf("db err: %v", err)
	}

	val3, err := db.Get([]byte("key"))
	if err != nil {
		log.Printf("db.Get err: %v", err)
	}
	log.Printf("val3: %s", val3)

	err = db.Put([]byte("key"), []byte("value"))
	if err != nil {
		log.Printf("db.Put err: %v", err)
	}

	val, err := db.Get([]byte("key"))
	if err != nil {
		log.Printf("db.Get err: %v", err)
	}
	log.Printf("val: %s", val)

	err = db.Put([]byte("key"), []byte("value222333322"))
	if err != nil {
		log.Printf("db.Put err: %v", err)
	}

	val2, err := db.Get([]byte("key"))
	if err != nil {
		log.Printf("db.Get err: %v", err)
	}
	log.Printf("val2: %s", val2)
	db.Close()
}
