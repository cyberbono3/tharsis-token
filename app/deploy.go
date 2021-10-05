package app

import (
	"errors"
	"fmt"

	"github.com/tharsis/token/erc20"
)

// DeployContract deploys ERC-20 contract using privateKey and ethClient that are stored on Client.
func (c *Client) DeployContract() error {
	if c.privateKey == nil {
		return errors.New("privateKey is nil")
	}

	if c.ethClient == nil {
		return errors.New("ethClient is nil")
	}

	auth,fromAddressStr, err := c.setupTransOpts()
	if err != nil {
		return err
	}

	// address, tx, instance, err := token.DeployToken(auth, client)
	addr, _, _, err := erc20.DeployErc20(auth, c.ethClient)
	if err != nil {
		return fmt.Errorf("DeployToken err: %q", err)
	}

	fmt.Printf("contract from %s has been successfully deployed at: %s", fromAddressStr, addr.Hex())

	return nil
}
