// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const { ethers } = require("hardhat");
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy
  const signers = await hre.ethers.getSigners();
  // const addr = prompt("Enter factory address: ");

  const factory = await ethers.getContractAt(
    "StorageNodeFactory",
    "0x5FbDB2315678afecb367f032d93F642f64180aa3",
    signers[0]
  );

  const tx = await factory.createStorageContract(
    ethers.utils.arrayify("0x123456"), // TLSCert
    "192.168.0.1:8000" // HOST
  );

  const b = await tx.wait();

  const eventInd = b.events.findIndex((f) => f.event == "EvNewStorageContract");
  const event = b.events[eventInd];
  console.log("New storage contract created at address", event.args.addr);
}
// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
