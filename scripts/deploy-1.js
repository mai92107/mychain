import hardhat from "hardhat";
const { ethers } = hardhat;

async function main(){
    const MyToken = await ethers.getContractFactory("MyToken");

    const initialSupply = ethers.parseUnits("1000", 18);
    
    const token = await MyToken.deploy(initialSupply);
    await token.waitForDeployment();

    console.log("MyToken deployed to:", await token.getAddress());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});