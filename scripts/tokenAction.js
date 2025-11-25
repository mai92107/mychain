import hardhat from "hardhat";
const { ethers } = hardhat;

async function main () {
    
    const [user1, user2, user3] = await ethers.getSigners();

    const tokenAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
    const MyToken = await ethers.getContractFactory("MyToken");
    const token = MyToken.attach(tokenAddress);

    console.log("------------ 初始餘額 ------------");
    console.log("user1: ", ethers.formatUnits(await token.balanceOf(user1.address)));
    console.log("user2: ", ethers.formatUnits(await token.balanceOf(user2.address)));

    console.log("\n------Transfer: user1 轉 10 給 user2 ----------");
    await token.transfer(user2.address, ethers.parseUnits("10", 18));
    console.log("user2 現在有 ", ethers.formatUnits(await token.balanceOf(user2.address)));
    console.log("user1 現在有 ", ethers.formatUnits(await token.balanceOf(user1.address)));

    console.log("\n------Approve: user2 認可 user3 可以花 5 ------");
    await token.connect(user2).approve(user3.address, ethers.parseUnits("5", 18));
    const allowance = await token.allowance(user2.address, user3.address);
    console.log("user3 獲得 user2 的 allowance: ", ethers.formatUnits(allowance));

    console.log("\n------transferfrom: user3 代替 user2 花 3 給 user1 ------");
    await token.connect(user3).transferFrom(user2.address, user1.address, ethers.parseUnits("3", 18));
    console.log("user3 現在有 ", ethers.formatUnits(await token.balanceOf(user3.address)));
    console.log("user2 現在有 ", ethers.formatUnits(await token.balanceOf(user2.address)));
    console.log("user1 現在有 ", ethers.formatUnits(await token.balanceOf(user1.address)));

}

main().catch((error) => {
    console.log(error);
    process.exitCode = 1;
})