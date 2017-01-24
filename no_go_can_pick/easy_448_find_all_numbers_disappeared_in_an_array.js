/**
 * Easy 448. Find All Numbers Disappeared in an Array
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
 *
 * @param {number[]} nums
 * @return {number[]}
 */
var findDisappearedNumbers = function(nums) {
    let res = []
    let numsLen = nums.length
    let container = new Set()

    nums.forEach((num) => {
        container.add(num)
    });

    let arr = Array.from(container)
    let max = Math.max(...arr, numsLen)

    for (let i=1; i<=max; i++) {
        if (!container.has(i)) {
            res.push(i)
        }
    }

    return res
}
