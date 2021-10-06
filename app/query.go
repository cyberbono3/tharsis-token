package app

import (
	"fmt"
	//"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) Query(account string) error {

	/*
	bal, err := c.ethClient.BalanceAt(context.Background(), common.HexToAddress(account), nil)
	if err != nil {
		return err
	}
	

	fmt.Printf("Token balance of %s is: %d \n", account, bal)
	return nil
	*/



	instance, err := c.GetContractInstance(ContractAddr)
	if err != nil {
		return err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(account))
	if err != nil {
		return fmt.Errorf("instance.BalanceOf: %q", err)
	}

	fmt.Printf("Token balance of %s is: %d \n", account, bal)
	
	return nil

}
