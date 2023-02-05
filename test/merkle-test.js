const { expect, util } = require("chai");
const { BigNumber, utils } = require("ethers");
const { ethers, network } = require("hardhat");
const { randomBytes, sign } = require("crypto");
const { equal } = require("assert");
const { MerkleTree } = require("merkletreejs");
const { arrayify } = require("ethers/lib/utils");
const fs = require("fs");
const {
  generateMerkleTree,
  generateMerkleProof,
  generateMerkleRoot,
  generateMerkleRootFromMerkleTree,
} = require("../utils/merkle");

const hash = (data) => {
  return Buffer.from(
    arrayify(
      ethers.utils.solidityKeccak256(["bytes"], [ethers.utils.arrayify(data)])
    )
  );
};

describe("Merkle Test", function () {
  let node = null;
  let snapshotId = "";
  it("verifies merkle leaf", async function () {
    MerkleContract = await ethers.getContractFactory("Merkle");
    node = await MerkleContract.deploy();

    await node.deployed();

    const leavesLength = 322;
    const leaves = Array(leavesLength);
    for (let i = 0; i < leaves.length; i++) {
      leaves[i] = hash(randomBytes(50));
    }
    // console.log("Leaves");
    // console.log(leaves.map((l) => l.toString("hex")));
    const rootHash = generateMerkleRoot(leaves);
    const leafIndex = 300;

    const proofR = generateMerkleProof(leafIndex, leaves);
    const proof = proofR.map((m) => m.hash);
    const tree = generateMerkleTree(leaves);
    const rr = generateMerkleRootFromMerkleTree(proofR);
    // console.log("Tree");
    // console.log(tree.map((t) => t.map((x) => x.toString("hex"))));
    const directions = proofR.map((d) => (d.direction == "left" ? 0 : 1));
    console.log("Segment:", leafIndex);
    // console.log(
    //   proofR.map((m) => ({
    //     direction: m.direction,
    //     hash: m.hash.toString("hex"),
    //   }))
    // );

    console.log("Root:", rootHash.toString("hex"));
    console.log("Verified Root:", rr.toString("hex"));
    expect(
      await node.verify(
        proof,
        rootHash,
        leaves[leafIndex],
        BigInt(leafIndex),
        directions
      )
    ).to.equal(true, "this should be verified correctly as it's valid");
  });
  xit("verifies incorrect leaf", async function () {
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
  xit("test", async function () {
    const leaves = new Array();
    const f = fs.readFileSync("test.jpeg");

    const fstats = fs.statSync("test.jpeg");
    console.log("Computing hash");

    const fileSize = fstats.size;
    const segmentsCount = Math.ceil(fileSize / (1024 * 1));

    const lastChunkSize = fileSize % 1024;

    let chunkSize;

    if (segmentsCount === 1) {
      chunkSize = lastChunkSize;
    } else {
      chunkSize = Math.floor((fileSize - lastChunkSize) / (segmentsCount - 1));
    }
    console.log("Seg:", segmentsCount);

    let i;
    for (i = 0; i < segmentsCount - 1; i++) {
      const r = new Uint8Array(
        f.buffer.slice(i * chunkSize, i * chunkSize + chunkSize)
      );
      // r.
      // console.log(r.byteLength);
      leaves.push(hash(r));
    }
    if (segmentsCount > 1) {
      const r = new Uint8Array(
        f.buffer.slice(i * chunkSize, i * chunkSize + lastChunkSize)
      );
      // r.
      // console.log(r.byteLength);
      leaves.push(hash(r));
    }

    const tree = new MerkleTree(leaves, hash, {
      duplicateOdd: true,
      sort: false,
      // isBitcoinTree: true,
    });
    const rootHash = tree.getRoot();
    const leafIndex = 521;
    const proofData = tree.getProof(leaves[leafIndex], leafIndex);
    const proof = proofData.map((d) => ethers.utils.arrayify(d.data));

    console.log("Root: ", rootHash.toString("hex"));
    // console.log(tree.getLeaves().map((t) => t.toString("hex")));
    console.log("Proof");
    console.log(proofData.length);
    console.log(proofData.map((m) => m.data.toString("hex")));

    expect(
      await node.verify(
        proof,
        arrayify(rootHash),
        arrayify(leaves[leafIndex]),
        BigInt(leafIndex) // different index
      )
    ).to.equal(true, "this should validate");
  });
});
