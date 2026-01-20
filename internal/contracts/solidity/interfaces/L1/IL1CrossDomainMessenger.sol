// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ICrossDomainMessenger} from "../universal/ICrossDomainMessenger.sol";

interface IL1CrossDomainMessenger is ICrossDomainMessenger {

    // function PORTAL() external view returns (IOptimismPortal);
    // function portal() external view returns (IOptimismPortal);
    // function systemConfig() external view returns (ISystemConfig);
    function version() external view returns (string memory);
    // function superchainConfig() external view returns (ISuperchainConfig);

    // function __constructor__() external;
}
