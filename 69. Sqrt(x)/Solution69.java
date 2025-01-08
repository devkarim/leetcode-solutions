class Solution69 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int res = sol.mySqrt(8);
    System.out.println(res);
  }
}

class Solution {
  public int mySqrt(int x) {
    if (x == 1 || x == 0)
      return x;
    int l = 1;
    int r = x / 2;
    while (l <= r) {
      int mid = l + (r - l) / 2;
      long res = (long) mid * mid;
      if (res == x)
        return mid;
      if (res < x) {
        l = mid + 1;
      } else {
        r = mid - 1;
      }
    }
    return r;
  }
}