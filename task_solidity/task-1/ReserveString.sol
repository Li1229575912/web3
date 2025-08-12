//反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

contract ReserveString {
    function reservestring(string calldata input) external pure returns (string memory) {
        bytes memory in_str = bytes(input);
        bytes memory out_str = new bytes(in_str.length);
        for (uint i = 0; i < in_str.length; i++) 
        {
            out_str[i] = in_str[in_str.length - 1 - i];
        }
        return string(out_str);
    }
}