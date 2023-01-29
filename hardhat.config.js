require("@nomiclabs/hardhat-waffle");

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: "0.8.4",
  networks: {
    ganache: {
      url: "http://127.0.0.1:7545",
      accounts: {
        mnemonic:
          "rural type globe skull action youth embrace globe river salmon enrich party",
        path: "m/44'/60'/0'/0/",
        initialIndex: 0,
        count: 10,
        passphrase: "",
      },
    },
    optimism_goerli: {
      url: "https://opt-goerli.g.alchemy.com/v2/lLvcVplpamR-7ZTitncsc_qSPkenwum9",
      accounts: [
        "c82b0a506338902b235d18cb3827e7fe8806aaf71d9994f0d771d849d660ee38",
      ],
    },
  },
};
