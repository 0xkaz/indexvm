package storage

import (
	"context"
	"crypto/ed25519"
	"encoding/binary"
	"log"

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
) error {
	log.Printf("IncBurn2Count: pk=%v, val=%d", pk, val)
	k := PrefixTestKey(pk)
	log.Printf("k=%v", k)
	v, err := db.GetValue(ctx, k)
	// if errors.Is(err, database.ErrNotFound) {
	// 	return nil
	// }
	log.Printf("v=%v", v)

	if err != nil {
		return err
	}

	// db.Insert(ctx, k, []byte{val})
	db.Insert(ctx, k, uint64toBytes(val))

	// return incBurn2Count(ctx, db, contentID, pk)
	return nil
}
