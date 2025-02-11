import java.util.Arrays;

class Solution300 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    //int[] nums = {4, 10, 4, 3, 8, 9};
    int[] nums = {0, 1, 0, 3, 2, 3};
    System.out.println(sol.lengthOfLIS(nums));
  }
}

class Solution {
  public int lengthOfLIS(int[] nums) {
    int[] dp = new int[nums.length];
    Arrays.fill(dp, 1);
    int res = 1;
    for (int start = nums.length - 1; start >= 0; start--) {
      for (int end = start + 1; end < nums.length; end++) {
        if (nums[end] <= nums[start]) continue;
        dp[start] = Math.max(dp[start], 1 + dp[end]);
        res = Math.max(res, dp[start]);
      }
    }
    return res;
  }
}
