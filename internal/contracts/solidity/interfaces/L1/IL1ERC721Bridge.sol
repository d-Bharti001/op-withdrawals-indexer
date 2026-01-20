// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC721Bridge} from "../universal/IERC721Bridge.sol";

interface IL1ERC721Bridge is IERC721Bridge {

    function bridgeERC721(
        address _localToken,
        address _remoteToken,
        uint256 _tokenId,
        uint32 _minGasLimit,
        bytes memory _extraData
    )
        external;
    function bridgeERC721To(
        address _localToken,
        address _remoteToken,
        address _to,
        uint256 _tokenId,
        uint32 _minGasLimit,
        bytes memory _extraData
    )
        external;
    function deposits(address, address, uint256) external view returns (bool);
    function finalizeBridgeERC721(
        address _localToken,
        address _remoteToken,
        address _from,
        address _to,
        uint256 _tokenId,
        bytes memory _extraData
    )
        external;
    // function paused() external view returns (bool);
    // function systemConfig() external view returns (ISystemConfig);
    function version() external view returns (string memory);
    // function superchainConfig() external view returns (ISuperchainConfig);
}
