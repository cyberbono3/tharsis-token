package app

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tharsis/token/erc20"
)

const (
	gasLimit = 10000000
	gasPrice = 3
	value    = 300

	rpcEndpoint  = "http://0.0.0.0:8545"
	ContractAddr = "0x332534B6704432bD43B61cdab476a5fe8F942963"
	Mnemonic     = "language fit sniff present wonder fish absent direct sheriff innocent thought educate bitter current design mother sunset name pelican rate clip eternal medal popular"
)

type Client struct {
	ethClient  *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

// TODO consider tocdelete mnemonic argument
// NewClient sets up a Client that contains ethClient and privKey
func NewClient(mnemonic string) (*Client, error) {
	ethClient, err := initEthClient(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	privKey, err := privKeyFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	return &Client{ethClient, privKey}, nil
}

// setupTransOpts yields an auth object that holds transaction options
func (c *Client) setupTransOpts() (*bind.TransactOpts, string, error) {
	fromAddress, err := addressFromPrivKey(c.privateKey)
	if err != nil {
		return nil, "", err
	}

	nonce, err := c.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, "", err
	}

	// TODO think to uncomment it
	/*
	   gasPrice, err := client.SuggestGasPrice(context.Background())
	   if err != nil {
	       log.Fatal(err)
	   }
	*/

	auth := bind.NewKeyedTransactor(c.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(300)
	auth.GasLimit = uint64(gasLimit)
	auth.GasPrice = big.NewInt(gasPrice)

	return auth, fromAddress.Hex(), nil
}

// GetContractInstance yields an contract instance from contract hex string.
func (c *Client) GetContractInstance(contractHexStr string) (*erc20.Erc20, error) {
	address := common.HexToAddress(contractHexStr)
	instance, err := erc20.NewErc20(address, c.ethClient)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

