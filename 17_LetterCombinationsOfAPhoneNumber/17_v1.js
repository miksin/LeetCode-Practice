/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function(digits) {
  if (digits.length < 1) {
      return []
  }
  
  var mapping = {
      2: ['a', 'b', 'c'],
      3: ['d', 'e', 'f'],
      4: ['g', 'h', 'i'],
      5: ['j', 'k', 'l'],
      6: ['m', 'n', 'o'],
      7: ['p', 'q', 'r', 's'],
      8: ['t', 'u', 'v'],
      9: ['w', 'x', 'y', 'z'],
  };
  
  var results = [];
  var stack = '';
  
  var combi = function(index) {
      if (index === digits.length) {
          results.push(stack);
          return;
      }
      
      mapping[digits[index]].forEach(function(l) {
          stack += l;
          combi(index + 1);
          stack = stack.slice(0, -1);
      });
  }
  
  combi(0);
  
  return results
};
