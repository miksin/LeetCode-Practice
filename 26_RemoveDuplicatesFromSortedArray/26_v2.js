/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  var i = 1;
  var startIndex = 0;
  while (nums.length > i) {
    if (nums[i] > nums[startIndex]) {
      if (i - startIndex - 1 > 0) {
        nums.splice(startIndex + 1, i - startIndex - 1);
      }
      startIndex += 1;
      i = startIndex;
    } else {
      i += 1;
    }
  }

  if (i - startIndex - 1 > 0) {
    nums.splice(startIndex + 1, i - startIndex - 1);
  }

  return nums.length;
};

var arr1 = [1, 1, 2]
var arr2 = [0,0,1,1,1,2,2,3,3,4]
var arr3 = [1, 1]
removeDuplicates(arr1)
removeDuplicates(arr2)
removeDuplicates(arr3)
console.log(arr1)
console.log(arr2)
console.log(arr3)
