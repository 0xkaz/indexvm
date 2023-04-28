package cmd

import (
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/hypersdk/examples/tokenvm/auth"
	hutils "github.com/ava-labs/hypersdk/utils"
	"github.com/ava-labs/indexvm/client"
	"github.com/ava-labs/indexvm/utils"
)

func defaultActor() (ids.ID, crypto.PrivateKey, *auth.ED25519Factory, *client.Client, error) {
	priv, err := GetDefaultKey()
	if err != nil {
		return ids.Empty, crypto.EmptyPrivateKey, nil, nil, err
	}
	chainID, uris, err := GetDefaultChain()
	if err != nil {
		return ids.Empty, crypto.EmptyPrivateKey, nil, nil, err
	}
	log.Printf("chainID: %v, uris: %v ", chainID, uris)
	// For [defaultActor], we always send requests to the first returned URI.
	return chainID, priv, auth.NewED25519Factory(priv), client.New(uris[0]), nil
}

func GetDefaultKey() (crypto.PrivateKey, error) {
	v, err := GetDefault(defaultKeyKey)
	if err != nil {
		return crypto.EmptyPrivateKey, err
	}
	if len(v) == 0 {
		return crypto.EmptyPrivateKey, ErrNoKeys
	}
	publicKey := crypto.PublicKey(v)
	priv, err := GetKey(publicKey)
	if err != nil {
		return crypto.EmptyPrivateKey, err
	}
	hutils.Outf("{{yellow}}address:{{/}} %s\n", utils.Address(publicKey))
	return priv, nil
}

func GetDefaultChain() (ids.ID, []string, error) {
	v, err := GetDefault(defaultChainKey)
	if err != nil {
		return ids.Empty, nil, err
	}
	if len(v) == 0 {
		return ids.Empty, nil, ErrNoChains
	}
	chainID := ids.ID(v)
	uris, err := GetChain(chainID)
	if err != nil {
		return ids.Empty, nil, err
	}
	hutils.Outf("{{yellow}}chainID:{{/}} %s\n", chainID)
	return chainID, uris, nil
}

func CloseDatabase() error {
	if db == nil {
		return nil
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("unable to close database: %w", err)
	}
	return nil
}
