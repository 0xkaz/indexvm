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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var burn3Cmd = &cobra.Command{
	Use:   "burn3 [options]",
	Short: "burn3 ",
	RunE:  burn3Func,
}

func burn3Func(_ *cobra.Command, args []string) error {
	log.Printf("burn3Func")
	priv, err := crypto.LoadKey(privateKeyFile)
	if err != nil {
		log.Printf("err: %v", err)
		return err
	}
	factory := auth.NewED25519Factory(priv)
	// factory := auth.NewDirectFactory(priv)

	ctx := context.Background()
	// currentChainID, priv, factory, cli, err := defaultActor()
	// if err != nil {
	// 	return err
	// }
	// log.Printf("currentChainID: %v", currentChainID)

	// log.Printf("priv: %v", priv)
	// to, value, err := getBurn3Op(args)
	value, err := getBurn3Op(args)
	if err != nil {
		return err
	}

	// amt, err := getBurn3Op(args)
	// if err != nil {
	// 	return err
	// }
	log.Printf("privateKeyFile: %v", privateKeyFile)

	log.Printf("value: %v", value)

	log.Printf("uri: %v", uri)

	cli := client.New(uri)
	submit, tx, _, err := cli.GenerateTransaction(ctx, &actions.Burn3{
		// Amount: amt,
		Value: value,
		// To:    to,
		To: priv.PublicKey(),
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

// func getBurn3Op(args []string) (value uint64, err error) {
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

func getBurn3Op(args []string) (value uint64, err error) {
	if len(args) != 1 {
		return 0, fmt.Errorf(
			"expected exactly 1 arguments, got %d",
			len(args),
		)
	}

	// addr, err := utils.ParseAddress(args[0])
	// if err != nil {
	// 	return crypto.EmptyPublicKey, 0, fmt.Errorf("%w: failed to parse address %s", err, args[0])
	// }
	value, err = hutils.ParseBalance(args[0])
	if err != nil {
		return 0, fmt.Errorf("%w: failed to parse %s", err, args[0])
	}
	return value, nil
}
