package cmd

import (
	"context"
	"log"

	"fmt"

	"github.com/ava-labs/hypersdk/crypto"
	hutils "github.com/ava-labs/hypersdk/utils"
	"github.com/ava-labs/indexvm/actions"
	"github.com/ava-labs/indexvm/auth"
	"github.com/ava-labs/indexvm/client"
	"github.com/spf13/cobra"
)

var burnCmd = &cobra.Command{
	Use:   "burn [options]",
	Short: "burn ",
	RunE:  burnFunc,
}

func burnFunc(_ *cobra.Command, args []string) error {
	log.Printf("burnFunc")
	priv, err := crypto.LoadKey(privateKeyFile)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	factory := auth.NewDirectFactory(priv)
	// log.Printf("priv: %v", priv)

	amt, err := getBurnOp(args)
	if err != nil {
		return err
	}
	log.Printf("privateKeyFile: %v", privateKeyFile)
	log.Printf("amt: %v", amt)
	ctx := context.Background()
	log.Printf("uri: %v", uri)

	cli := client.New(uri)
	submit, tx, _, err := cli.GenerateTransaction(ctx, &actions.Spend{
		Amount: amt,
	}, factory)

	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	log.Printf("submit: %v", submit)
	log.Printf("tx: %v", tx)
	//
	// if err := submit(ctx); err != nil {
	// 	return err
	// }
	// // wait for transactions
	// if err := cli.WaitForTransaction(ctx, tx.ID()); err != nil {
	// 	return err
	// }
	// color.Green("burned %d", amt)

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
