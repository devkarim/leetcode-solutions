class Solution69 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int res = sol.mySqrt(8);
    System.out.println(res);
  }
}

class Solution {
  public int mySqrt(int x) {
    int l = 0;
    int r = x;
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