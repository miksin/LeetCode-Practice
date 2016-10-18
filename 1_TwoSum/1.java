public class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();

        for (int i = 0; i < nums.length; i++){
            int x = nums[i];
            Integer get_value = map.get(target - x);

            if (get_value != null){
                int[] return_value = {get_value, i};
                return return_value;
            }

            map.put(x, i);
        }

        int [] null_solution = {};
        return null_solution;
    }
}