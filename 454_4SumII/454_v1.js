/**
 * @param {number[]} A
 * @param {number[]} B
 * @param {number[]} C
 * @param {number[]} D
 * @return {number}
 */
var fourSumCount = function(A, B, C, D) {
  const hashTable = {}
  let count = 0

  for (let i = 0; i < A.length; i++) {
      for (let j = 0; j < B.length; j++) {
          const value = A[i] + B[j]
          hashTable[value] = (hashTable[value] || 0) + 1
      }
  }

  for (let i = 0; i < C.length; i++) {
      for (let j = 0; j < D.length; j++) {
          const value = C[i] + D[j]
          count += (hashTable[-value] || 0)
      }
  }

  return count
};
