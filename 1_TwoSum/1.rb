# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    hashTable = Hash.new
    
    nums.each_with_index do |value, index|
        find_value = target - value
        
        if hashTable.key?(find_value)
            return [hashTable[find_value], index]
        end
        
        hashTable[value] = index
    end
    []
end