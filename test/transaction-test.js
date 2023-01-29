const { expect } = require("chai");
const { BigNumber } = require("ethers");
const { ethers, network } = require("hardhat");
const { randomBytes } = require("crypto");

describe("Transactions", function () {
  let node = null;
  let snapshotId = "";
  before(async () => {});
  beforeEach(async () => {
    snapshotId = await network.provider.send("evm_snapshot");
    // console.log("snapshot:", snapshotId);
    const owner = (await ethers.getSigners())[0];

    StorageNodeContract = await ethers.getContractFactory("StorageNode");
    node = await StorageNodeContract.deploy(
      ethers.utils.arrayify("0x123456"), // TLSCert
      "http://localhost:8000", // HOST
      owner.address
    );

    await node.deployed();
  });
  afterEach(async () => {
    await network.provider.send("evm_revert", [snapshotId]);
  });
  it("validates transaction conclusion", async function () {
    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: BigInt(1000 * 2),
    });

    //ethers.utils.parseEther("1.0")
    // console.log("BL:", await signers[0].getBalance());
    const FILE_HASH = ethers.utils.arrayify(randomBytes(32));
    const TStart = Date.now();
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    await node.concludeTransaction(
      1, // user node
      signers[0].address,
      FILE_HASH,
      100, // file size
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segment s count
      BigInt(1000), // bid Amount in wei
      {
        value: 1000,
      }
    );

    await network.provider.send("evm_increaseTime", [100]);
    await network.provider.send("evm_mine");
    // console.log("Fin:");
    await node.finishTransaction(signers[0].address, FILE_HASH);
  });
  it("conclusion of non-existent entry should throw", async function () {
    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: BigInt(1000 * 2),
    });

    //ethers.utils.parseEther("1.0")
    // console.log("BL:", await signers[0].getBalance());
    const FILE_HASH = ethers.utils.arrayify(randomBytes(32));

    const TStart = Date.now();

    await expect(
      node.concludeTransaction(
        1, // user node
        signers[0].address,
        FILE_HASH,
        100, // file size
        BigInt(Math.round(TStart / 1000)), // timerStart
        BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
        10, // prove timeout length in seconds
        10, // conclude timeout
        BigInt(100), // segment s count
        BigInt(1000), // bid Amount in wei
        {
          value: 1000,
        }
      )
    ).to.be.revertedWith("no transaction entry found to conclude");

    // console.log("Fin:");
  });
  it("validates duplication is throwing error", async function () {
    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: BigInt(1000 * 2),
    });

    //ethers.utils.parseEther("1.0")
    // console.log("BL:", await signers[0].getBalance());
    const FILE_HASH = ethers.utils.arrayify(randomBytes(32));
    const TStart = Date.now();
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    await network.provider.send("evm_increaseTime", [5]); // still in conclusion period
    await network.provider.send("evm_mine");

    const TStart2 = Date.now() + 1000; // Different T start
    await expect(
      node.concludeTransaction(
        0, // storage node
        signers[0].address,
        FILE_HASH, // merkle hash
        100, // file size in bytes
        BigInt(Math.round(TStart2 / 1000)), // timerStart
        BigInt(Math.round((TStart2 + 10000) / 1000)), // timerEnd
        10, // prove timeout length in seconds
        10, // conclude timeout
        BigInt(100), // segments count
        BigInt(1000) // bid Amount in wei
      )
    ).to.be.revertedWith("file already stored or waiting for conclusion");

    // console.log("Fin:");
  });
  it("validates node can try again after conclusion expires", async function () {
    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: BigInt(1000 * 2),
    });

    //ethers.utils.parseEther("1.0")
    // console.log("BL:", await signers[0].getBalance());
    const FILE_HASH = ethers.utils.arrayify(randomBytes(32));
    const TStart = Date.now();
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    await network.provider.send("evm_increaseTime", [11]); // (expires in 10 second, we try after 11)
    await network.provider.send("evm_mine");

    const TStart2 = Date.now() + 15000; // Different T start, after 16 seconds
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart2 / 1000)), // timerStart
      BigInt(Math.round((TStart2 + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    await node.concludeTransaction(
      1, // user node
      signers[0].address,
      FILE_HASH,
      100, // file size
      BigInt(Math.round(TStart2 / 1000)), // timerStart
      BigInt(Math.round((TStart2 + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segment s count
      BigInt(1000), // bid Amount in wei
      {
        value: 1000,
      }
    );

    await network.provider.send("evm_increaseTime", [100]); // (finish)
    await network.provider.send("evm_mine");

    await node.finishTransaction(signers[0].address, FILE_HASH);

    // console.log("Fin:");
  });

  it("should revert with `concludeTimeoutLength mismatch'`", async function () {
    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: ethers.utils.parseEther("1.0"),
    });

    const FILE_HASH = ethers.utils.arrayify(randomBytes(32));
    const TStart = Date.now();
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    // await network.provider.send("evm_increaseTime", [3600]);
    // await network.provider.send("evm_mine");
    await expect(
      node.concludeTransaction(
        1, // user node
        signers[0].address,
        FILE_HASH,
        100, // file size
        BigInt(Math.round(TStart / 1000)), // timerStart
        BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
        10, // prove timeout length in seconds
        5, // conclude timeout
        BigInt(100), // segments count
        BigInt(1000), // bid Amount in wei
        {
          value: 1000,
        }
      )
    ).to.be.revertedWith("concludeTimeoutLength mismatch");
  });

  it("should revert with `proveTimeout mismatch'`", async function () {
    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: ethers.utils.parseEther("1.0"),
    });

    const FILE_HASH = ethers.utils.arrayify(randomBytes(32));
    const TStart = Date.now();
    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      10, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000) // bid Amount in wei
    );

    // await network.provider.send("evm_increaseTime", [3600]);
    // await network.provider.send("evm_mine");
    await expect(
      node.concludeTransaction(
        1, // user node
        signers[0].address,
        FILE_HASH,
        100, // file size
        BigInt(Math.round(TStart / 1000)), // timerStart
        BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
        5, // prove timeout length in seconds
        10, // conclude timeout
        BigInt(100), // segments count
        BigInt(1000), // bid Amount in wei
        {
          value: 1000,
        }
      )
    ).to.be.revertedWith("proveTimeout mismatch");
  });
});
