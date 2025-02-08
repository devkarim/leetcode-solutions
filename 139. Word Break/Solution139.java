import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution139 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "cars";
    List<String> wordDict = Arrays.asList("car", "rs", "ca");
    System.out.println(sol.wordBreak(s, wordDict));
  }
}

class Solution {
  private Map<Integer, Boolean> cache;

  public boolean wordBreak(String s, List<String> wordDict) {
    // int start = 0;
    // while (start < s.length()) {
    //   int temp = start;
    //   for (String choice : wordDict) {
    //     int choiceLen = choice.length();
    //     int end = start + choiceLen;
    //     if (end <= s.length() && s.substring(start, end).equals(choice)) {
    //       start = end;
    //       break;
    //     }      
    //   }
    //   if (temp == start) return false;
    // }
    cache = new HashMap<>();
    return dfs(s, wordDict, 0);
  }

  private boolean dfs(String s, List<String> wordDict, int start) {
    if (start >= s.length()) return true;
    if (cache.containsKey(start)) return cache.get(start);
    for (String c : wordDict) {
      int cLen = c.length();
      int end = start + cLen;
      if (end <= s.length() && s.substring(start, end).equals(c)) {
        boolean success = dfs(s, wordDict, end);
        cache.put(start, success);
        if (success) return true;
      }
    }
    cache.put(start, false);
    return false;
  }
}
