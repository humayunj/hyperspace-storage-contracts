const { expect } = require("chai");
const { BigNumber } = require("ethers");
const { ethers, network } = require("hardhat");

describe("Storage Node", function () {
  it("deploys contract", async function () {
    const StorageNodeContract = await ethers.getContractFactory("StorageNode");
    const node = await StorageNodeContract.deploy(
      ethers.utils.arrayify("0x123456"), // TLSCert
      "http://localhost:8000" // HOST
    );

    await node.deployed();

    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: ethers.utils.parseEther("1.0"),
    });

    console.log("BL:", await signers[0].getBalance());
    const TStart = Date.now();
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      ethers.utils.arrayify(Array.from(new Array(20)).map((_, i) => i)), // hash
      100, // file size
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      3, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    // await network.provider.send("evm_increaseTime", [3600]);
    // await network.provider.send("evm_mine");
    await node.concludeTransaction(
      1, // user node
      signers[0].address,
      ethers.utils.arrayify(Array.from(new Array(20)).map((_, i) => i)), // hash
      100, // file size
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000), // bid Amount in wei
      {
        value: 1000,
      }
    );

    await network.provider.send("evm_increaseTime", [5]);
    await network.provider.send("evm_mine");
    console.log("Fin:");
    console.log(await node.finishTransaction(signers[0].address));

    // expect(await greeter.greet()).to.equal("Hello, world!");

    // const setGreetingTx = await greeter.setGreeting("Hola, mundo!");

    // // wait until the transaction is mined
    // await setGreetingTx.wait();

    // expect(await greeter.greet()).to.equal("Hola, mundo!");
  });
});
