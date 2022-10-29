const { expect } = require("chai");
const { BigNumber, utils } = require("ethers");
const { ethers, network } = require("hardhat");
const { randomBytes } = require("crypto");

describe("Transaction key tests", function () {
  it("validates compute function", async function () {
    const StorageNodeContract = await ethers.getContractFactory("StorageNode");
    const node = await StorageNodeContract.deploy(
      ethers.utils.arrayify("0x123456"), // TLSCert
      "http://localhost:8000" // HOST
    );

    await node.deployed();

    const signers = await ethers.getSigners();
    const fileMerkleRootHash = randomBytes(32);

    const userAddress = signers[0].address;
    const abi = utils.solidityPack(
      ["address", "bytes32"],
      [userAddress, ethers.utils.arrayify(fileMerkleRootHash)]
    );
    const validKey = utils.solidityKeccak256(["bytes"], [abi]);
    const computedKey = await node.computeKey(
      userAddress,
      ethers.utils.arrayify(fileMerkleRootHash)
    );

    expect(computedKey).to.equal(validKey);
  });
});
