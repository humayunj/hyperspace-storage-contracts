const { expect, util } = require("chai");
const { BigNumber, utils } = require("ethers");
const { ethers, network } = require("hardhat");
const { randomBytes, sign } = require("crypto");
const { equal } = require("assert");
const { MerkleTree } = require("merkletreejs");
const { arrayify } = require("ethers/lib/utils");

const hash = (data) => {
  return ethers.utils.solidityKeccak256(
    ["bytes"],
    [ethers.utils.arrayify(data)]
  );
};

describe("Merkle Test", function () {
  let node = null;
  let snapshotId = "";
  it("verifies merkle leaf", async function () {
    MerkleContract = await ethers.getContractFactory("Merkle");
    node = await MerkleContract.deploy();

    await node.deployed();

    const leavesLength = 4;
    const leaves = Array(leavesLength);
    for (let i = 0; i < leaves.length; i++) {
      leaves[i] = hash(randomBytes(50));
    }
    const tree = new MerkleTree(leaves, hash);
    const rootHash = tree.getRoot();
    const leafIndex = 1;
    const proofData = tree.getProof(leaves[leafIndex], leafIndex);

    const proof = proofData.map((d) => ethers.utils.arrayify(d.data));

    expect(
      await node.verify(
        proof,
        arrayify(rootHash),
        arrayify(leaves[leafIndex]),
        BigInt(leafIndex)
      )
    ).to.equal(true, "this should be verified correctly as it's valid");
  });
  it("verifies incorrect leaf", async function () {
    MerkleContract = await ethers.getContractFactory("Merkle");
    node = await MerkleContract.deploy();

    await node.deployed();

    const leavesLength = 4;
    const leaves = Array(leavesLength);
    for (let i = 0; i < leaves.length; i++) {
      leaves[i] = hash(randomBytes(50));
    }
    const tree = new MerkleTree(leaves, hash);
    const rootHash = tree.getRoot();
    const leafIndex = 1;
    const proofData = tree.getProof(leaves[leafIndex], leafIndex);
    const proof = proofData.map((d) => ethers.utils.arrayify(d.data));

    expect(
      await node.verify(
        proof,
        arrayify(rootHash),
        arrayify(leaves[leafIndex]),
        BigInt(leafIndex + 1) // different index
      )
    ).to.equal(false, "this should't validate");
  });
});
