
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
    int[] dp = new int[nums.length];
    dp[0] = nums[0];
    dp[1] = Math.max(nums[0], nums[1]);
    for (int i = 2; i < nums.length; i++) {
      dp[i] = Math.max(dp[i - 1], nums[i] + dp[i - 2]);
    }
    return dp[nums.length - 1];
  }
}