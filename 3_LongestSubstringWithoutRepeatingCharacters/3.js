/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    "use strict";

    var max_len = 0
    var start = 0
    var char_set = new Map()

    for (let i = 0; i < s.length; i++){
        var c = s.charAt(i)

        if (char_set.has(c)){
            var new_start = char_set.get(c) + 1
            while (start < new_start){
                char_set.delete(s.charAt(start++))
            }
        }

        max_len = Math.max(i - start + 1, max_len)
        char_set.set(c, i)
    }

    return max_len
};