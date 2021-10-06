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
2. CLI command `token mint <ethereum_address> <amount`
3. Input:
```
token mint 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7 1000
```
4. Output
```
1000 tokens for account address 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7 have been minted
```

### Derive Ethereum address from mnemonic phrase
1. CLI Command Usage `token ethaddress <mnemonic words>`
2. Input:
```
token ethaddress token ethaddress drastic early glass silver head satoshi hammer dawn source rubber basic balcony civil dentist oxygen spice solid script know dial tired outside conduct siege
```
2. Output:
```
Ethereum address is: 0x23F667B3f5Bc3893205B6A4edDd1D1A1629aE402 
```

### Query ERC-20 contract balance (razobratjsja s balansom)
1. Checkout [Tharsis Token repo](https://github.com/cyberbono3/tharsis-token)
2. Query token balance of a contract or any account address by running `token query <account_address>`
```
token query 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7
```
3. It outputs token balance:
```
Token balance of 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7 is: 3847775300 
```

### Transfer tokens from owner to an arbitary account

Solidity function
```
function transfer(address to, uint tokens) public returns (bool success) {
```

#### Preparation 
1. Add new `robert` key into ethermint
```
ai@ai-ThinkPad-T450:~/go/src/github.com/tharsis/ethermint$ ethermintd keys add robert --keyring-backend test

- name: robert
  type: local
  address: ethm1y0mx0vl4hsufxgzmdf8dm5w3593f4eqzmm6nfz
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"Az9bCjGcuDSwxQBROUebG//sjjQz1gBKkNL6UyDMuK3K"}'
  mnemonic: ""


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

drastic early glass silver head satoshi hammer dawn source rubber basic balcony civil dentist oxygen spice solid script know dial tired outside conduct siege

```
2. Derive an Ethereum address from `robert` key. This address will be used as `to` address.
```
token ethaddress drastic early glass silver head satoshi hammer dawn source rubber basic balcony civil dentist oxygen spice solid script know dial tired outside conduct siege

Ethereum address is: 0x23F667B3f5Bc3893205B6A4edDd1D1A1629aE402
```
3. Make sure you have some tokens at sender address (`andrei` key).
```
token query 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7

Token balance of 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7 is: 3847775300
```

TODO address all edge cases (overflow, underflow etc, transfer to 0x0000, transfer to itself)

#### Execution
1. Input
```
token transfer 0x23F667B3f5Bc3893205B6A4edDd1D1A1629aE402 1000
```
2. Output
```
bigIntTokens 100
100 tokens from 0xdded6aC7e0A13db5eE70810Bf07E4a875859e1A7 to 0x23F667B3f5Bc3893205B6A4edDd1D1A1629aE402 have been successfully transferred 
tx hash: 0x404e9bf4fef733aef1b09e7f9f9bf1452f09f34d2740f606be68d59d0a64df68 
```

### TODO address all corner cases of `transfer` function:
1. No value to pay for a gas to run `transfer tokens transaction`
2. Transaction execution fails if `to` is 0 address.
3. Transaction execution fails if `balances[msg.sender]` < `tokens`
4. Transaction execution fails if `balances[to]` + `tokens` < `balances[to]`
etc






