import java.util.HashMap;
import java.util.Map;

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
  private Map<Integer, Integer> cache;

  public int coinChange(int[] coins, int amount) {
    if (amount == 0) return 0;
    cache = new HashMap<>();
    int minCoins = dfs(coins, amount);
    return Integer.MAX_VALUE == minCoins ? -1 : minCoins;
  }

  public int dfs(int[] coins, int amount) {
    if (amount == 0) return 0;
    if (cache.containsKey(amount)) return cache.get(amount);
    int min = Integer.MAX_VALUE;
    for (int coin : coins) {
      if (amount - coin >= 0) {
        int minCoins = dfs(coins, amount - coin);
        if (minCoins != Integer.MAX_VALUE) {
          min = Math.min(min, minCoins + 1);
        }
      }
    }
    cache.put(amount, min);
    return min;
  }
}