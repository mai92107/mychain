import hardhat from "hardhat";
const { ethers } = hardhat;

async function main() {
    const [user1, user2] = await ethers.getSigners();

    const tokenAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
    const autoSenderAddress = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9";

    const MyToken = await ethers.getContractFactory("MyToken");
    const MySender = await ethers.getContractFactory("AutoSender");

    const token = MyToken.attach(tokenAddress);
    const sender = MySender.attach(autoSenderAddress);                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                

    console.log("------Approve --- user1 批准 autoSender 可以花 100 token --- 取得授權後 接著存 100 到合約 ----")
    await token.approve(autoSenderAddress, ethers.parseUnits("100", 18));
    await sender.deposit(ethers.parseUnits("100",18));

    console.log(" user2 呼叫 sendReward ");
    await sender.sendReward(user2.address);

    console.log(" user2 餘額： ", ethers.formatUnits(await token.balanceOf(user2.address)));
}

main((error) => {
    console.log(error);
    process.exitCode = 1;
})