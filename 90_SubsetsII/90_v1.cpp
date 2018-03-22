class Solution {
public:
    void dfs(int n, const vector<int>& nums, vector<int>& set, vector<vector<int>>& powerset) {
        if (n >= nums.size()) {
            if (find(powerset.begin(), powerset.end(), set) == powerset.end()) {
                powerset.push_back(set);
            }            
            return;
        }
        
        set.push_back(nums[n]);
        dfs(n + 1, nums, set, powerset);
        set.pop_back();
        dfs(n + 1, nums, set, powerset); 
    }
    
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        
        vector<vector<int>> powerset;
        vector<int> set;
        
        dfs(0, nums, set, powerset);
        
        return powerset;
    }
};