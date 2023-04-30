// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package auth

import (
	"context"

	"log"

	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/indexvm/consts"
	"github.com/ava-labs/indexvm/storage"
	"github.com/ava-labs/indexvm/utils"
)

func GetActor(auth chain.Auth) crypto.PublicKey {
	log.Printf("GetActor: %T", auth)
	switch a := auth.(type) {
	case *ED25519:
		log.Printf("GetActor ED25519")
		return a.Signer
	case *Direct:
		log.Printf("GetActor Direct")
		return a.Signer
	case *Delegate:
		log.Printf("GetActor Delegate")
		return a.Actor
	default:
		log.Printf("GetActor else")
		return crypto.EmptyPublicKey
	}
}

func GetSigner(auth chain.Auth) crypto.PublicKey {
	log.Printf("GetSigner: %T", auth)
	switch a := auth.(type) {
	case *ED25519:
		log.Printf("GetSigner ED25519")
		return a.Signer
	case *Direct:
		log.Printf("GetSigner Direct")
		return a.Signer
	case *Delegate:
		log.Printf("GetSigner Signer")
		return a.Signer
	default:
		log.Printf("GetSigner else")
		return crypto.EmptyPublicKey
	}
}

func Authorized(
	ctx context.Context,
	db chain.Database,
	action chain.Action,
	actor crypto.PublicKey,
	signer crypto.PublicKey,
	actorPays bool,
) error {
	log.Printf("Authorized: action=%T", action)

	log.Printf("Authorized: actor=%v", actor)
	log.Printf("Authorized: signer=%v", signer)
	log.Printf("Authorized: actorPays=%v", actorPays)

	actionPerms, miscPerms, err := storage.GetPermissions(ctx, db, actor, signer)
	if err != nil {
		return err
	}

	log.Printf("Authorized: miscPerms: %d", miscPerms)
	log.Printf("Authorized: actionPerms: %d", actionPerms)

	index, _, exists := consts.ActionRegistry.LookupType(action)
	if !exists {
		log.Printf("Authorized: consts.ActionRegistry.LookupType. exists==false ")
		return ErrActionMissing
	}

	log.Printf("Authorized: index: %d", index)

	if !utils.CheckBit(actionPerms, index) {
		log.Printf("Authorized: utils.CheckBit == false ")
		return ErrNotAllowed
	}
	if actorPays && !utils.CheckBit(miscPerms, actorPaysBit) {
		log.Printf("!utils.CheckBit")
		log.Printf("Authorized: miscPerms: %d", miscPerms)
		log.Printf("Authorized: actorPaysBit: %d", actorPaysBit)
		// log.Printf("Authorized: actorPays: %v", actorPays)

		return ErrNotAllowed
	}
	return nil
}
