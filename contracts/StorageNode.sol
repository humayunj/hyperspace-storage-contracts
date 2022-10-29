//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;
import "./Merkle.sol";

// Version PROTOTYPE: 1

// import "hardhat/console.sol";

contract StorageNode is Merkle {
    bytes public TLSCert;
    string public HOST;

    uint256 public lockedCollateral = 0;

    struct Transaction {
        bytes32 merkleRootHash;
        uint32 size;
        uint256 timerStart; // UNIX timestamp (in seconds)
        uint256 timerEnd;
        uint64 proveTimeoutLength; // in seconds
        uint32 segmentsCount;
        bool userConcluded;
        uint64 concludeTimeoutLength;
        uint256 bidAmount;
        uint256 validationRequestTime;
        uint32 validationSegmentInd;
    }
    uint256 public mappingLength = 0;
    // Contains the address list (keys) of the mappings
    bytes32[] public mappingsList;

    //Store mapping address => Transaction
    mapping(bytes32 => Transaction) transactionMapping; // single address => single-file

    enum CallerType {
        StorageNode,
        ClientNode
    }
    /**
     * @dev Generated for storage node to begin validation protocol
     */
    event EvProveStorage(
        address userAddress,
        bytes32 fileMerkleRootHash,
        uint256 timestamp,
        uint256 expiryTimestamp,
        uint32 segmentIndex
    );

    constructor(bytes memory _TLSCert, string memory _HOST) {
        TLSCert = _TLSCert;
        HOST = _HOST;
    }

    function computeKey(address userAddress, bytes32 merkleRootHash)
        public
        pure
        returns (bytes32)
    {
        return keccak256(abi.encodePacked(userAddress, merkleRootHash));
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
        bytes32 merkleRootHash,
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
            bytes32 ref = computeKey(userAddress, merkleRootHash);
            Transaction storage t = transactionMapping[ref];

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
            t.validationRequestTime = 0;
            t.validationSegmentInd = 0;

            mappingLength += 1;
        } else if (callerType == CallerType.ClientNode) {
            bytes32 ref = computeKey(msg.sender, merkleRootHash);
            Transaction storage t = transactionMapping[ref];
            require(t.merkleRootHash == merkleRootHash, "root mismatch");
            require(t.size == fileSize, "size mismatch");
            require(t.segmentsCount == segmentsCount, "segments mismatch");
            require(t.timerStart == timerStart, "timerstart mismatch");
            require(t.timerEnd == timerEnd, "timerend mismatch");
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
    function finishTransaction(address userAddress, bytes32 merkleRootHash)
        public
    {
        bytes32 ref = computeKey(userAddress, merkleRootHash);
        Transaction storage t = transactionMapping[ref];

        require(t.size > 0, "invalid opr");
        require(t.userConcluded == true, "invalid tx");
        require(block.timestamp >= t.timerEnd, "not expired");

        lockedCollateral -= t.bidAmount * 2;
        // To-do: emit finish event

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
    function validateStorage(
        address userAddress,
        bytes32 fileRootHash,
        uint32 segmentIndex
    ) public {
        bytes32 ref = computeKey(userAddress, fileRootHash);
        Transaction storage t = transactionMapping[ref];
        require(t.size > 0, "mapping doesn't exists");
        require(t.userConcluded == true, "user hasn't concluded");
        require(
            segmentIndex >= 0 && segmentIndex < t.segmentsCount,
            "invalid segmentIndex"
        );
        require(
            block.timestamp > (t.validationRequestTime + t.proveTimeoutLength),
            "validation is already in progress"
        );
        t.validationRequestTime = block.timestamp;
        t.validationSegmentInd = segmentIndex;
        emit EvProveStorage(
            userAddress,
            fileRootHash,
            (block.timestamp),
            (block.timestamp + uint256(t.proveTimeoutLength)),
            segmentIndex
        );
    }

    /// TODO
    /**
     * @dev Invoked by Storage Node to submit storage proof.
     */
    function processValidation(
        address userAddress,
        bytes32 rootHash,
        bytes calldata data,
        bytes32[] calldata proof
    ) public returns (bool) {
        bytes32 ref = computeKey(userAddress, rootHash);
        Transaction storage t = transactionMapping[ref];
        require(t.size > 0, "invalid tx");
        require(t.userConcluded == true, "tx not concluded");
        uint32 segmentInd = t.validationSegmentInd;
        bytes32 leafHash = keccak256(data);

        bool isValid = verify(proof, rootHash, leafHash, segmentInd);
        if (isValid) {
            /// Todo: emit validated event
            t.validationRequestTime = 0;
        }
        return isValid;
    }

    /**
     * @dev invoked by client node
     */
    function validationExpired(address userAddress, bytes32 rootHash) public {
        bytes32 ref = computeKey(userAddress, rootHash);
        Transaction storage t = transactionMapping[ref];
        require(t.size > 0, "invalid tx");
        require(t.userConcluded == true, "tx not concluded");
        require(t.validationRequestTime > 0, "validation not started");
        require(
            block.timestamp < t.validationRequestTime + t.proveTimeoutLength,
            "validation window not expired"
        );
        /// Todo: release amount to user

        uint256 transferAmount = t.bidAmount * 2; //twice
        lockedCollateral -= transferAmount;
        payable(userAddress).transfer(transferAmount);

        /**
         * Todo: emit validation expired event
         */
        t.size = 0; // remove tx
        t.userConcluded = false;
    }

    /**
     * @dev Fallback function to receive ether ( deposit )
     */
    receive() external payable {
        // React to receiving ether
    }
}
