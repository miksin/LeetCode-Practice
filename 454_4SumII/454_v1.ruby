# @param {Integer[]} a
# @param {Integer[]} b
# @param {Integer[]} c
# @param {Integer[]} d
# @return {Integer}
def four_sum_count(a, b, c, d)
  hashTable = Hash.new
  count = 0

  a.each do |i|
      b.each do |j|
          hashTable[i + j] ||= 0
          hashTable[i + j] += 1
      end
  end

  c.each do |i|
      d.each do |j|
          count += (hashTable[-i - j] || 0)
      end
  end

  count
end
