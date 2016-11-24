/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    var mapping = {
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    check = []
    
    for (let i = 0; i < s.length; i++){
        var c = s[i]
        if (mapping[c] !== undefined){
            if (check.length === 0 || check.slice(-1)[0] != mapping[c])
                return false;
            check.pop()
        } else {
            check.push(c)
        }
    }
    
    return check.length === 0
};
