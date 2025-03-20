import java.util.Arrays;

class Solution322 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // int[] coins = {1,2,5};
    // int amount = 11;
    int[] coins = {2,5,10,1};
    int amount = 27;
    System.out.println(sol.coinChange(coins, amount));
  }
}

class Solution {
  public int coinChange(int[] coins, int amount) {
    if (amount == 0) return 0;
    int[] dp = new int[amount + 1];
    Arrays.fill(dp, amount + 1);
    dp[0] = 0;
    for (int i = 1; i <= amount; i++) {
      for (int j = 0; j < coins.length; j++) {
        if (coins[j] <= i) {
          dp[i] = Math.min(dp[i], dp[i - coins[j]] + 1);
        }
      }
    }
    return dp[amount] > amount ? -1 : dp[amount];
  }
}
