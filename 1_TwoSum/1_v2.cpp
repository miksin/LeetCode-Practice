class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> hashTable;
        for (int i = 0; i < nums.size(); i++){
            int x = nums[i];
            auto iter = hashTable.find(target - x);

            if (iter != hashTable.end()){
                return {iter->second, i};
            }

            hashTable[x] = i;
        }
        return vector<int>();
    }
};
