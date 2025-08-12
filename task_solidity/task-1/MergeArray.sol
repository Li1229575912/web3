// 将两个有序数组合并为一个有序数组。

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract MergeSortedArrays {
    // 合并两个有序的 uint 数组
    function merge(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        uint[] memory merged = new uint[](len1 + len2);

        uint i = 0; // arr1 索引
        uint j = 0; // arr2 索引
        uint k = 0; // merged 索引

        // 逐一比较，小的先放入 merged
        while (i < len1 && j < len2) {
            if (arr1[i] <= arr2[j]) {
                merged[k] = arr1[i];
                i++;
            } else {
                merged[k] = arr2[j];
                j++;
            }
            k++;
        }

        // 处理 arr1 剩余元素
        while (i < len1) {
            merged[k] = arr1[i];
            i++;
            k++;
        }

        // 处理 arr2 剩余元素
        while (j < len2) {
            merged[k] = arr2[j];
            j++;
            k++;
        }

        return merged;
    }
}
