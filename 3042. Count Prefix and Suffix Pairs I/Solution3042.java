public class Solution3042 {
  public static void main(String[] args) {
    String[] test = { "a", "aba", "ababa", "aa" };

    Solution sol = new Solution();
    int res = sol.countPrefixSuffixPairs(test);
    System.out.println(res);
  }
}

class Solution {
  public int countPrefixSuffixPairs(String[] words) {
    int n = words.length;
    int res = 0;
    for (int i = 0; i < n; i++) {
      String str1 = words[i];
      for (int j = i + 1; j < n; j++) {
        String str2 = words[j];
        if (i < j && isPrefixAndSuffix(str1, str2)) {
          res++;
        }
      }
    }
    return res;
  }

  public boolean isPrefixAndSuffix(String str1, String str2) {
    return str2.startsWith(str1) && str2.endsWith(str1);
  }
}
