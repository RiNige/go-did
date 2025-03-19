// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AccountRegistry {
    mapping(address => bool) public accounts;

    event AccountRegistered(address indexed account);

    function registerAccount() public {
        require(!accounts[msg.sender], "Account already registered");
        accounts[msg.sender] = true;
        emit AccountRegistered(msg.sender);
    }

    function isAccountRegistered(address account) public view returns (bool) {
        return accounts[account];
    }
}