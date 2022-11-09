const { expect } = require("chai");
const { BigNumber, utils } = require("ethers");
const { ethers, network } = require("hardhat");
const { randomBytes } = require("crypto");

describe("Factory tests", function () {
  let snapshotId = -1;
  beforeEach(async () => {
    snapshotId = await network.provider.send("evm_snapshot");
  });
  afterEach(async () => {
    await network.provider.send("evm_revert", [snapshotId]);
  });
  it("creates node contract", async function () {
    const FactoryContract = await ethers.getContractFactory(
      "StorageNodeFactory"
    );
    const StorageNodeContract = await ethers.getContractFactory("StorageNode");
    const iface = FactoryContract.interface;
    const factory = await FactoryContract.deploy();

    await factory.deployed();

    const tx = await factory.createStorageContract(
      ethers.utils.arrayify("0x123456"), // TLSCert
      "http://localhost:8000" // HOST
    );

    const b = await tx.wait();

    const eventInd = b.events.findIndex(
      (f) => f.event == "EvNewStorageContract"
    );
    const event = b.events[eventInd];

    expect(event.args.addr).to.be.properAddress;
    // expect(computedKey).to.equal(validKey);
  });
  it("validates deployed nodes list", async function () {
    const FactoryContract = await ethers.getContractFactory(
      "StorageNodeFactory"
    );
    const StorageNodeContract = await ethers.getContractFactory("StorageNode");
    const iface = FactoryContract.interface;
    const factory = await FactoryContract.deploy();

    await factory.deployed();
    const deployContract = async () => {
      const tx = await factory.createStorageContract(
        ethers.utils.arrayify("0x123456"), // TLSCert
        "http://localhost:8000" // H
      );

      const b = await tx.wait();

      const eventInd = b.events.findIndex(
        (f) => f.event == "EvNewStorageContract"
      );
      const event = b.events[eventInd];

      return event.args.addr;
    };

    const address = [];
    const total = 5;
    let i = total;
    while (--i >= 0) {
      const addr = await deployContract();
      expect(addr).to.be.properAddress;
      address.push(addr);
    }

    const contractAddr = await factory.getStorageContracts();

    while (i++ < total - 1) {
      const a = await contractAddr[i];
      // console.log(address[i], "==", a);
      expect(address[i]).to.be.equal(a);
    }
  });
});
