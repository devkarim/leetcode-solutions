class Solution1137 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 25;
    System.out.println(sol.tribonacci(n));    
  }
}

class Solution {
  public int tribonacci(int n) {
    int[] dp = new int[3];
    if (n == 0) return 0;
    if (n <= 2) return 1;
    dp[0] = 0;
    dp[1] = 1;
    dp[2] = 1;
    for (int i = 3; i <= n; i++) {
      int temp2 = dp[2];
      dp[2] = dp[0] + dp[1] + dp[2];
      dp[0] = dp[1];
      dp[1] = temp2;
    }
    return dp[2];
  }
}