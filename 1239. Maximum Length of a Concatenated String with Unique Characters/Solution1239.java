import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution1239 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // List<String> arr = new ArrayList<>(Arrays.asList("un", "iq", "ue"));
    // List<String> arr = new ArrayList<>(Arrays.asList("aa", "bb"));
    List<String> arr = new ArrayList<>(Arrays.asList("cha", "r", "act", "ers"));
    System.out.println(sol.maxLength(arr));
  }
}

class Solution {
  private boolean[] used;
  
  public int maxLength(List<String> arr) {
    used = new boolean[26];
    return backtrack(arr, 0);
  }

  public int backtrack(List<String> arr, int idx) {
    if (idx >= arr.size()) return getCurrentStringSize();
    String s = arr.get(idx);
    int withCurr = 0;
    if (canMark(s)) {
      // include current string if only
      // the string is not used before
      markString(s, true);
      withCurr = backtrack(arr, idx + 1);
      // exclude current string
      markString(s, false);
    }
    int withoutCurr = backtrack(arr, idx + 1);
    return Math.max(withCurr, withoutCurr);
  }

  private boolean canMark(String s) {
    boolean[] duplicates = new boolean[26];
    for (char c : s.toCharArray()) {
      if (used[c - 'a'] || duplicates[c - 'a']) return false;
      duplicates[c - 'a'] = true;
    }
    return true;
  }

  private void markString(String s, boolean mark) {
    for (char c : s.toCharArray()) {
      used[c - 'a'] = mark;
    }
  }

  private int getCurrentStringSize() {
    int sum = 0;
    for (boolean u : used) {
      sum += u ? 1 : 0;
    }
    return sum;
  }
}