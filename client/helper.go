// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"
	"log"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/client"
	"github.com/ava-labs/hypersdk/utils"
)

func (cli *Client) GenerateTransaction(
	ctx context.Context,
	action chain.Action,
	factory chain.AuthFactory,
	modifiers ...client.Modifier,
) (func(context.Context) error, *chain.Transaction, uint64, error) {
	// Gather chain metadata
	log.Printf("GenerateTransaction: action=%v", action)
	log.Printf("GenerateTransaction: factory=%v", factory)
	g, err := cli.Genesis(ctx)
	if err != nil {
		log.Printf("GenerateTransaction: err=%v", err)
		return nil, nil, 0, err
	}
	log.Printf("g: %v", g)
	_, _, chainID, err := cli.Network(ctx) // TODO: store in object to fetch less frequently
	if err != nil {
		log.Printf("GenerateTransaction: err=%v", err)
		return nil, nil, 0, err
	}
	log.Printf("chainID: %v", chainID)
	return cli.Client.GenerateTransaction(
		ctx,
		&Parser{chainID, g},
		action,
		factory,
		modifiers...)
}

func (cli *Client) WaitForBalance(ctx context.Context, addr string, min uint64) error {
	return client.Wait(ctx, func(ctx context.Context) (bool, error) {
		unlocked, _, _, err := cli.Balance(ctx, addr)
		if err != nil {
			return false, err
		}
		shouldExit := unlocked >= min
		if !shouldExit {
			utils.Outf(
				"{{yellow}}waiting for %s balance: %s{{/}}\n",
				utils.FormatBalance(min),
				addr,
			)
		}
		return shouldExit, nil
	})
}

func (cli *Client) WaitForTransaction(ctx context.Context, txID ids.ID) error {
	return client.Wait(ctx, func(ctx context.Context) (bool, error) {
		_, accepted, err := cli.GetTx(ctx, txID)
		if err != nil {
			return false, err
		}
		return accepted, nil
	})
}
