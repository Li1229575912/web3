//用 solidity 实现整数转罗马数字

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntToRoma {

    // 将整数转换为罗马数字
    function intToRoma(uint256 num) public pure returns (string memory) {
        // 定义罗马数字和对应的数值
        string[13] memory romanSymbols = [
            "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"
        ];
        uint256[13] memory values;
        
        string memory result = "";

        values[0] = 1000;
        values[1] = 900;
        values[2] = 500;
        values[3] = 400;
        values[4] = 100;
        values[5] = 90;
        values[6] = 50;
        values[7] = 40;
        values[8] = 10;
        values[9] = 9;
        values[10] = 5;
        values[11] = 4;
        values[12] = 1;

        // 从最大值到最小值依次处理
        for (uint256 i = 0; i < 13; i++) {
            while (num >= values[i]) {
                num -= values[i];
                result = string(abi.encodePacked(result, romanSymbols[i]));
            }
        }
        
        
        return result;
    }
}
