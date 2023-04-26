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

var _ chain.Action = (*SetData)(nil)

type SetData struct {
	Content ids.ID `json:"content"`
}

func (s *SetData) StateKeys(rauth chain.Auth) [][]byte {
	actor := auth.GetActor(rauth)
	return [][]byte{
		storage.PrefixBalanceKey(actor),
		storage.PrefixContentKey(s.Content),
	}
	// keys := [][]byte{storage.PrefixBalanceKey(actor)}
	// if i.Parent != ids.Empty {
	// 	keys = append(keys, storage.PrefixContentKey(s.Parent))
	// 	if i.Searcher != crypto.EmptyPublicKey {
	// 		keys = append(keys, storage.PrefixBalanceKey(i.Searcher))
	// 	}
	// }
	// if i.Royalty > 0 {
	// 	keys = append(keys, storage.PrefixContentKey(i.ContentID()))
	// }
	// if i.Servicer != crypto.EmptyPublicKey && i.Commission > 0 {
	// 	// You can be serviced with or without a [Parent]
	// 	keys = append(keys, storage.PrefixBalanceKey(i.Servicer))
	// }
	// return keys
}

func (s *SetData) Execute(
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
	// if s.Schema == ids.Empty {
	// 	return &chain.Result{Success: false, Units: unitsUsed, Output: OutputInvalidSchema}, nil
	// }

	if err := storage.SetDataContent(ctx, db, s.Content, actor); err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	// stateLockup, err := genesis.GetStateLockup(r)
	// if err != nil {
	// 	return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	// }
	// if err := storage.UnlockBalance(ctx, db, actor, stateLockup); err != nil {
	// 	return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	// }
	return &chain.Result{Success: true, Units: unitsUsed}, nil
}

func (*SetData) MaxUnits(chain.Rules) uint64 {
	return consts.IDLen
}

func (u *SetData) Marshal(p *codec.Packer) {
	p.PackID(u.Content)
}

func UnmarshalSetData(p *codec.Packer) (chain.Action, error) {
	var setdata SetData
	p.UnpackID(true, &setdata.Content)
	return &setdata, p.Err()
}

func (*SetData) ValidRange(chain.Rules) (int64, int64) {
	return -1, -1
}
