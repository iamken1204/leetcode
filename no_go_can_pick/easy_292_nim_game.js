/**
 * Easy 292. Nim Game
 * https://leetcode.com/problems/nim-game/
 *
 * @param {number} n
 * @return {boolean}
 */
var canWinNim = function(n) {
    let flag = 3 + 1
    return n % flag !== 0
}
