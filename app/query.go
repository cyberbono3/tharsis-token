package app

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) Query(address string) error {
	instance, err := c.GetContractInstance(ContractAddr)
	if err != nil {
		return err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return fmt.Errorf("instance.BalanceOf: %q", err)
	}

	fmt.Printf("Token balance of %s is: %d \n", address, bal)

	return nil

}
