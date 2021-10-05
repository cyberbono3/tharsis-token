package app

import (
	"testing"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)


// TODO fix error "no contract code at given address"
// add scenario with from address does not match ownerAddr to test onlyOwner modifier
func TestMint(t *testing.T){
	amount := "1000"
	client, err := NewClient(Mnemonic)
	require.NoError(t, err)

	ownerAddr, err := addressFromPrivKey(client.privateKey)
	require.NoError(t, err)

	instance, err := client.GetContractInstance(ContractAddr)
	require.NoError(t, err)

	// balBefore 0 tokens
	balBefore, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
	require.NoError(t, err)

	// mint 1000 tokens
	err = client.Mint(ownerAddr.Hex(), amount)
	require.NoError(t, err)

	// balAfter 1000 tokens
	balAfter, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
	require.NoError(t, err)

	amountBigInt, err := bigIntFromStr(amount, false)  
	require.NoError(t, err)

	sub := big.NewInt(0).Sub(balAfter, balBefore)  
	require.Equal(t, amountBigInt, sub)
}