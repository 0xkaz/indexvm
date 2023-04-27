package cmd

import (
	"log"

	"fmt"

	hutils "github.com/ava-labs/hypersdk/utils"
	"github.com/spf13/cobra"
)

var burnCmd = &cobra.Command{
	Use:   "burn [options]",
	Short: "burn ",
	RunE:  burnFunc,
}

func burnFunc(_ *cobra.Command, args []string) error {
	log.Printf("burnFunc")

	amt, err := getBurnOp(args)
	if err != nil {
		return err
	}
	log.Printf("amt: %v", amt)

	// ctx := context.Background()
	// // cli := client.New(uri)
	// // c := controller.New()
	// cli := client.New(uri)
	// networkID, subnetID, chainID, err := cli.Network(context.Background())
	// if err != nil {
	// 	return err
	// }
	// log.Printf("networkID=%d subnetID=%s chainID=%s", networkID, subnetID, chainID)
	// u,l,e,err:= cli.Balance(ctx, addr)
	// u, l, e, err := cli.Balance(ctx, fmt.Sprintf("%v", addr))
	// // u, l, err := storage.GetBalanceFromState(ctx, c.inner.ReadState, addr)
	// if err != nil {
	// 	log.Printf("")
	// 	return err
	// }
	// log.Printf("u: %v, l: %v, e: %v ", u, l, e)

	return nil
}

func getBurnOp(args []string) (value uint64, err error) {
	if len(args) != 1 {
		return 0, fmt.Errorf(
			"expected exactly 1 arguments, got %d",
			len(args),
		)
	}

	value, err = hutils.ParseBalance(args[0])
	if err != nil {
		return 0, fmt.Errorf("%w: failed to parse %s", err, args[0])
	}
	return value, nil
}
