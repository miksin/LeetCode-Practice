/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    var reverseInt = Number(x.toString().split("").reverse().join(""))
    return !isNaN(reverseInt) && reverseInt == x
};
