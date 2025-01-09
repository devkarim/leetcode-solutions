class Solution2185 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String[] test = { "pay", "attention", "practice", "attend" };
    String pref = "at";
    int res = sol.prefixCount(test, pref);
    System.out.println(res);
  }
}

class Solution {
  public int prefixCount(String[] words, String pref) {
    int res = 0;
    for (String w : words) {
      if (w.startsWith(pref)) {
        res++;
      }
    }
    return res;
  }
}
