import java.util.Arrays;

class Solution91 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "106";
    System.out.println(sol.numDecodings(s));   
  }
}

class Solution {
  private int[] cache;
  
  public int numDecodings(String s) {
    if (s.length() == 1) return s.charAt(0) == '0' ? 0 : 1;
    cache = new int[s.length()];
    Arrays.fill(cache, -1);
    return dp(s, 0);
  }

  public int dp(String s, int idx) {
    if (idx >= s.length()) return 1;
    if (s.charAt(idx) == '0') return 0;
    if (cache[idx] != -1) return cache[idx];
    int res = dp(s, idx + 1);
    if (s.length() > idx + 1) {
      if (s.charAt(idx) == '1' || (s.charAt(idx) == '2' && s.charAt(idx + 1) < '7')) {
        res += dp(s, idx + 2);
      }
    }
    cache[idx] = res;
    return res;
  }
}
