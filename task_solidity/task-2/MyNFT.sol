// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721URIStorage, Ownable {
    uint256 private _nextTokenId;

    constructor(string memory name_, string memory symbol_) ERC721(name_, symbol_) Ownable(msg.sender) {}

    /// @notice 仅所有者可以铸造，并设置 tokenURI（一般为 ipfs://CID/xxx.json）
    function safeMint(address to, string memory tokenURI_) external onlyOwner returns (uint256 tokenId) {
        tokenId = ++_nextTokenId;
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, tokenURI_);
    }

    /// @notice 批量铸造（可选）
    function batchMint(address to, string[] calldata uris) external onlyOwner {
        for (uint256 i = 0; i < uris.length; i++) {
            uint256 tokenId = ++_nextTokenId;
            _safeMint(to, tokenId);
            _setTokenURI(tokenId, uris[i]);
        }
    }

    function nextTokenId() external view returns (uint256) {
        return _nextTokenId + 1;
    }
}
