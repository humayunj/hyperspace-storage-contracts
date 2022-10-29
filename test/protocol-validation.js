const { expect } = require("chai");
const { BigNumber } = require("ethers");
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

describe("Validation Protocol", function () {
  let node = null;
  let snapshotId = "";
  before(async () => {});
  beforeEach(async () => {
    snapshotId = await network.provider.send("evm_snapshot");
    // console.log("snapshot:", snapshotId);
    StorageNodeContract = await ethers.getContractFactory("StorageNode");
    node = await StorageNodeContract.deploy(
      ethers.utils.arrayify("0x123456"), // TLSCert
      "http://localhost:8000" // HOST
    );

    await node.deployed();
  });
  afterEach(async () => {
    await network.provider.send("evm_revert", [snapshotId]);
  });
  it("should trigger validation protocol", async function () {
    const PROVE_TIMEOUT = 10;

    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: ethers.utils.parseEther("1.0"),
    });

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
      PROVE_TIMEOUT, // prove timeout length in seconds
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
      PROVE_TIMEOUT, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000), // bid Amount in wei
      {
        value: 1000,
      }
    );

    const t = node.validateStorage(signers[0].address, FILE_HASH, 43);
    const r = await t;
    const b = await r.wait();

    const eventInd = b.events.findIndex((f) => f.event == "EvProveStorage");
    const event = b.events[eventInd];
    let nodeTime = (await event.getBlock()).timestamp;

    expect(event).to.not.null;

    expect(event.args.length).to.equal(5, "event args length mismatch");
    expect(ethers.utils.hexlify(event.args[0])).to.equal(
      signers[0].address.toLowerCase()
    );
    expect(ethers.utils.hexlify(event.args[1])).to.equal(
      ethers.utils.hexlify(FILE_HASH).toLowerCase()
    );
    expect(event.args[2].toString()).to.equal(BigInt(nodeTime).toString());
    expect(event.args[3].toString()).to.equal(
      BigInt(nodeTime + PROVE_TIMEOUT).toString()
    );
    expect(event.args[4].toString()).to.equal(BigInt(43).toString());
  });

  it("should prove storage", async function () {
    const PROVE_TIMEOUT = 10;

    const signers = await ethers.getSigners();
    await signers[0].sendTransaction({
      to: node.address,
      value: ethers.utils.parseEther("1.0"),
    });

    // console.log("BL:", await signers[0].getBalance());
    const FILE_SIZE = 1 * 1024 * 1024 * 1024; // 1 GB
    const LEAF_SIZE = 10 * 1024; // 10 KB
    const LEAVES_LENGTH = Math.ceil(FILE_SIZE / LEAF_SIZE); //

    const leaves = Array(LEAVES_LENGTH);
    const LEAF_INDEX = 43;
    const SLICE = randomBytes(LEAF_SIZE);
    const bytes = randomBytes(LEAF_SIZE);
    const bytesHash = hash(bytes);

    for (let i = 0; i < LEAVES_LENGTH; i++) {
      leaves[i] = bytesHash;
      // console.log("Generateing leave", i);
    }
    leaves[LEAF_INDEX] = SLICE;

    const tree = new MerkleTree(leaves, hash);
    const rootHash = tree.getRoot();

    const proofData = tree.getProof(leaves[LEAF_INDEX], LEAF_INDEX);
    const proof = proofData.map((d) => ethers.utils.arrayify(d.data));

    const FILE_HASH = ethers.utils.arrayify(rootHash);

    const TStart = Date.now();

    await node.concludeTransaction(
      0, // storage node
      signers[0].address,
      FILE_HASH, // merkle hash
      100, // file size in bytes
      BigInt(Math.round(TStart / 1000)), // timerStart
      BigInt(Math.round((TStart + 10000) / 1000)), // timerEnd
      PROVE_TIMEOUT, // prove timeout length in seconds
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
      PROVE_TIMEOUT, // prove timeout length in seconds
      10, // conclude timeout
      BigInt(100), // segments count
      BigInt(1000), // bid Amount in wei
      {
        value: 1000,
      }
    );

    const t = node.validateStorage(signers[0].address, FILE_HASH, LEAF_INDEX);
    const r = await t;
    const b = await r.wait();

    const eventInd = b.events.findIndex((f) => f.event == "EvProveStorage");
    const event = b.events[eventInd];
    let nodeTime = (await event.getBlock()).timestamp;

    expect(event).to.not.null;

    expect(event.args.length).to.equal(5, "event args length mismatch");

    const userAddress = event.args[0];
    const fileMerkleRoot = event.args[1];
    const timestamp = event.args[2];
    const expiryTimestamp = event.args[3];
    const segmentInd = event.args[4];
  });
});
