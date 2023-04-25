package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/crypto"

	"github.com/ava-labs/indexvm/utils"
)

func SetDataContent(
	ctx context.Context,
	db chain.Database,
	contentID ids.ID,
	pk crypto.PublicKey,
	// royalty uint64,
) error {
	log.Printf("SetDataContent: contentID=%s", contentID)
	// if royalty == 0 {
	// 	return ErrInsufficientTip
	// }
	owner, _, exists, err := GetContent(ctx, db, contentID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf(
			// "%w: owner=%s royalty=%d",
			"%w: owner=%s",
			ErrContentAlreadyExists,
			utils.Address(owner),
			// royalty,
		)
	}
	// return SetContent(ctx, db, contentID, pk, royalty)
	return SetData(ctx, db, contentID, pk)
}

func SetData(
	ctx context.Context,
	db chain.Database,
	contentID ids.ID,
	owner crypto.PublicKey,
	// royalty uint64,
) error {
	log.Printf("SetData: contentID=%s, owner=%v", contentID, owner)
	v := make([]byte, 40)
	copy(v, owner[:])
	// binary.BigEndian.PutUint64(v[32:], royalty)
	return db.Insert(ctx, PrefixContentKey(contentID), v)
}

func AddDataContent(
	ctx context.Context,
	db chain.Database,
	contentID ids.ID,
	pk crypto.PublicKey,
) error {
	log.Printf("AddDataContent: contentID=%s", contentID)
	owner, _, exists, err := GetContent(ctx, db, contentID)
	log.Printf("AddDataContent: owner=%v, exists=%v, err=%v", owner, exists, err)
	if err != nil {
		log.Printf("AddDataContent: err=%v", err)
		return err
	}
	if exists {
		return fmt.Errorf(
			// "%w: owner=%s royalty=%d",
			"%w: owner=%s",
			ErrContentAlreadyExists,
			utils.Address(owner),
			// royalty,
		)
	}
	// return SetData(ctx, db, contentID, pk)
	return nil
}
