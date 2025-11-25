import hardhat from "hardhat";
const { ethers } = hardhat;

async function main() {
  const Storage = await ethers.getContractFactory("Storage");
  const storage = await Storage.deploy();

  await storage.waitForDeployment?.();

  console.log("Storage deployed to:", await storage.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
