// 实现罗马数字转整数

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomaToInt {
// 把罗马字符映射成对应的整数
    function charToValue(bytes1 romanChar) internal pure returns (uint256) {
        if (romanChar == "I") return 1;
        if (romanChar == "V") return 5;
        if (romanChar == "X") return 10;
        if (romanChar == "L") return 50;
        if (romanChar == "C") return 100;
        if (romanChar == "D") return 500;
        if (romanChar == "M") return 1000;
        revert("Invalid Roman numeral");
    }

    // 主函数：罗马数字转整数
    function romanToInt(string memory roman) public pure returns (uint256) {
        bytes memory s = bytes(roman);
        uint256 total = 0;
        uint256 prevValue = 0;

        // 从右往左遍历
        for (uint256 i = s.length; i > 0; i--) {
            uint256 value = charToValue(s[i - 1]);
            if (value < prevValue) {
                total -= value; // 小数在大数左边，减
            } else {
                total += value; // 否则加
            }
            prevValue = value;
        }
        return total;
    }
}