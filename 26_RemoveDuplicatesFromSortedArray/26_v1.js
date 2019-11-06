/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  if (nums.length < 1) {
    return nums
  }

  var i = 1;
  var prev = nums[0]
  while (nums.length > i) {
    if (nums[i] > prev) {
      prev = nums[i]
      i += 1
    } else {
      nums.splice(i, 1)
    }
  }

  return nums.length
};

console.log(removeDuplicates([1, 1, 2]))
console.log(removeDuplicates([0,0,1,1,1,2,2,3,3,4]))
