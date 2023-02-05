// SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "hardhat/console.sol";

contract Merkle {
    function verify(
        bytes32[] memory proof,
        bytes32 root,
        bytes32 leaf,
        uint256 index,
        uint8[] memory directions
    ) public view returns (bool) {
        bytes32 hash = leaf;

        for (uint256 i = 0; i < proof.length - 1; i++) {
            bytes32 proofElement = proof[i + 1];

            if (directions[i + 1] == 1) {
                hash = keccak256(abi.encodePacked(hash, proofElement));
            } else {
                hash = keccak256(abi.encodePacked(proofElement, hash));
            }
            // console.logBytes32(hash);

            // index = index / 2;
        }

        return hash == root;

        // bytes32 computedHash = leaf;

        // for (uint256 i = 0; i < proof.length; i++) {
        //     bytes32 proofElement = proof[i];

        //     if (computedHash <= proofElement) {
        //         // Hash(current computed hash + current element of the proof)
        //         computedHash = keccak256(
        //             abi.encodePacked(computedHash, proofElement)
        //         );
        //     } else {
        //         // Hash(current element of the proof + current computed hash)
        //         computedHash = keccak256(
        //             abi.encodePacked(proofElement, computedHash)
        //         );
        //     }
        // }

        // // Check if the computed hash (root) is equal to the provided root
        // return computedHash == root;
    }
}
