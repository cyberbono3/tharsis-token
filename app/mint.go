package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) Mint(address string, tokens string) error {
	// set transaction opts
	auth, _, err := c.setupTransOpts()
	if err != nil {
		return err
	}

	// load contract
	instance, err := c.GetContractInstance(ContractAddr)
	if err != nil {
		return err
	}

	amount, err := bigIntFromStr(tokens, true)
	if err != nil {
		return err
	}

	tx, err := instance.Mint(auth, common.HexToAddress(address), amount)
	if err != nil {
		return err
	}

	fmt.Printf("%s tokens for account address %s has been minted", tokens, tx.Hash().Hex())

	return nil
}
