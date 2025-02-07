class Solution152 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {2, 3, -2, 4};
    System.out.println(sol.maxProduct(nums));
  }
}

class Solution {
  public int maxProduct(int[] nums) {
    if (nums.length == 1) return nums[0];
    int min = 1;
    int max = 1;
    int res = max(nums);
    for (int n : nums) {
      if (n == 0) {
        min = max = 1;
        continue;
      }
      int temp = max;
      max = Math.max(Math.max(n * max, n * min), n);
      min = Math.min (Math.min(n * temp, n * min), n);
      res = Math.max(Math.max(min, max), res);
    }
    return res;
  }

  private int max(int[] nums) {
    int res = Integer.MIN_VALUE;
    for (int n : nums) {
      res = Math.max(res, n);
    }
    return res;
  }
}

