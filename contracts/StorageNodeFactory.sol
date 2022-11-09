// SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "./StorageNode.sol";

contract StorageNodeFactory {
    StorageNode[] public storageContracts;

    event EvNewStorageContract(
        address addr,
        bytes TLSCert,
        string host,
        address owner
    );

    function getStorageContracts() public view returns (StorageNode[] memory) {
        return storageContracts;
    }

    function createStorageContract(
        bytes calldata _TLSCert,
        string calldata _host
    ) public returns (address) {
        StorageNode node = new StorageNode(_TLSCert, _host, msg.sender);
        storageContracts.push(node);

        emit EvNewStorageContract(address(node), _TLSCert, _host, msg.sender);
        return address(node);
    }

    /**
     * Todo: to be done
     */
    function removeContract(address addr) public {}
}
