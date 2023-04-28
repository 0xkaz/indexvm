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
	"github.com/ava-labs/indexvm/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var burn2Cmd = &cobra.Command{
	Use:   "burn2 [options]",
	Short: "burn2 ",
	RunE:  burn2Func,
}

func burn2Func(_ *cobra.Command, args []string) error {
	log.Printf("burn2Func")
	priv, err := crypto.LoadKey(privateKeyFile)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	factory := auth.NewDirectFactory(priv)
	// log.Printf("priv: %v", priv)
	to, value, err := getBurn2Op(args)
	if err != nil {
		return err
	}

	// amt, err := getBurn2Op(args)
	// if err != nil {
	// 	return err
	// }
	log.Printf("privateKeyFile: %v", privateKeyFile)
	log.Printf("to: %v", to)
	log.Printf("value: %v", value)
	ctx := context.Background()
	log.Printf("uri: %v", uri)

	cli := client.New(uri)
	submit, tx, _, err := cli.GenerateTransaction(ctx, &actions.Burn2{
		// Amount: amt,
		Value: value,
		To:    to,
	}, factory)

	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	log.Printf("submit..")
	log.Printf("tx: %v", tx)
	//
	if err := submit(ctx); err != nil {
		log.Printf("submit err=%v", err)
		return err
	}
	// wait for transactions
	if err := cli.WaitForTransaction(ctx, tx.ID()); err != nil {
		return err
	}
	color.Green("burned %d", value)

	return nil
}

// func getBurn2Op(args []string) (value uint64, err error) {
// 	if len(args) != 1 {
// 		return 0, fmt.Errorf(
// 			"expected exactly 1 arguments, got %d",
// 			len(args),
// 		)
// 	}

// 	value, err = hutils.ParseBalance(args[0])
// 	if err != nil {
// 		return 0, fmt.Errorf("%w: failed to parse %s", err, args[0])
// 	}
// 	return value, nil
// }

func getBurn2Op(args []string) (to crypto.PublicKey, value uint64, err error) {
	if len(args) != 2 {
		return crypto.EmptyPublicKey, 0, fmt.Errorf(
			"expected exactly 2 arguments, got %d",
			len(args),
		)
	}

	addr, err := utils.ParseAddress(args[0])
	if err != nil {
		return crypto.EmptyPublicKey, 0, fmt.Errorf("%w: failed to parse address %s", err, args[0])
	}
	value, err = hutils.ParseBalance(args[1])
	if err != nil {
		return crypto.EmptyPublicKey, 0, fmt.Errorf("%w: failed to parse %s", err, args[1])
	}
	return addr, value, nil
}
