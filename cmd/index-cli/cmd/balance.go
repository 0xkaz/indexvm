package cmd

import (
	"context"
	"log"

	"fmt"

	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/indexvm/client"
	"github.com/ava-labs/indexvm/utils"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance [options]",
	Short: "balance ",
	RunE:  balanceFunc,
}

func balanceFunc(_ *cobra.Command, args []string) error {
	log.Printf("balanceFunc")

	addr, err := getBalanceOp(args)
	if err != nil {
		return err
	}
	log.Printf("addr: %v", addr)

	ctx := context.Background()
	// cli := client.New(uri)
	// c := controller.New()
	cli := client.New(uri)
	networkID, subnetID, chainID, err := cli.Network(context.Background())
	if err != nil {
		return err
	}
	log.Printf("networkID=%d subnetID=%s chainID=%s", networkID, subnetID, chainID)
	// u, l, e, err := cli.Balance(ctx, fmt.Sprintf("%v", addr))
	u, l, e, err := cli.Balance(ctx, args[0])

	// u, l, err := storage.GetBalanceFromState(ctx, c.inner.ReadState, addr)
	if err != nil {
		log.Printf("cli.Balance err=%v", err)
		return err
	}
	log.Printf("u: %v, l: %v, e: %v ", u, l, e)
	// port, err := cli.BlocksPort(ctx)
	// if err != nil {
	// 	return err
	// }
	// host, err := utils.GetHost(uri)
	// if err != nil {
	// 	return err
	// }
	// scli, err := vm.NewBlockRPCClient(fmt.Sprintf("%s:%d", host, port))
	// if err != nil {
	// 	return err
	// }
	// defer scli.Close()
	// parser, err := cli.Parser(ctx)
	// if err != nil {
	// 	return err
	// }
	// totalTxs := float64(0)
	// start := time.Now()
	// utils.Outf("{{green}}watching for new blocks 👀{{/}}\n")
	// for ctx.Err() == nil {
	// 	blk, _, err := scli.Listen(parser)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	totalTxs += float64(len(blk.Txs))
	// 	utils.Outf(
	// 		"{{yellow}}height:{{/}}%d {{yellow}}txs:{{/}}%d {{yellow}}units:{{/}}%d {{yellow}}root:{{/}}%s {{yellow}}avg TPS:{{/}}%f\n", //nolint:lll
	// 		blk.Hght,
	// 		len(blk.Txs),
	// 		blk.UnitsConsumed,
	// 		blk.StateRoot,
	// 		totalTxs/time.Since(start).Seconds(),
	// 	)
	// }
	// color.Green("balance of %s: %d", addr, amt)

	return nil
}

func getBalanceOp(args []string) (addr crypto.PublicKey, err error) {
	if len(args) != 1 {
		return crypto.EmptyPublicKey, fmt.Errorf(
			"expected exactly 1 arguments, got %d",
			len(args),
		)
	}

	addr, err = utils.ParseAddress(args[0])
	if err != nil {
		return crypto.EmptyPublicKey, fmt.Errorf("%w: failed to parse address %s", err, args[0])
	}
	// value, err = hutils.ParseBalance(args[1])
	// if err != nil {
	// 	return crypto.EmptyPublicKey, fmt.Errorf("%w: failed to parse %s", err, args[1])
	// }
	// return addr, value, nil
	return addr, nil
}
