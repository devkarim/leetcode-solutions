import java.util.ArrayList;
import java.util.List;
import java.util.HashSet;
import java.util.Set;

class Solution93 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "25525511135";
    System.out.println(sol.restoreIpAddresses(s));
  }
}

class Solution {
  private List<String> res;
  private Set<Integer> subset;

  public List<String> restoreIpAddresses(String s) {
    res = new ArrayList<>();
    subset = new HashSet<>();
    backtrack(s, 0, 0);
    return res;
  }

  private void backtrack(String s, int start, int numberOfDots) {
    // check if we finished the number of dots
    if (numberOfDots == 4) {
      if (start >= s.length()) {
        res.add(dotsToString(s));
      }
      return;
    }
    for (int end = start; end < start + 3 && end < s.length(); end++) {
      String sub = s.substring(start, end + 1);
      if (isValidNumber(sub)) {
        subset.add(end);
        backtrack(s, end + 1, numberOfDots + 1);
        subset.remove(end);
      }
    }
  }

  private boolean isValidNumber(String s) {
    // check if string has any leading zeros
    if (s.startsWith("0") && s.length() != 1)
      return false;
    int n = Integer.parseInt(s);
    // check if out of boundaries
    if (n > 255 || n < 0)
      return false;
    return true;
  }

  private String dotsToString(String s) {
    StringBuilder str = new StringBuilder();
    for (int i = 0; i < s.length(); i++) {
      str.append(s.charAt(i));
      // append dot at specified location except the ending
      if (subset.contains(i) && i != s.length() - 1) {
        str.append(".");
      }
    }
    return str.toString();
  }
}
