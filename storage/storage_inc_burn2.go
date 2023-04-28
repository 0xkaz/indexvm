package storage

import (
	"context"
	"log"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/crypto"
)

func IncBurn2Count(
	ctx context.Context,
	db chain.Database,
	contentID ids.ID,
	pk crypto.PublicKey,
	// royalty uint64,
) error {
	log.Printf("IncBurn2Count: contentID=%s", contentID)
	// // if royalty == 0 {
	// // 	return ErrInsufficientTip
	// // }
	// owner, _, exists, err := GetContent(ctx, db, contentID)
	// if err != nil {
	// 	return err
	// }
	// if exists {
	// 	return fmt.Errorf(
	// 		// "%w: owner=%s royalty=%d",
	// 		"%w: owner=%s",
	// 		ErrContentAlreadyExists,
	// 		utils.Address(owner),
	// 		// royalty,
	// 	)
	// }
	// return SetContent(ctx, db, contentID, pk, royalty)
	return incBurn2Count(ctx, db, contentID, pk)

}

func incBurn2Count(
	ctx context.Context,
	db chain.Database,
	contentID ids.ID,
	owner crypto.PublicKey,
	// royalty uint64,
) error {
	log.Printf("incBurn2Count: contentID=%s, owner=%v", contentID, owner)
	v := make([]byte, 40)
	copy(v, owner[:])
	// binary.BigEndian.PutUint64(v[32:], royalty)
	return db.Insert(ctx, PrefixContentKey(contentID), v)
}
