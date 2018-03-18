class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        const int m = nums1.size();
        const int n = nums2.size();
        
        if (m > n) {
            // Insure that the size of nums1 always less or equal than another.
            return findMedianSortedArrays(nums2, nums1);
        }
     
        int k = (m + n + 1) / 2;
        
        int l = 0;
        int r = m;
        
        while(l < r) {
            int m1 = (l + r) / 2;
            int m2 = k - m1;

            // If m1 take a little big value (take the right side neighbor)
            // and it still less than m2 value,
            // means we take too small value of m1.
            if (nums1[m1] < nums2[m2 - 1]) {
                l = m1 + 1;
            } else {
                r = m1;
            }
        }
        
        int m1 = (l + r) / 2;
        int m2 = k - m1;
        
        // Check the boundary.
        int mid1 = max((m1 == 0)? INT_MIN : nums1[m1 - 1], (m2 == 0)? INT_MIN : nums2[m2 - 1]);
        
        // Odd length case.
        if ((m + n) % 2 == 1) {
            return mid1;
        }
        
        // Even length case.
        int mid2 = min((m1 >= m)? INT_MAX : nums1[m1], (m2 >= n)? INT_MAX : nums2[m2]);
        return (mid1 + mid2) / 2.0;
    }
};