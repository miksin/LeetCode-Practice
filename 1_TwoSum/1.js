/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var hashTable = new Map();

    for (let i = 0; i < nums.length; i++) {
        var x = nums[i];
        var get_value = hashTable.get(target - x);

        if (get_value !== undefined) {
            return [get_value, i];
        }

        hashTable.set(x, i);
    }

    return []
};