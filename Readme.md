# Tharsis coding challenge

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
ethermintd keys add andrei 
```
`mnemonic` words are: `sight cotton inmate increase build victory emerge flee rhythm begin physical copy elite drill trash immense doctor doll bundle person whale discover they witness`


4. We have to fund `andrei` key with some tokens to pay for future gas costs. Transfer some tokens from `mykey` to `andrei` and confirm the transaction.
```
ethermintd tx bank send ethm18a9c2rz2faq5f6s5zlmancvalcpvc7qrgar48y ethm1egn3vmrezgenc406ca6dw96fgr6t27swae8fc6  1000000000aphoton --fees 20aphoton
```
where `ethm18a9c2rz2faq5f6s5zlmancvalcpvc7qrgar48y` is `mykey` address, `ethm1egn3vmrezgenc406ca6dw96fgr6t27swae8fc6` is `andrei` address. 

### Program Installion

1. Checkout my [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token) and install it.
```
make install
```
2. Complile ERC-20 contract using solc and abigen following Ethereum book [guide](https://goethereumbook.org/smart-contract-compile/)

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools

solc --abi erc20/token.sol
solc --bin erc20/token.sol
abigen --bin=erc20/token.bin --abi=erc20/token.bin --pkg=erc20 --out=erc20/Token.go --alias _totalSupply=TotalSupply1
```

## Program Execution

### Deploy ERC-20 contract

1. Checkout [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token 
2. Deploy Erc-20 [token](https://github.com/cyberbono3/tharsis-token/blob/master/erc20/token.sol) on Ethermint using `token deploy <mnemonic words>` command and saved `mnemonic` words from Ethermint Preparation Step 3. You can read full desciption about `deploy` command [here](https://github.com/cyberbono3/tharsis-token/blob/master/cmd/deploy.go)
```
token deploy sight cotton inmate increase build victory emerge flee rhythm begin physical copy elite drill trash immense doctor doll bundle person whale discover they witness
```
3. See deployment confirmation that look like that:
```
contract has been successfully deployed at:  0xd3C2901CE8AfF95176C7812DA97235238D419D0F
tx hex 0x91004fe4f5b26096bcbe3e385c39ce1778532bae50e7a554387278956e8a3d53
deploy called
```
4. Hardcode contract address and mnemonic in `app/client.go` 
```
ContractAddr = "0x9491f4c3f45956903FCD1Abbf404097a82995072"
Mnemonic = "sight cotton inmate increase build victory emerge flee rhythm begin physical copy elite drill trash immense doctor doll bundle person whale discover they witness"
```
This is some UX issue for a client. You can fix it by implementing CLI commands or using config file. I keep it out of scope of this task due to lack of time.

### Query ERC-20 contract balance 
1. Checkout [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token)
2. Query token balance of a contract or any account address by running `token query <contract_address> <account_address>`
```
token query 0xd3C2901CE8AfF95176C7812DA97235238D419D0F
```
3. It should output a total supply of contract,
```
token query 00xd3C2901CE8AfF95176C7812DA97235238D419D0F
```
4. Unfortunately, it yields an error 
```
Error: TotalSupply1 err: "no contract code at given address"
```
5. The error arises here:
```
totalSupply, err := instance.TotalSupply1(&bind.CallOpts{})
if err != nil {
	return fmt.Errorf("TotalSupply1 err: %q", err)
}
```
6. Namely, `contract.Call` method yields an error in [token.go](https://github.com/cyberbono3/tharsis-token/erc20/Token.go). need more time to debug it.
```
func (_Erc20 *Erc20Caller) TotalSupply1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "_totalSupply")
}
```
TODO scenario add an account in `ethermint` and run `token query <contract_address> <account_address>`

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

#### Corner cases in transfer scenario that has to be addressed and tested:
1. No value to pay for a gas to run `transfer tokens transaction`
2. Transaction execution fails if `to` is 0 address.
3. Transaction execution fails if `balances[msg.sender]` < `tokens`
4. Transaction execution fails if `balances[to]` + `tokens` < `balances[to]`









