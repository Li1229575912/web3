// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title Simple ERC20 Token (owner-mintable)
/// @notice 作业示例：遵循 IERC20 的最小实现 + 仅所有者可增发
contract SimpleERC20 {
    // ===== 基本信息 =====
    string public name;
    string public symbol;
    uint8 public immutable decimals; // 常见为 18
    uint256 public totalSupply;

    // ===== 所有者（用于 mint 权限）=====
    address public owner;
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    // ===== 核心存储 =====
    mapping(address => uint256) private _balances;
    // allowance[owner][spender] => amount
    mapping(address => mapping(address => uint256)) private _allowances;

    // ===== 事件（与 IERC20 保持一致）=====
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    constructor(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 initialSupply   // 按 decimals 计量，如 1000 * 10**18
    ) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        owner = msg.sender;

        // 初始铸造给部署者
        _mint(msg.sender, initialSupply);
    }

    // ===== 标准只读 =====
    function balanceOf(address account) external view returns (uint256) {
        return _balances[account];
    }
    function allowance(address tokenOwner, address spender) external view returns (uint256) {
        return _allowances[tokenOwner][spender];
    }

    // ===== 转账 =====
    function transfer(address to, uint256 amount) external returns (bool) {
        _transfer(msg.sender, to, amount);
        return true;
    }

    // ===== 授权 + 代扣转账 =====
    function approve(address spender, uint256 amount) external returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(address from, address to, uint256 amount) external returns (bool) {
        uint256 currentAllowance = _allowances[from][msg.sender];
        require(currentAllowance >= amount, "ERC20: insufficient allowance");
        // 更新授权（如需“无限授权”语义，可在currentAllowance==type(uint256).max时跳过减少）
        unchecked { _approve(from, msg.sender, currentAllowance - amount); }
        _transfer(from, to, amount);
        return true;
    }

    // ===== 仅所有者可增发 =====
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    // ===== 内部工具函数 =====
    function _transfer(address from, address to, uint256 amount) internal {
        require(to != address(0), "ERC20: transfer to zero");
        uint256 fromBal = _balances[from];
        require(fromBal >= amount, "ERC20: balance too low");
        unchecked {
            _balances[from] = fromBal - amount;
        }
        _balances[to] += amount;
        emit Transfer(from, to, amount);
    }

    function _approve(address tokenOwner, address spender, uint256 amount) internal {
        require(spender != address(0) && tokenOwner != address(0), "ERC20: zero addr");
        _allowances[tokenOwner][spender] = amount;
        emit Approval(tokenOwner, spender, amount);
    }

    function _mint(address to, uint256 amount) internal {
        require(to != address(0), "ERC20: mint to zero");
        totalSupply += amount;
        _balances[to] += amount;
        emit Transfer(address(0), to, amount);
    }
}
