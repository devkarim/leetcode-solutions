import java.util.Arrays;
import java.util.List;

class Solution139 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "cars";
    List<String> wordDict = Arrays.asList("car", "rs", "cas", "cars");
    System.out.println(sol.wordBreak(s, wordDict));
  }
}

class Solution {
  public boolean wordBreak(String s, List<String> wordDict) {
    int n = s.length();
    boolean[] dp = new boolean[n + 1];
    dp[n] = true;

    for (int start = n; start >= 0; start--) {
      for (String choice : wordDict) {
        int choiceLen = choice.length();
        int end = start + choiceLen;
        if (end <= s.length() && s.substring(start, end).equals(choice)) {
          dp[start] = dp[end];
        }
        if (dp[start]) break;
      }
    }

    return dp[0];
  }
}
