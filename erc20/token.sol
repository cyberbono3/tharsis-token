pragma solidity >=0.5.11;

library SafeMath {
    function add(uint a, uint b) public pure returns (uint c) {
        c = a + b;
        require(c >= a);
    }
    function sub(uint a, uint b) public pure returns (uint c) {
        require(b <= a);
        c = a - b;
    }
    function mul(uint a, uint b) public pure returns (uint c) {
        c = a * b;
        require(a == 0 || c / a == b);
    }
    function div(uint a, uint b) public pure returns (uint c) {
        require(b > 0);
        c = a / b;
    }
}

// todo add address validation
contract Token {

    using SafeMath for uint256;

    string public symbol;
    string public  name;
    uint8 public decimals;
    uint public _totalSupply;
    address public owner;

    mapping(address => uint) balances;
    mapping(address => mapping(address => uint)) allowed;

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);

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
        balances[msg.sender] = SafeMath.sub(balances[msg.sender], tokens);
        balances[to] = SafeMath.add(balances[to], tokens);
        emit Transfer(msg.sender, to, tokens);
        return true;
    }


    // ------------------------------------------------------------------------
    // Token owner can approve for spender to transferFrom(...) tokens
    // from the token owner's account
    //
    // https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md
    // recommends that there are no checks for the approval double-spend attack
    // as this should be implemented in user interfaces 
    // ------------------------------------------------------------------------
    function approve(address spender, uint tokens) public returns (bool success) {
        allowed[msg.sender][spender] = tokens;
        emit Approval(msg.sender, spender, tokens);
        return true;
    }


    // ------------------------------------------------------------------------
    // Transfer tokens from the from account to the to account
    // 
    // The calling account must already have sufficient tokens approve(...)-d
    // for spending from the from account and
    // - From account must have sufficient balance to transfer
    // - Spender must have sufficient allowance to transfer
    // - 0 value transfers are allowed
    // ------------------------------------------------------------------------
    function transferFrom(address from, address to, uint tokens) public returns (bool success) {
        balances[from] = SafeMath.sub(balances[from], tokens);
        allowed[from][msg.sender] = SafeMath.sub(allowed[from][msg.sender], tokens);
        balances[to] = SafeMath.add(balances[to], tokens);
        emit Transfer(from, to, tokens);
        return true;
    }


    // ------------------------------------------------------------------------
    // Returns the amount of tokens approved by the owner that can be
    // transferred to the spender's account
    // ------------------------------------------------------------------------
    function allowance(address tokenOwner, address spender) public view returns (uint remaining) {
        return allowed[tokenOwner][spender];
    }

    // add safemath methods
    // add address validation
    function mint(address account, uint256 amount) internal {
        _totalSupply += amount;
        balances[account] += amount;
        emit Transfer(address(0), account, amount);
    }
}
