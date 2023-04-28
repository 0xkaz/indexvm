package storage

import (
	"context"
	"crypto/ed25519"
	"encoding/binary"
	"errors"
	"log"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/crypto"
)

// [testPrefix] + [signer]
func PrefixTestKey(signer crypto.PublicKey) (k []byte) {
	k = make([]byte, 1+ed25519.PublicKeySize)
	k[0] = testPrefix
	copy(k[1:], signer[:])
	return
}
func uint64toBytes(val uint64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	// binary.LittleEndian.PutUint64(b, uint64(val))
	return
}
func bytes2uint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
	// return binary.LittleEndian.Uint64(b)
}
func IncBurn2Count(
	ctx context.Context,
	db chain.Database,
	pk crypto.PublicKey,
	val uint64,
) (uint64, error) {
	log.Printf("IncBurn2Count: pk=%v, val=%d", pk, val)
	k := PrefixTestKey(pk)
	log.Printf("k=%v", k)
	v, err := db.GetValue(ctx, k)
	if errors.Is(err, database.ErrNotFound) {
		// return 0, nil
		v = uint64toBytes(0)
	}
	if err != nil {
		// return 0, err
		v = uint64toBytes(0)
	}

	log.Printf("v=%v", v)
	log.Printf("bytes2uint64(v)=%d", bytes2uint64(v))

	// db.Insert(ctx, k, []byte{val})
	err = db.Insert(ctx, k, uint64toBytes(val+bytes2uint64(v)))
	if err != nil {
		log.Printf("err=%v", err)
		return 0, err
	}
	// return incBurn2Count(ctx, db, contentID, pk)
	return val + bytes2uint64(v), nil
}
