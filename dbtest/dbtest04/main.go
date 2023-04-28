// ref: github.com/cockroachdb/pebble

package main

import (
	"encoding/binary"
	"log"

	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/indexvm/storage"

	"github.com/ava-labs/hypersdk/pebble"
)

func bytes2string(b []byte) string {
	return string(b)
}

func uint64ToBytes(val uint64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	// binary.LittleEndian.PutUint64(b, uint64(val))
	return
}
func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
	// return binary.LittleEndian.Uint64(b)
}
func main() {
	privateKeyFile := "../../.index-cli.pk"
	priv, err := crypto.LoadKey(privateKeyFile)
	if err != nil {
		log.Printf("LoadKey err: %v", err)
		return
	}
	priv.ToHex()
	log.Printf("priv: %v", priv)
	log.Printf("priv.ToHex(): %v", priv.ToHex())
	log.Printf("priv.PublicKey(): %v", priv.PublicKey())
	log.Printf("priv.PublicKey(): %s", priv.PublicKey())
	log.Printf("priv.PublicKey() (hex): %x", priv.PublicKey())

	key := storage.PrefixTestKey(priv.PublicKey())
	log.Printf("key: %v", key)
	log.Printf("=========")
	// ctx := context.Background()

	// db, err := pebble.Open("demo", &pebble.Options{})
	cfg := pebble.NewDefaultConfig()
	db, err := pebble.New("demo", cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	b, err := db.Has(key)
	if b == false || err != nil {
		log.Printf("db.Has err: %v", err)

		if err := db.Put(key, uint64ToBytes(0)); err != nil {
			log.Fatal(err)
		}
	}
	value, err := db.Get(key)
	log.Printf("bytesToUint64(value) + 1 : %d ", bytesToUint64(value)+1)

	db.Put(key, uint64ToBytes(bytesToUint64(value)+1))

}
