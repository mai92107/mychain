// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract AutoSender is Ownable {
    IERC20 public token;
    uint256 public rewardAmount = 1 * 10**18;

    event RewardSent(address indexed to, uint256 amount);

    constructor(address tokenAddress) Ownable(msg.sender){
        token = IERC20(tokenAddress);
    }

    function setRewardAmt(uint256 amount) external onlyOwner{
        rewardAmount = amount;
    }

    function sendReward(address to) external {
        require(token.balanceOf(address(this)) >= rewardAmount, "Not Enough Tokens...");
        token.transfer(to, rewardAmount);
        emit RewardSent(to, rewardAmount);
    }

    function deposit(uint256 amount) external {
        token.transferFrom(msg.sender, address(this), amount);
    }
}