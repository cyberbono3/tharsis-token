package app

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func bigIntFromStr(str string, toWei bool) (*big.Int, error) {
	// convert string to *big.Int
	n := new(big.Int)
	var ok bool
	if toWei {
		n, ok = n.SetString(str+"000000000000000000", 10) //str token in wei
	} else {
		n, ok = n.SetString(str, 10)
	}

	if !ok {
		return nil, errors.New("unable convert string to *big.Int")
	}

	return n, nil
}

func addressFromPrivKey(privKey *ecdsa.PrivateKey) (common.Address, error) {

	publicKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress, nil
}

// initEthClient sets up ethClient based on rpcEndpoint provided
func initEthClient(rpcEndpoint string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// privKeyFromMnemonic derives a private key from mnemonic phrase
func privKeyFromMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false) // account.Address.Hex(),
	if err != nil {
		return nil, err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
