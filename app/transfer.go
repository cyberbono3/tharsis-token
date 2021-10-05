package app

import (
	"context"
//	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/crypto"
)
/*
ai@ai-ThinkPad-T450:~/go/src/github.com/tharsis/ethermint$ ethermintd keys add robert --keyring-backend test

- name: robert
  type: local
  address: ethm1y0mx0vl4hsufxgzmdf8dm5w3593f4eqzmm6nfz
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"Az9bCjGcuDSwxQBROUebG//sjjQz1gBKkNL6UyDMuK3K"}'
  mnemonic: ""


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

drastic early glass silver head satoshi hammer dawn source rubber basic balcony civil dentist oxygen spice solid script know dial tired outside conduct siege
*/
// TODO fix it - this is incorrect implemntation, see https://goethereumbook.org/transfer-tokens/ for more details
//1.  add new key to ethermint
//2.  convert it's mnemonic to the Ethereum address - this is "to"
//3.  mint tokens  on owner ethereum account
// 4. transfer

func (c *Client) TransferTokens(to string, tokens string) error {
	// get contract instance
	fromAddress, err := addressFromPrivKey(c.privateKey)
	if err != nil {
		return err
	}

	nonce, err := c.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := c.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(ContractAddr)

	//transferFnSignature := []byte("transfer(address,uint256)")
	//hash := crypto.Keccak256Hash(transferFnSignature) //sha3.NewKeccak256() is deprecated
	// TODO resolve issue hash.Sum undefined, tried Keccak256 method above
	//methodID := hash.Sum(nil)[:4]
	//fmt.Println(hexutil.Encode(methodID))             // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	bigIntTokens, err := bigIntFromStr(tokens, true)
	if err != nil {
		return err
	}

	paddedAmount := common.LeftPadBytes(bigIntTokens.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	//data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := c.ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := c.ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	return nil

}
