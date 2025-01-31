class Solution746 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // int[] cost = new int[]{1,100,1,1,1,100,1,1,100,1};
    int[] cost = new int[]{10, 15, 20};
    System.out.println(sol.minCostClimbingStairs(cost));
  }
}

class Solution {
  public int minCostClimbingStairs(int[] cost) {
    int[] cache = new int[cost.length];
    initCache(cache);
    return Math.min(dp(cost, 0, cache), dp(cost, 1, cache));
  }

  private int dp(int[] cost, int n, int[] cache) {
    if (n >= cost.length) return 0; 
    if (cache[n] != -1) return cache[n];
    cache[n] = cost[n] + Math.min(dp(cost, n + 1, cache), dp(cost, n + 2, cache));
    return cache[n];
  }

  private void initCache(int[] cache) {
    for (int i = 0; i < cache.length; i++) {
      cache[i] = -1;
    }
  }
}