// 创建一个名为Voting的合约，包含以下功能：
//- 一个mapping来存储候选人的得票数
//- 一个vote函数，允许用户投票给某个候选人
//- 一个getVotes函数，返回某个候选人的得票数
//- 一个resetVotes函数，重置所有候选人的得票数


// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Voting {
    // mapping 存储候选人的得票数
    mapping(string => uint256) private votes;

    // 存储候选人列表，方便 resetVotes 遍历
    string[] private candidates;

    // 记录候选人是否已经存在（避免重复添加）
    mapping(string => bool) private candidateExists;

    // 添加候选人（部署合约后，先调用这个）
    function addCandidate(string memory _name) public {
        require(!candidateExists[_name], "Candidate already exists");
        candidates.push(_name);
        candidateExists[_name] = true;
        votes[_name] = 0;
    }

    // 投票给某个候选人
    function vote(string memory _name) public {
        require(candidateExists[_name], "Candidate does not exist");
        votes[_name] += 1;
    }

    // 查询某个候选人的票数
    function getVotes(string memory _name) public view returns (uint256) {
        require(candidateExists[_name], "Candidate does not exist");
        return votes[_name];
    }

    // 重置所有候选人的票数
    function resetVotes() public {
        for (uint256 i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }

    // 获取所有候选人
    function getAllCandidates() public view returns (string[] memory) {
        return candidates;
    }
}
