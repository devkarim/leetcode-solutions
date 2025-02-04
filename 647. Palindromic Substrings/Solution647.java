class Solution647 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "aaa";
    System.out.println(sol.countSubstrings(s));
  }
}

class Solution {
  public int countSubstrings(String s) {
    int cnt = 0;
    int n = s.length();
    for (int curr = 0; curr < n; curr++) {
      // if odd
      int l = curr;
      int r = curr;
      while (l >= 0 && r < n && s.charAt(l) == s.charAt(r)) {
        cnt++;
        r++;
        l--;
      }
      // if even
      l = curr;
      r = curr + 1;
      while (l >= 0 && r < n && s.charAt(l) == s.charAt(r)) {
        cnt++;
        r++;
        l--;
      }
    }
    return cnt;      
  }
}