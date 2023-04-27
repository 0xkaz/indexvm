// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package controller

import (
	"net/http"

	"github.com/ava-labs/indexvm/storage"
	"github.com/ava-labs/indexvm/utils"
)

type BalanceArgs struct {
	Address string `json:"address"`
}

type BalanceReply struct {
	Exists   bool   `json:"exists"`
	Unlocked uint64 `json:"unlocked"`
	Locked   uint64 `json:"locked"`
	Amount   uint64 `json:"amount"` // Unlocked + Locked
}

func (h *Handler) Balance(req *http.Request, args *BalanceArgs, reply *BalanceReply) error {
	ctx, span := h.c.inner.Tracer().Start(req.Context(), "Handler.Balance")
	defer span.End()

	addr, err := utils.ParseAddress(args.Address)
	if err != nil {
		return err
	}
	u, l, err := storage.GetBalanceFromState(ctx, h.c.inner.ReadState, addr)
	if err != nil {
		return err
	}
	reply.Exists = l > 0
	reply.Unlocked = u
	reply.Locked = l
	reply.Amount = u + l
	return err
}
