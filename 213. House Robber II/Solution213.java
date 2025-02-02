class Solution213 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {1, 2, 3};
    System.out.println(sol.rob(nums));    
  }
}

class Solution {
  public int rob(int[] nums) {
    // rules:
    // given: 0...n-1
    // first(0) and last(n-1) are neighbors
    // rob either 0..n-2 or 1..n-1
    if (nums.length == 1) return nums[0];
    if (nums.length == 2) return Math.max(nums[0], nums[1]);
    int n = nums.length;
    return Math.max(dp(nums, 0, n - 1), dp(nums, 1, n));
  }

  private int dp(int[] nums, int start, int n) {
    int prevPrev = nums[start];
    int prev = Math.max(nums[start], nums[start + 1]);
    for (int i = start + 2; i < n; i++) {
      int temp = prev;
      prev = Math.max(prev, nums[i] + prevPrev);
      prevPrev = temp;
    }
    return prev;
  }
  
  // private int dp(int[] nums, int idx, int n) {
  //   if (idx >= n) return 0;
  //   if (cache[idx] != -1) return cache[idx];
  //   cache[idx] = Math.max(dp(nums, idx + 1, n), nums[idx] + dp(nums, idx + 2, n));
  //   return cache[idx];
  // }
}