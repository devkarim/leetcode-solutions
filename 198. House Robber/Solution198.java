import java.util.Arrays;

class Solution198 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // int[] nums = {1, 2, 3, 1};
    int[] nums = {100, 1, 2, 104};
    System.out.println(sol.rob(nums));
  }
}

class Solution {
  private int[] cache;

  public int rob(int[] nums) {
    cache = new int[nums.length];
    Arrays.fill(cache, -1);
    return dp(nums, 0);
  }

  private int dp(int[] nums, int idx) {
    if (idx >= nums.length) return 0;
    if (cache[idx] != -1) return cache[idx];
    cache[idx] = Math.max(dp(nums, idx + 1), nums[idx] + dp(nums, idx + 2));
    return cache[idx];
  }
}