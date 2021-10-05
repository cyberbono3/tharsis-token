package app

import "fmt"

func DeriveEthAddressFrom(mnemonic string) error {
	privKey, err := privKeyFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	addr, err := addressFromPrivKey(privKey)
	if err != nil {
		return err
	}

	fmt.Printf("Ethereum address is: %s \n", addr.Hex())

	return nil
}