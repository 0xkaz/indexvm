// ref: github.com/cockroachdb/pebble

package main

import (
	"encoding/binary"
	"log"

	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/indexvm/storage"

	// "github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble"
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
	// log.Printf("priv (hex): %x", priv)
	key := storage.PrefixTestKey(priv.PublicKey())
	log.Printf("key: %v", key)

	db, err := pebble.Open("demo", &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	value, closer, err := db.Get(key)
	if err != nil {
		log.Printf("db.Get err: %v", err)

		// if err := db.Set(key, []byte("aaa"), pebble.Sync); err != nil {
		if err := db.Set(key, uint64ToBytes(0), pebble.Sync); err != nil {
			log.Fatal(err)
		}
		value = uint64ToBytes(0)
	}
	value = uint64ToBytes(bytesToUint64(value) + 1)

	log.Printf("key=%x, val=%d\n", key, bytesToUint64(value))
	if err := closer.Close(); err != nil {
		log.Fatal(err)
	}

	if err := db.Set(key, value, pebble.Sync); err != nil {
		log.Fatal(err)
	}

	// if err := db.Close(); err != nil {
	// 	log.Fatal(err)
	// }

}
