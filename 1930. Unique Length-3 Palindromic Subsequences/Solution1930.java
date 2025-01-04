import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

class Solution1930 {
  public static void main(String[] args) {
    String[] testCases = { "aabca", "adc", "bbcbaba" };
    int[] answer = { 3, 0, 4 };

    Solution sol = new Solution();

    for (int i = 0; i < testCases.length; i++) {
      String t = testCases[i];
      int n = sol.countPalindromicSubsequence(t);
      if (answer[i] == n) {
        System.out.println("Accepted at Test Case " + (i + 1));
      } else {
        System.out.println("Wrong Answer at Test Case " + (i + 1));
      }
    }
  }
}

class Solution {
  public int countPalindromicSubsequence(String s) {
    Map<Character, Integer> rightMap = new HashMap<Character, Integer>();
    for (char c : s.toCharArray()) {
      rightMap.put(c, rightMap.getOrDefault(c, 0) + 1);
    }
    Set<String> res = new HashSet<String>();
    Set<Character> leftSet = new HashSet<Character>();
    for (char c : s.toCharArray()) {
      rightMap.put(c, rightMap.get(c) - 1);
      for (char ch : leftSet) {
        if (rightMap.get(ch) != 0) {
          res.add("(" + c + "," + ch + ")");
        }
      }
      leftSet.add(c);
    }
    return res.size();
  }
}
