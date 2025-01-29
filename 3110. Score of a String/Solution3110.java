class Solution3110 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "hello";
    System.out.println(sol.scoreOfString(s));
  }
}

class Solution {
  public int scoreOfString(String s) {
      int res = 0;
      for (int i = 0; i < s.length() - 1; i++) {
          res += Math.abs(s.charAt(i) - s.charAt(i + 1));
      }
      return res;
  }
}