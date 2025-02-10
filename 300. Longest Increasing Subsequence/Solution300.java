import java.util.HashMap;
import java.util.Map;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

class Solution300 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {4, 10, 4, 3, 8, 9};
    System.out.println(sol.lengthOfLIS(nums));
  }
}

class Solution {
  private int[][] cache;

  public int lengthOfLIS(int[] nums) {
    int n = nums.length;
    cache = new int[n][n + 1];
    for (int[] row : cache) {
      Arrays.fill(row, -1);
    }
    return dfs(nums, 0, -1);
  }

  private int dfs(int[] nums, int start, int last) {
    if (start >= nums.length) {
      return 0;
    }
    if (cache[start][last + 1] != -1) return cache[start][last + 1];
    // exclude current item
    int res = dfs(nums, start + 1, last);
    // include current element if it is larger than last item
    if (last == -1 || nums[last] < nums[start]) {
      res = Math.max(res, 1 + dfs(nums, start + 1, start));
    }
    cache[start][last + 1] = res;
    return res;
  }
}
