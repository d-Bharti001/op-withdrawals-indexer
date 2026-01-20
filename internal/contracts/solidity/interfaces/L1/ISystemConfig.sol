// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface ISystemConfig {
    enum UpdateType {
        BATCHER,
        FEE_SCALARS,
        GAS_LIMIT,
        UNSAFE_BLOCK_SIGNER,
        EIP_1559_PARAMS,
        OPERATOR_FEE_PARAMS,
        MIN_BASE_FEE,
        DA_FOOTPRINT_GAS_SCALAR
    }

    struct Addresses {
        address l1CrossDomainMessenger;
        address l1ERC721Bridge;
        address l1StandardBridge;
        address optimismPortal;
        address optimismMintableERC20Factory;
    }

    event ConfigUpdate(uint256 indexed version, UpdateType indexed updateType, bytes data);
    event FeatureSet(bytes32 indexed feature, bool indexed enabled);

    function VERSION() external view returns (uint256);
    // function basefeeScalar() external view returns (uint32);
    function batchInbox() external view returns (address addr_);
    function batcherHash() external view returns (bytes32);
    // function blobbasefeeScalar() external view returns (uint32);
    function disputeGameFactory() external view returns (address addr_);
    // function gasLimit() external view returns (uint64);
    // function eip1559Denominator() external view returns (uint32);
    // function eip1559Elasticity() external view returns (uint32);
    function getAddresses() external view returns (Addresses memory);
    function l1CrossDomainMessenger() external view returns (address addr_);
    function l1ERC721Bridge() external view returns (address addr_);
    function l1StandardBridge() external view returns (address addr_);
    function l2ChainId() external view returns (uint256);
    // function maximumGasLimit() external pure returns (uint64);
    // function minimumGasLimit() external view returns (uint64);
    // function operatorFeeConstant() external view returns (uint64);
    // function operatorFeeScalar() external view returns (uint32);
    // function minBaseFee() external view returns (uint64);
    // function daFootprintGasScalar() external view returns (uint16);
    function optimismMintableERC20Factory() external view returns (address addr_);
    function optimismPortal() external view returns (address addr_);
    function overhead() external view returns (uint256);
    function owner() external view returns (address);
    function startBlock() external view returns (uint256 startBlock_);
    function version() external pure returns (string memory);
    function paused() external view returns (bool);
    // function superchainConfig() external view returns (ISuperchainConfig);
    function guardian() external view returns (address);
    function isFeatureEnabled(bytes32) external view returns (bool);
}
