class Solution441 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 5;
    System.out.println(sol.arrangeCoins(n));
  }
}

class Solution {
  public int arrangeCoins(int n) {
    int l = 1;
    int r = n;
    while (l <= r) {
      int mid = l + (r - l) / 2;
      long coins = (long) mid * (mid + 1) / 2;
      if (coins == n) return mid;
      if (coins > n) {
        r = mid - 1;
      } else {
        l = mid + 1;
      }
    }
    return r;
  }
}
