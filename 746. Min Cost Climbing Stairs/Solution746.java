class Solution746 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] cost = new int[]{1,100,1,1,1,100,1,1,100,1};
    // int[] cost = new int[]{10, 15, 20};
    System.out.println(sol.minCostClimbingStairs(cost));
  }
}

class Solution {
  public int minCostClimbingStairs(int[] cost) {
    int n = cost.length;
    int[] dp = new int[2];

    for (int i = 2; i <= n; i++) {
      int tmp = dp[1];
      dp[1] = Math.min(dp[1] + cost[i - 1], dp[0] + cost[i - 2]);
      dp[0] = tmp;
    }

    return dp[1];
  }
}