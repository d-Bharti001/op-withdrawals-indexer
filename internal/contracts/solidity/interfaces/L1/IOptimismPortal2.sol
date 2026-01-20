// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IOptimismPortal2 {
    error ContentLengthMismatch();
    error EmptyItem();
    error InvalidDataRemainder();
    error InvalidHeader();
    error ReinitializableBase_ZeroInitVersion();
    error OptimismPortal_AlreadyFinalized();
    error OptimismPortal_BadTarget();
    error OptimismPortal_CallPaused();
    error OptimismPortal_CalldataTooLarge();
    error OptimismPortal_GasEstimation();
    error OptimismPortal_GasLimitTooLow();
    error OptimismPortal_ImproperDisputeGame();
    error OptimismPortal_InvalidDisputeGame();
    error OptimismPortal_InvalidMerkleProof();
    error OptimismPortal_InvalidOutputRootProof();
    error OptimismPortal_InvalidProofTimestamp();
    error OptimismPortal_InvalidRootClaim();
    error OptimismPortal_NoReentrancy();
    error OptimismPortal_ProofNotOldEnough();
    error OptimismPortal_Unproven();
    error OptimismPortal_InvalidLockboxState();
    error OutOfGas();
    error UnexpectedList();
    error UnexpectedString();

    event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData);
    event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success);
    event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to);
    event WithdrawalProvenExtension1(bytes32 indexed withdrawalHash, address indexed proofSubmitter);

    receive() external payable;

    function checkWithdrawal(bytes32 _withdrawalHash, address _proofSubmitter) external view;
    // function depositTransaction(
    //     address _to,
    //     uint256 _value,
    //     uint64 _gasLimit,
    //     bool _isCreation,
    //     bytes memory _data
    // )
    //     external
    //     payable;
    // function disputeGameFactory() external view returns (IDisputeGameFactory);
    // function disputeGameFinalityDelaySeconds() external view returns (uint256);
    // function superchainConfig() external view returns (ISuperchainConfig);
    // function finalizeWithdrawalTransaction(Types.WithdrawalTransaction memory _tx) external;
    // function finalizeWithdrawalTransactionExternalProof(
    //     Types.WithdrawalTransaction memory _tx,
    //     address _proofSubmitter
    // )
    //     external;
    function finalizedWithdrawals(bytes32) external view returns (bool);
    // function guardian() external view returns (address);
    // function l2Sender() external view returns (address);
    function numProofSubmitters(bytes32 _withdrawalHash) external view returns (uint256);
    function proofMaturityDelaySeconds() external view returns (uint256);
    function proofSubmitters(bytes32, uint256) external view returns (address);
    // function proveWithdrawalTransaction(
    //     Types.WithdrawalTransaction memory _tx,
    //     uint256 _disputeGameIndex,
    //     Types.OutputRootProof memory _outputRootProof,
    //     bytes[] memory _withdrawalProof
    // )
    //     external;
    // function provenWithdrawals(
    //     bytes32,
    //     address
    // )
    //     external
    //     view
    //     returns (IDisputeGame disputeGameProxy, uint64 timestamp);
    // function systemConfig() external view returns (ISystemConfig);
    function version() external pure returns (string memory);

    // function __constructor__(uint256 _proofMaturityDelaySeconds) external;
}
