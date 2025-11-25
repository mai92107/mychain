// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Storage {
    uint256 private number;

    // 設定值
    function set(uint256 _number) public {
        number = _number;
    }

    // 取得值
    function get() public view returns (uint256) {
        return number;
    }
}

