/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    var unitLen = (numRows == 1)? 1: numRows*2 - 2
    var rows = new Array(numRows).fill('')
    
    for (let i = 0; i < s.length; i++){
        var unitIdx = i % unitLen
        var rowIdx = (unitIdx < numRows)? unitIdx: unitLen - unitIdx
        rows[rowIdx] += s[i]
    }
    
    return rows.join('')
};