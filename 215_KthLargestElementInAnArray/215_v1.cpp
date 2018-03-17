/*
    Based on quick sort that
    only partition the needed parts.
*/

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        partialSort(nums, k, 0, nums.size() - 1);
        return nums[k - 1];
    }

    void swap(vector<int>& nums, int i, int j) {
        int x = nums[i];
        nums[i] = nums[j];
        nums[j] = x;
    }

    void partialSort(vector<int>& nums, int k, int l, int r) {
        if (l < r) {
            int p = partition(nums, l, r);
            if (k - 1 < p) {
                partialSort(nums, k, l, p - 1);
            } else if (k - 1 > p) {
                partialSort(nums, k, p + 1, r);
            }
        }
    }

    int partition(vector<int>& nums, int l, int r) {
        int pivot = nums[r];
        int i = l - 1;
        for (int j = l; j < r; j++) {
            if (pivot <= nums[j]) {
                i++;
                swap(nums, i, j);
            }
        }
        swap(nums, i + 1, r);

        return i + 1;
    }
};