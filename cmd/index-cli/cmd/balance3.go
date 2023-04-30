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

var balance3Cmd = &cobra.Command{
	Use:   "balance3 [options]",
	Short: "balance3 ",
	RunE:  balanceFunc,
}

func balance3Func(_ *cobra.Command, args []string) error {
	log.Printf("balance3Func")
	priv, err := crypto.LoadKey(privateKeyFile)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}

	// addr, err := getBalance3Op(args)
	err = getBalance3Op(args)
	if err != nil {
		return err
	}
	// log.Printf("addr: %v", addr)

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
	u, l, e, err := cli.Balance(ctx, utils.Address(priv.PublicKey()))

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

// func getBalance3Op(args []string) (addr crypto.PublicKey, err error) {
func getBalance3Op(args []string) (err error) {
	log.Printf("args: %v", args)
	if len(args) != 0 {
		// return crypto.EmptyPublicKey, fmt.Errorf(
		return fmt.Errorf(
			"expected exactly 0 arguments, got %d",
			len(args),
		)
	}
	// addr, err = utils.ParseAddress(args[0])
	// if err != nil {
	// 	return crypto.EmptyPublicKey, fmt.Errorf("%w: failed to parse address %s", err, args[0])
	// }
	// // value, err = hutils.ParseBalance(args[1])
	// // if err != nil {
	// // 	return crypto.EmptyPublicKey, fmt.Errorf("%w: failed to parse %s", err, args[1])
	// // }
	// // return addr, value, nil
	// return addr, nil
	return nil
}
