// ref: github.com/cockroachdb/pebble

package main

import (
	"context"
	"encoding/binary"
	"log"

	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/indexvm/storage"
)

// func bytes2string(b []byte) string {
// 	return string(b)
// }

func uint64ToBytes(val uint64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	// binary.LittleEndian.PutUint64(b, uint64(val))
	return
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func main() {

	db, err := NewMyDatabase()
	defer db.Close()
	privateKeyFile := "../../.index-cli.pk"
	priv, err := crypto.LoadKey(privateKeyFile)
	if err != nil {
		log.Printf("LoadKey err: %v", err)
		return
	}
	priv.ToHex()
	log.Printf("priv.PublicKey() (hex): %x", priv.PublicKey())

	key := storage.PrefixTestKey(priv.PublicKey())
	log.Printf("key: %v", key)
	log.Printf("=========")
	ctx := context.TODO()
	ret, err := storage.IncBurn2Count(ctx, db, priv.PublicKey(), 12)

	log.Printf("ret: %v", ret)

}
