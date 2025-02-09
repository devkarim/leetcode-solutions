class Solution441 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 5;
    System.out.println(sol.arrangeCoins(n));
  }
}

class Solution {
  public int arrangeCoins(int n) {
    int max = n;
    for (int i = 1; i <= max; i++) {
      n -= i;
      if (n == 0) return i;
      if (n < 0) return i - 1;
    }
    return -1;
  }
}
