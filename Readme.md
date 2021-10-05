# Tharsis coding challenge (WIP)

This application  written in Go exposes CLI commands for quering and transfering tokens using ERC-20 contract deployed on Ethermint local node.

## Description:
This application contains the following packages:

1.**app** takes care about application logic

2.**cmd** contains CLI logic

3.**erc20** contains files associated with ERC20 token


## Getting Started

### Ethermint preparation

1. Checkout [Ethermint](https://github.com/tharsis/ethermint) repo and install it following Ethermint installation [guide](https://ethermint.dev/quickstart/installation.html)
2. Run ethermint Node on a background by running `./init.sh`. It will start a node and set up create a new key `mykey` and fund it with `aphoton` tokens
3. Add a new key `andrei` and save `mnemonic` words.
```
ethermintd keys add andrei --keyring-backend test
```
Output:
```
- name: andrei
  type: local
  address: ethm1mhkk43lq5y7mtmnssy9lqlj2sav9ncd8tpma3a
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A7Hpzr+EEz8K1Qtg/Wo5Pb1Je5uDZvBGBL9RLng1owTO"}'
  mnemonic: ""


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

language fit sniff present wonder fish absent direct sheriff innocent thought educate bitter current design mother sunset name pelican rate clip eternal medal popular

```

4. View `mykey` address 
```
ethermintd keys list
```
Output:
```
- name: andrei  type: local
  address: ethm1mhkk43lq5y7mtmnssy9lqlj2sav9ncd8tpma3a
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A7Hpzr+EEz8K1Qtg/Wo5Pb1Je5uDZvBGBL9RLng1owTO"}'
  mnemonic: ""
- name: mykey
  type: local
  address: ethm157h9798qgrxfxzv6t40dx98lkqty8eqcay9mfs
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AmemhAX8WVP8qjffv0zmc3frUYAEGGKldYTRMUv1OfEI"}'
  mnemonic: ""
```

5. We have to fund `andrei` key with some tokens to pay for future gas costs. Transfer some tokens from `mykey` to `andrei` and confirm the transaction.
```
ethermintd tx bank send ethm157h9798qgrxfxzv6t40dx98lkqty8eqcay9mfs ethm1mhkk43lq5y7mtmnssy9lqlj2sav9ncd8tpma3a  1000000000aphoton --fees 20aphoton
```
where `ethm157h9798qgrxfxzv6t40dx98lkqty8eqcay9mfs` is `mykey` address, `ethm1mhkk43lq5y7mtmnssy9lqlj2sav9ncd8tpma3a` is `andrei` address. 
Output:
```
{"body":{"messages":[{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"ethm157h9798qgrxfxzv6t40dx98lkqty8eqcay9mfs","to_address":"ethm1mhkk43lq5y7mtmnssy9lqlj2sav9ncd8tpma3a","amount":[{"denom":"aphoton","amount":"1000000000"}]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[{"denom":"aphoton","amount":"20"}],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: Y
code: 0
codespace: ""
data: ""
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 3F2F06FA30D7CF7E5F891A8D41198389C426EC80ED8999679F4308DE075CE7F3
```

### Program Installion

1. Checkout my [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token) and install it.
```
make install
```
2. Complile ERC-20 contract using solc and abigen following Ethereum book [guide](https://goethereumbook.org/smart-contract-compile/)

```
git clone github.com/ethereum/go-ethereum.git
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools

solc --abi erc20/token.sol
solc --bin erc20/token.sol
abigen --bin=erc20/token.bin --abi=erc20/token.bin --pkg=erc20 --out=erc20/Token.go --alias _totalSupply=TotalSupply1
```

## Program Execution

### Deploy ERC-20 contract

1. Checkout [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token)
2. Hardcode mnemonic phrase of `andrei` key in `app/client.go` (the better option to use some config in encrypted fashion)
```
Mnemonic = "language fit sniff present wonder fish absent direct sheriff innocent thought educate bitter current design mother sunset name pelican rate clip eternal medal popular"
```
3. Deploy Erc-20 [token](https://github.com/cyberbono3/tharsis-token/blob/master/erc20/token.sol) on Ethermint using `token deploy`  command. It will use `andrei` account under the hood to pay for gas costs. You can read full desciption about `deploy` command [here](https://github.com/cyberbono3/tharsis-token/blob/master/cmd/deploy.go)
```
token deploy
```
4. See deployment confirmation that look like that:
```
contract from 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7 has been successfully deployed at: 0x699115980439687bEfC301549599edF5e6A28716
deploy called
```
5. Hardcode contract address in `app/client.go` (future use of config for that)
```
ContractAddr = "0x699115980439687bEfC301549599edF5e6A28716"
```

### Mint ERC-20 Tokens
1. Checkout [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token)



### Query ERC-20 contract balance 
1. Checkout [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token)
2. Query token balance of a contract or any account address by running `token query <account_address>`
```
token query 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7
```
3. It should output a total supply of contract,
4. Unfortunately, it raises an error: 
```
Error: no contract code at given address
```
5. The error arises here:
```
Error: instance.BalanceOf: "no contract code at given address"
```
6. Namely, `instance.BalanceOf method yields an error:
```
bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
if err != nil {
	return fmt.Errorf("instance.BalanceOf: %q", err)
}
```


### Transfer tokens from owner to an arbitary account (WORK IN PROGRESS)
Solidity function
```
function transfer(address to, uint tokens) public returns (bool success) {
```
1. Check owner's token balance using `balanceOf`. 
2. Mint tokens on owner acccount\if owner's token balance is zero.
3. Use `ethermitd keys add to` to add destination account and derive it's ethereum address from mnemonic phrase.
4. Follow the [guide](https://goethereumbook.org/transfer-tokens/) to transfer `tokens` from `owner` to `to` address.
```
token trasfer <to_address> <tokens>

```
### TODO
* Resolve an error 
* Address and test corner cases in transfer scenario :
	1. No value to pay for a gas to run `transfer tokens transaction`
	2. Transaction execution fails if `to` is 0 address.
	3. Transaction execution fails if `balances[msg.sender]` < `tokens`
	4. Transaction execution fails if `balances[to]` + `tokens` < `balances[to]`

* Add cobra CLI tests `deploy_test.go`,`query_test.go`,`transfer_test.go`,`name_test.go`

* Add more testing 

* Fix `token name` error





