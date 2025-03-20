// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

contract DIDHashRegistry {
    mapping(string => string) public didToHash; // DID to SHA256 Encryption
    
    event HashUpdated(string indexed did, string hash);

    function setHash(string calldata did, string calldata hash) external {
        didToHash[did] = hash;
        emit HashUpdated(did, hash);
    }

    function getHash(string calldata did) external view returns (string memory) {
        return didToHash[did];
    }
}