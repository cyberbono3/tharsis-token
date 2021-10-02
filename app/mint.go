package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) Mint(to string, tokens string) error {
	auth, err := c.setupTransOpts()
	if err != nil {
		return err
	}

	instance, err := c.GetContractInstance(ContractAddr)
	if err != nil {
		return err
	}

	amount, err := bigIntTokensFromStr(tokens)
	if err != nil {
		return err
	}

	tx, err := instance.Mint(auth, common.HexToAddress(to), amount)
	if err != nil {
		return err
	}

	fmt.Printf("%d tokens for account address %s has been minted", tokens, tx.Hash().Hex())

	return nil
}
