package app

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/tharsis/token/erc20"
)

const(
	gasLimit = 10000000
	gasPrice = 100
	value = 300

	rpcEndpoint = "http://0.0.0.0:8545"
	ContractAddr = "0x2A41ea95F96A87dC21f9F0dfD1f9848357D9149a"
	mnemonic = "sight cotton inmate increase build victory emerge flee rhythm begin physical copy elite drill trash immense doctor doll bundle person whale discover they witness"
)

type Client struct {
	ethClient  *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

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


func (c *Client) setupTransOpts() (*bind.TransactOpts, error) {
	publicKeyECDSA, ok := c.privateKey.Public().(*ecdsa.PublicKey)
    if !ok {
		return nil, errors.New("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := c.ethClient.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return nil, err
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

	return auth, nil
}

// returns instance, contractAddrStr, error
func (c *Client) DeployContract() error {
	if c.privateKey == nil {
		return errors.New("privateKey is nil")
	}

	if c.ethClient == nil {
		return errors.New("ethClient is nil")
	}

    auth, err := c.setupTransOpts()
	if err != nil {
		return err
	}

	// address, tx, instance, err := token.DeployToken(auth, client)
    addr, tx, _, err := erc20.DeployErc20(auth, c.ethClient)
    if err != nil {
        return fmt.Errorf("DeployToken err: %q", err)
    }

	fmt.Println("contract has been successfully deployed at: ", addr.Hex())
    fmt.Println("tx hex", tx.Hash().Hex()) 

	return nil
}


func (c *Client) GetContractInstance(contractHexStr string) (*erc20.Erc20, error)   {
	address := common.HexToAddress(contractHexStr)
	instance, err := erc20.NewErc20(address, c.ethClient)
	if err != nil {
  		return nil, err
	}

	return instance, nil
}


func DisplayTokenBalance(instance *erc20.Erc20, addr string) (*big.Int, error) {
	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		return nil, err
	}

	return bal, nil
}


func (c *Client) TransferTokens(to string, tokens *big.Int) error {
   // get contract instance 
    auth, err := c.setupTransOpts()
	if err != nil {
		return err
	}

	instance, err := c.GetContractInstance(ContractAddr)
	if err != nil {
		return err
	}

	
	tx, err := instance.Transfer(auth, common.HexToAddress(to), tokens)
	if err != nil {
		return err
	}

	fmt.Printf("transfer tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	return nil
}


func initEthClient(rpcEndpoint string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return client, nil
}



// add godoc
func privKeyFromMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil,err
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
