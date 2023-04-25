package actions

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
	"github.com/ava-labs/hypersdk/utils"
	"github.com/ava-labs/indexvm/auth"
	"github.com/ava-labs/indexvm/storage"
)

var _ chain.Action = (*AddData)(nil)

type AddData struct {
	Content ids.ID `json:"content"`
}

func (s *AddData) StateKeys(rauth chain.Auth) [][]byte {
	actor := auth.GetActor(rauth)
	return [][]byte{storage.PrefixBalanceKey(actor), storage.PrefixContentKey(s.Content)}
}

func (s *AddData) Execute(
	ctx context.Context,
	r chain.Rules,
	db chain.Database,
	_ int64,
	rauth chain.Auth,
	_ ids.ID,
) (*chain.Result, error) {
	actor := auth.GetActor(rauth)
	unitsUsed := s.MaxUnits(r) // max units == units

	if s.Content == ids.Empty {
		return &chain.Result{Success: false, Units: unitsUsed, Output: OutputInvalidContent}, nil
	}

	if err := storage.AddDataContent(ctx, db, s.Content, actor); err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	return &chain.Result{Success: true, Units: unitsUsed}, nil
}

func (*AddData) MaxUnits(chain.Rules) uint64 {
	return consts.IDLen
}

func (u *AddData) Marshal(p *codec.Packer) {
	p.PackID(u.Content)
}

func UnmarshalAddData(p *codec.Packer) (chain.Action, error) {
	var setdata AddData
	p.UnpackID(true, &setdata.Content)
	return &setdata, p.Err()
}

func (*AddData) ValidRange(chain.Rules) (int64, int64) {
	return -1, -1
}
