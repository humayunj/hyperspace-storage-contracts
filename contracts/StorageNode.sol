//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

// Version PROTOTYPE: 1.1

// import "hardhat/console.sol";

contract StorageNode {
    bytes public TLSCert;
    string public HOST;

    uint256 public lockedCollateral = 0;

    struct Transaction {
        bytes20 merkleRootHash;
        uint32 size;
        uint256 timerStart;
        uint256 timerEnd;
        uint64 proveTimeoutLength;
        uint32 segmentsCount;
        bool userConcluded;
        uint64 concludeTimeoutLength;
        uint256 bidAmount;
    }
    uint256 public mappingLength = 0;
    // Contains the address list (keys) of the mappings
    address[] public mappingsList;

    //Store mapping address => Transaction
    mapping(address => Transaction) transactionMapping; // single address => single-file

    enum CallerType {
        StorageNode,
        ClientNode
    }

    constructor(bytes memory _TLSCert, string memory _HOST) {
        TLSCert = _TLSCert;
        HOST = _HOST;
    }

    /**
     * @dev Invoked by Storage Node first. Storage node submits uploaded
     *  file details, which are then added to a map. Client Node must invoke
     *  this function with similar params to ensure consistency within
     *  `timerstart + concludeTimeoutLength` time period, the transaction will
     *  be expired otherwise.
     */
    function concludeTransaction(
        CallerType callerType,
        address userAddress,
        bytes20 merkleRootHash,
        uint32 fileSize,
        uint256 timerStart,
        uint256 timerEnd,
        uint64 proveTimeoutLength,
        uint64 concludeTimeoutLength,
        uint32 segmentsCount,
        uint256 bidAmount
    ) public payable {
        if (callerType == CallerType.StorageNode) {
            /**
             * Proposal: transationMapping key = keccak256(concat(userAddress,merkleRootHash))
             */
            Transaction storage t = transactionMapping[userAddress];

            // check if we already have file
            require(t.size == 0, "file already stored");
            require(
                timerEnd > block.timestamp,
                "timerEnd must be > current timestamp"
            );
            require(
                address(this).balance >= bidAmount * 2,
                "insufficient collateral"
            );

            t.merkleRootHash = merkleRootHash;
            t.size = fileSize;
            t.segmentsCount = segmentsCount;
            t.timerStart = timerStart;
            t.timerEnd = timerEnd;
            t.proveTimeoutLength = proveTimeoutLength;
            t.userConcluded = false;
            t.concludeTimeoutLength = concludeTimeoutLength;
            t.bidAmount = bidAmount;

            mappingLength += 1;
        } else if (callerType == CallerType.ClientNode) {
            Transaction storage t = transactionMapping[msg.sender];
            require(t.merkleRootHash == merkleRootHash, "root mismatch");
            require(t.size == fileSize, "size mismatch");
            require(t.segmentsCount == segmentsCount, "segments mismatch");
            require(t.timerStart == timerStart, "timestart mismatch");
            require(t.timerEnd == timerEnd, "timeend mismatch");
            require(
                t.proveTimeoutLength == proveTimeoutLength,
                "proveTimeout mismatch"
            );
            require(
                t.concludeTimeoutLength == concludeTimeoutLength,
                "concludeTimeoutLength mismatch"
            );
            require(t.bidAmount == bidAmount, "bid amount mismatch");
            require(t.userConcluded == false, "user concluded already");
            require(msg.value >= bidAmount, "amount must be >= bidAmount"); // refund policy?
            require(
                block.timestamp <= (t.timerStart + t.concludeTimeoutLength),
                "conclude Timed out"
            ); // todo: remove `msg.sender` from transactionMapping[]
            require(
                address(this).balance >= bidAmount * 2,
                "insufficient collateral"
            );
            lockedCollateral += bidAmount * 2;

            t.userConcluded = true;

            //return remaining amount
            payable(address(msg.sender)).transfer(msg.value - bidAmount);
        }
    }

    /**
     * @dev Invoked by Storage Node to finish the transaction and unlocking the
     * `collateral + reward`. The block timestamp must be greater than or equal to
     *  the agreed timerEnd. Storage Node will no longer be liable for this file storage.
     */
    function finishTransaction(address userAddress) public {
        Transaction storage t = transactionMapping[userAddress];
        require(t.size > 0, "invalid opr");
        require(t.userConcluded == true, "invalid tx");
        require(block.timestamp >= t.timerEnd, "not expired");

        lockedCollateral -= t.bidAmount * 2;

        t.size = 0; //reset slot
    }

    /**
     * @dev To be invoked by Storage Node to withdraw specific non-collateral amount.
     */
    function withdraw(uint256 amount, address target) public {
        uint256 b = address(this).balance - lockedCollateral;
        require(b >= amount, "insufficent blnc");
        payable(target).transfer(amount);
    }

    // TODO
    /**
     * @dev Invoked by Client Node to verify file storage with in specific time period.
     */
    function validateStorage() public {}

    //  TODO
    /**
     * @dev Invoked by Storage Node to submit storage proof.
     */
    function processProof() public {}

    /**
     * @dev Fallback function to receive ether ( deposit )
     */
    receive() external payable {
        // React to receiving ether
    }
}
