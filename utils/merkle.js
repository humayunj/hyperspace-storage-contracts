const crypto = require("crypto");

const LEFT = "left";
const RIGHT = "right";

const hashes = ["123", "123", "123", "123", "123"];
const keccak256 = require("keccak256");

const getLeafNodeDirectionInMerkleTree = (hash, merkleTree) => {
  const hashIndex = merkleTree[0].findIndex((h) => h === hash);
  return hashIndex % 2 === 0 ? LEFT : RIGHT;
};

function ensureEven(hashes) {
  if (hashes.length % 2 !== 0) {
    hashes.push(Buffer.from(hashes[hashes.length - 1]));
  }
}

function generateMerkleRoot(hashes) {
  if (!hashes || hashes.length === 0) {
    return null;
  }
  ensureEven(hashes);
  const combinedHashes = [];
  for (let i = 0; i < hashes.length; i += 2) {
    // console.log(hashes[i], hashes[i + 1]);
    const hashConcat = Buffer.concat([hashes[i], hashes[i + 1]]);
    // console.log(">", hashConcat);
    const hash = keccak256(hashConcat);
    combinedHashes.push(hash);
  }
  if (combinedHashes.length === 1) {
    return combinedHashes[0];
  }
  return generateMerkleRoot(combinedHashes);
}

function generateMerkleTree(hashes) {
  if (!hashes || hashes.length === 0) {
    return "";
  }
  const tree = [hashes];

  const generate = (hashes, tree) => {
    if (hashes.length == 1) return hashes;

    ensureEven(hashes);

    const combinedHashes = [];
    for (let i = 0; i < hashes.length; i += 2) {
      const hashConcat = Buffer.concat([hashes[i], hashes[i + 1]]);
      const hash = keccak256(hashConcat);
      combinedHashes.push(hash);
    }
    tree.push(combinedHashes);
    return generate(combinedHashes, tree);
  };
  generate(hashes, tree);
  return tree;
}

function generateMerkleProof(index, hashes) {
  if (!hashes || index < 0 || hashes.length == 0) {
    return null;
  }
  const tree = generateMerkleTree(hashes);
  console.log(tree);
  const merkleProof = [
    {
      hash: hashes[index],
      direction: index % 2 == 0 ? LEFT : RIGHT,
    },
  ];

  let hashIndex = index;
  for (let level = 0; level < tree.length - 1; level++) {
    const isLeftChild = hashIndex % 2 === 0;
    const siblingDirection = isLeftChild ? RIGHT : LEFT;
    const siblingIndex = isLeftChild ? hashIndex + 1 : hashIndex - 1;
    const siblingNode = {
      hash: tree[level][siblingIndex],
      direction: siblingDirection,
    };
    console.log("Tree([Level]length)", tree[level].length);
    if (siblingIndex >= tree[level].length) {
      console.log("sibling out of range: ", siblingIndex, tree[level].length);
      return null;
    }
    merkleProof.push(siblingNode);
    hashIndex = Math.floor(hashIndex / 2);
  }
  return merkleProof;
}

function generateMerkleRootFromMerkleTree(merkleProof) {
  if (!merkleProof || merkleProof.length == 0) {
    return "";
  }
  const merkleRootFromProof = merkleProof.reduce((hashProof1, hashProof2) => {
    if (hashProof2.direction === RIGHT) {
      const hash = keccak256(Buffer.concat([hashProof1.hash, hashProof2.hash]));
      //   console.log(">", hash);
      return { hash };
    }
    const hash = keccak256(Buffer.concat([hashProof2.hash, hashProof1.hash]));
    // console.log(">", hash);

    return { hash };
  });
  return merkleRootFromProof.hash;
}

function generateMerkleRootFromMerkleTreeDir(path, dir) {
  if (!path || path.length == 0) {
    return null;
  }
  const merkleRootFromProof = path.reduce((hashProof1, hashProof2, i) => {
    console.log(hashProof1);
    if (dir[i] === 1) {
      const hash = keccak256(Buffer.concat([hashProof1, hashProof2]));
      //   console.log(">", hash);
      return hash;
    }
    const hash = keccak256(Buffer.concat([hashProof2, hashProof1]));
    // console.log(">", hash);

    return hash;
  });
  return merkleRootFromProof;
}
module.exports = {
  generateMerkleProof,
  generateMerkleRoot,
  generateMerkleRootFromMerkleTree,
  generateMerkleTree,
};

// const merkleRoot = generateMerkleRoot(hashes);
// console.log("merkleRoot: ", merkleRoot);

// const merkleTree = generateMerkleTree(hashes);
// console.log("MerkleTree\n", merkleTree);

// const merkleProof = generateMerkleProof(hashes[2], hashes);
// console.log("MerkleProof\n", merkleProof);
// const root = generateMerkleRootFromMerkleTree(merkleProof);
// console.log("root\n", root);

// console.log(
//   ":",
//   generateMerkleRoot([
//     Buffer.from("hello", "utf-8"),
//     Buffer.from("world!"),
//     Buffer.from("hello", "utf-8"),
//     Buffer.from("hello", "utf-8"),
//     Buffer.from("hello", "utf-8"),
//   ])
// );

console.log(
  ">",
  generateMerkleRootFromMerkleTreeDir(
    [
      "9f12c1d8cc60092c50c808fd4d52ad3ff83924d3ddac23c5e1cd8ab8f6603d6e",
      "d3727d3ef00bb92430637883a480ec59b9c4d0a07d95a8a14115be36b83f1928",
      "5e88896fdec0726487be0a96098afc96619354054f48e06a7da238cacb4cda7a",
      "bf4554c805820a0d1b0e923cb621982a1c8c1db48a24499a012bcd1baff4e0d7",
      "7e671f30031fb5ce4649a4b6d0043b965a986d1a7c42e837b71fdcf584e1b5b7",
      "53300ad31fe4bc179997071d7c8aaa1703f91f84bcc10e4b0780edae4e656af1",
      "5e4b4d72601bce80f02ae2d3146ca250dd33045f2a244ac9ddfd98bd7db607e2",
      "c90b3bc273b89814dc2c066c87a3e512aeec0ab4d4ae08de88ca64445e29ef8a",
      "506598e4a09198634a0041c2b10431aaa9096ef00b0b7e47cfa457e11ff5025b",
      "8d1e33ade261ae3f59cd3a619c0efa5290f0dc9e277ce0034b5b918a39acd400",
      "9bd45d86ad6e95e38d9bdc010c09e3800c75106359aa9feca457eaccee8d66f4",
    ].map((b) => Buffer.from(b, "hex")),
    [0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
  )
);
