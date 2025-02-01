
class Solution198 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // int[] nums = {1, 2, 3, 1};
    int[] nums = {100, 1, 2, 104};
    System.out.println(sol.rob(nums));
  }
}

class Solution {
  public int rob(int[] nums) {
    if (nums.length == 1) return nums[0];
    int prevPrev = nums[0];
    int prev = Math.max(nums[0], nums[1]);
    for (int i = 2; i < nums.length; i++) {
      int temp = prev;
      prev = Math.max(prev, nums[i] + prevPrev);
      prevPrev = temp;
    }
    return prev;
  }
}