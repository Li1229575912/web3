// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract BinarySearch {
    // 二分查找
    // arr 必须是升序排列
    // 如果找到 target，返回索引；否则返回 -1
    function binarySearch(uint[] memory arr, uint target) public pure returns (int) {
        int left = 0;
        int right = int(arr.length) - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (arr[uint(mid)] == target) {
                return mid; // 找到目标
            } 
            else if (arr[uint(mid)] < target) {
                left = mid + 1; // 去右半边找
            } 
            else {
                right = mid - 1; // 去左半边找
            }
        }
        return -1; // 没找到
    }
}
