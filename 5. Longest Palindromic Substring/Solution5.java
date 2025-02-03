class Solution5 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // String s = "abb";
    // String s = "cbbd";
    // String s = "babad";
    String s = "cbbc";
    System.out.println(sol.longestPalindrome(s));
  }
}

class Solution {
  public String longestPalindrome(String s) {
    int start = 0;
    int resLen = 0;
    int n = s.length();
    for (int curr = 0; curr < n; curr++) {
      int l = curr;
      int r = curr;
      // if odd
      int[] resOdd = expand(s, l, r, resLen);
      if (resOdd[0] > resLen) {
        resLen = resOdd[0];
        start = resOdd[1];
      }
      // if even
      l = curr;
      r = curr + 1;
      int[] resEven = expand(s, l, r, resLen);
      if (resEven[0] > resLen) {
        resLen = resEven[0];
        start = resEven[1];
      }
    }
    return s.substring(start, start + resLen);
  }

  private int[] expand(String s, int l, int r, int max) {
    int[] res = new int[2];
    while (l >= 0 && r < s.length() && s.charAt(l) == s.charAt(r)) {
      if (r - l + 1 > max) {
        res[0] = r - l + 1; 
        res[1] = l;
      }
      l--;
      r++;
    }
    return res;
  }
}