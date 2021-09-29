pragma solidity >=0.5.11;

// @dev It is a good practice to use SafeMath.sol library here, but due to an error `Fatal: Contract has additional library references, please use other mode(e.g. --combined-json) to catch library infos` 
// with `abigen --bin=erc20/token.bin --abi=erc20/token.abi --pkg=erc20 --out=erc20/Token.go --alias _totalSupply=TotalSupply1`
// dealing with multiple contracts is problematic. Subsequently, I decided to remove it and add manual underflow, overflow protection
contract Token {

    string public symbol;
    string public  name;
    uint8 public decimals;
    uint public _totalSupply;
    address public owner;

    mapping(address => uint) balances;

    event Transfer(address indexed from, address indexed to, uint tokens);
    // TODO set onlyOwner
    // ------------------------------------------------------------------------
    // Constructor
    // ------------------------------------------------------------------------
    constructor() public {
        symbol = "GLD";
        name = "GOLD";
        decimals = 18;
        _totalSupply = 100000000000000000000000000;
        owner = msg.sender;
      //  balances[0x5A86f0cafD4ef3ba4f0344C138afcC84bd1ED222] = _totalSupply;
     //   emit Transfer(address(0), 0x5A86f0cafD4ef3ba4f0344C138afcC84bd1ED222, _totalSupply);
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }


    // ------------------------------------------------------------------------
    // Total supply
    // ------------------------------------------------------------------------
    function totalSupply() public view returns (uint) {
        return _totalSupply;
    }

    // ------------------------------------------------------------------------
    // Get the token balance for account tokenOwner
    // ------------------------------------------------------------------------
    function balanceOf(address tokenOwner) public view returns (uint) {
        return balances[tokenOwner];
    }


    // ------------------------------------------------------------------------
    // Transfer the balance from token owner's account to to account
    // - Owner's account must have sufficient balance to transfer
    // - 0 value transfers are allowed
    // ------------------------------------------------------------------------
    function transfer(address to, uint tokens) public returns (bool success) {
        require(to != address(0));
        require(balances[msg.sender] >= tokens); // undeflow protection
        require(balances[to] + tokens >= balances[to]);  // overflow protection

        balances[msg.sender] -= tokens;
        balances[to] += tokens;
        
        emit Transfer(msg.sender, to, tokens);
        return true;
    }

    
    // add address validation
    // check for overflow using Safemath
    function mint(address account, uint256 amount) internal onlyOwner {
        require(account != address(0));
        require(_totalSupply + amount >= _totalSupply);  // overflow protection
        require(balances[account] + amount >= balances[account]);  // overflow protection

        _totalSupply += amount;
        balances[account] += amount;
        emit Transfer(address(0), account, amount);
    }
}
