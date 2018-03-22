class Solution {
public:
    void dfs(int n, const vector<int>& nums, vector<int>& set, vector<vector<int>>& powerset) {
        if (n >= nums.size()) {
            powerset.push_back(set);
            return;
        }
        
        dfs(n + 1, nums, set, powerset);
        set.push_back(nums[n]);
        dfs(n + 1, nums, set, powerset);
        set.pop_back();
    }
    
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> powerset;
        vector<int> set;
        
        dfs(0, nums, set, powerset);
        
        return powerset;
    }
};