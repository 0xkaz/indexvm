// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package actions

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/hypersdk/utils"
	"github.com/ava-labs/indexvm/auth"
	"github.com/ava-labs/indexvm/genesis"
	"github.com/ava-labs/indexvm/storage"
)

var _ chain.Action = (*Burn2)(nil)

type Burn2 struct {
	// To is the recipient of the [Value].
	To crypto.PublicKey `json:"to"`

	// Amount are transferred to [To].
	Value uint64 `json:"value"`
}

func (t *Burn2) StateKeys(rauth chain.Auth) [][]byte {
	return [][]byte{
		storage.PrefixBalanceKey(auth.GetActor(rauth)),
		storage.PrefixBalanceKey(t.To),
		// TODO: Make conditional if account already exists
		storage.PrefixPermissionsKey(t.To, t.To),
	}
}

func (t *Burn2) Execute(
	ctx context.Context,
	r chain.Rules,
	db chain.Database,
	_ int64,
	rauth chain.Auth,
	_ ids.ID,
) (*chain.Result, error) {
	actor := auth.GetActor(rauth)
	unitsUsed := t.MaxUnits(r) // max units == units
	if t.Value == 0 {
		return &chain.Result{Success: false, Units: unitsUsed, Output: OutputValueZero}, nil
	}
	stateLockup, err := genesis.GetStateLockup(r)
	if err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	if err := storage.SubUnlockedBalance(ctx, db, actor, t.Value); err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	alreadyExists, err := storage.AddUnlockedBalance(ctx, db, t.To, t.Value, false)
	if err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	if alreadyExists {
		return &chain.Result{Success: true, Units: unitsUsed}, nil
	}
	// new accounts must lock funds for balance and perms
	if err := storage.LockBalance(ctx, db, t.To, stateLockup*2); err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	// new accounts have default perms
	if err := storage.SetPermissions(ctx, db, t.To, t.To, consts.MaxUint8, consts.MaxUint8); err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	return &chain.Result{Success: true, Units: unitsUsed}, nil
}

func (*Burn2) MaxUnits(chain.Rules) uint64 {
	return crypto.PublicKeyLen + consts.Uint64Len
}

func (t *Burn2) Marshal(p *codec.Packer) {
	p.PackPublicKey(t.To)
	p.PackUint64(t.Value)
}

func UnmarshalBurn2(p *codec.Packer) (chain.Action, error) {
	var bb Burn2
	p.UnpackPublicKey(&bb.To)
	bb.Value = p.UnpackUint64(true) // use [Clear] to empty
	return &bb, p.Err()
}

func (*Burn2) ValidRange(chain.Rules) (int64, int64) {
	return -1, -1
}
