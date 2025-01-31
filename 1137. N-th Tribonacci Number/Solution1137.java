class Solution1137 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 25;
    System.out.println(sol.tribonacci(n));    
  }
}

class Solution {
  public int tribonacci(int n) {
    int[] cache = new int[n + 1];
    return dp(n, cache);
  }

  private int dp(int n, int[] cache) {
    if (n <= 1) return n;
    if (n == 2) return 1;
    if (cache[n] != 0) return cache[n];
    cache[n] = dp(n - 1, cache) + dp(n - 2, cache) + dp(n - 3, cache);
    return cache[n];
  }
}