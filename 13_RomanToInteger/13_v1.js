/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    var alpha = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    var buf = 0, sum = 0
    
    for (let i = 0; i < s.length; i++){
        var num = alpha[s[i]]
        if (buf === 0){
            buf = num
        } else if (buf < num){
            sum = sum - buf + num
            buf = 0
        } else {
            sum += buf
            buf = num
        }
        console.log(sum)
    }
    
    sum += buf
    return sum
};
