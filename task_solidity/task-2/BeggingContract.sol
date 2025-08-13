// 讨饭合约 地址：0x6D593D300004279B048F1b43d89381A94c9491cD

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract BeggingContract {
    // ===== 所有者 =====
    address public owner;

    // ===== 记录每位捐赠者的累计捐赠金额（wei）=====
    mapping(address => uint256) private donations;

    // ===== 事件 =====
    event Donated(address indexed donor, uint256 amount);
    event Withdrawn(address indexed owner, uint256 amount);

    // ===== 构造：设置所有者 =====
    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    // ===== 捐赠函数：带 data 调用时使用 =====
    function donate() public payable {
        require(msg.value > 0, "No ETH sent");
        donations[msg.sender] += msg.value;
        emit Donated(msg.sender, msg.value);
    }

    // ===== 直接转账（无 data）也记账 =====
    receive() external payable {
        require(msg.value > 0, "No ETH sent");
        donations[msg.sender] += msg.value;
        emit Donated(msg.sender, msg.value);
    }

    // ===== 仅所有者提取全部余额（使用 transfer 按题意实现）=====
    function withdraw() external onlyOwner {
        uint256 amount = address(this).balance;
        require(amount > 0, "No balance");

        (bool ok, ) = payable(owner).call{value: amount}("");
        require(ok, "withdraw failed");
        emit Withdrawn(owner, amount);
    }

    // ===== 查询某地址的累计捐赠 =====
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    // 更换所有者，方便运营
    function transferOwnership(address newOwner) external onlyOwner {
        require(newOwner != address(0), "zero address");
        owner = newOwner;
    }
}
