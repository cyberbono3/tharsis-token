package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
)

func (c *Client) TransferTokens(to string, tokens string) error {
	// get contract instance
	fromAddress, err := addressFromPrivKey(c.privateKey)
	if err != nil {
		return err
	}

	nonce, err := c.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(ContractAddr)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	//fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//fmt.Println(hexutil.Encode(paddedAddress))

	bigIntTokens, err := bigIntFromStr(tokens)
	if err != nil {
		return err
	}
	fmt.Println("bigIntTokens", bigIntTokens)

	paddedAmount := common.LeftPadBytes(bigIntTokens.Bytes(), 32)
	//fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, tokenAddress, bigIntTokens, gasLimit, big.NewInt(gasPrice), data)

	chainID, err := c.ethClient.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("NetworkID err: %q", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), c.privateKey)
	if err != nil {
		return fmt.Errorf("SignTx err: %q", err)
	}

	err = c.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("SendTransaction err: %q", err)
	}

	fmt.Printf("%d tokens from %s to %s have been successfully transferred \n", bigIntTokens, fromAddress.Hex(), to)
	fmt.Printf("tx hash: %s \n", signedTx.Hash().Hex())

	return nil

}
