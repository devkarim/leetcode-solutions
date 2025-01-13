class Solution3223 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "abaacbcbb";
    System.out.println(sol.minimumLength(s));
  }
}

class Solution {
  public int minimumLength(String s) {
    int n = s.length();
    int res = n;
    int[] freq = new int[26];
    for (int i = 0; i < n; i++) {
      freq[s.charAt(i) - 'a']++;
    }
    for (int occurrences : freq) {
      while (occurrences > 2) {
        res -= 2;
        occurrences -= 2;
      }
    }
    return res;
  }
}
