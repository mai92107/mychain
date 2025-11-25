import hardhat from "hardhat";
const { ethers } = hardhat;

async function main() {
    const tokenAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
    const AutoSender = await ethers.getContractFactory("AutoSender");
    const autoSender = await AutoSender.deploy(tokenAddress);

    await autoSender.waitForDeployment();

    console.log("AutoSender 部署到： ", await autoSender.getAddress());
}

main((error) => {
    console.log(error);
    process.exitCode = 1;
})