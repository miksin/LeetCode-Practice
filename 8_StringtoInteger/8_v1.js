/**
 * @param {string} str
 * @return {number}
 */
var myAtoi = function(str) {
    var INT_MAX = 2147483647
    var INT_MIN = -2147483648
    
    var m = /^\s*([\+-]?\d+).*$/.exec(str)
    
    if (m === null)
        return 0
        
    var result = parseInt(m[0])
    
    return result > 0? Math.min(result, INT_MAX) : Math.max(result, INT_MIN)
};
