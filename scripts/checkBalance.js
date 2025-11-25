import hardhat from "hardhat";
const { ethers } = hardhat;

async function main() {
    const [user1, user2] = await ethers.getSigners();

    const tokenAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
    const MyToken = await ethers.getContractFactory("MyToken");
    const token = MyToken.attach(tokenAddress);
    
    const balance1 = await token.balanceOf(await user1.getAddress());
    const balance2 = await token.balanceOf(await user2.getAddress());

    console.log(`User1 Balance: ${ethers.formatUnits(balance1, 18)} MTK`);
    console.log(`User2 Balance: ${ethers.formatUnits(balance2, 18)} MTK`);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});