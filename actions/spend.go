// just use units
package actions

import (
	"context"
	"encoding/binary"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
	"github.com/ava-labs/indexvm/auth"
	"github.com/ava-labs/indexvm/storage"
)

var _ chain.Action = (*Spend)(nil)

type Spend struct {
	Amount uint64 `json:"amount"`
	// Content ids.ID `json:"content"`
}

func (s *Spend) StateKeys(rauth chain.Auth) [][]byte {
	actor := auth.GetActor(rauth)
	// return [][]byte{storage.PrefixBalanceKey(actor), storage.PrefixContentKey(s.Content)}
	b := make([]byte, 32)
	binary.LittleEndian.PutUint64(b, uint64(s.Amount))
	return [][]byte{storage.PrefixBalanceKey(actor), b}
}

func (s *Spend) Execute(
	ctx context.Context,
	r chain.Rules,
	db chain.Database,
	_ int64,
	_ chain.Auth,
	// rauth chain.Auth,
	_ ids.ID,
) (*chain.Result, error) {
	// actor := auth.GetActor(rauth)
	// unitsUsed := s.MaxUnits(r) // max units == units

	// if s.Content == ids.Empty {
	// 	// use full units even if rollback.
	// 	return &chain.Result{Success: false, Units: unitsUsed, Output: OutputInvalidContent}, nil
	// }

	return &chain.Result{Success: true, Units: s.Amount}, nil
	// return &chain.Result{Success: true, Units: unitsUsed}, nil
}

func (*Spend) MaxUnits(chain.Rules) uint64 {
	return consts.IDLen // 32
}

func (u *Spend) Marshal(p *codec.Packer) {
	// p.PackID(u.Content)
}

func UnmarshalSpend(p *codec.Packer) (chain.Action, error) {
	var s Spend
	// p.UnpackID(true, &s.Content)
	return &s, p.Err()
}

func (*Spend) ValidRange(chain.Rules) (int64, int64) {
	return -1, -1
}
