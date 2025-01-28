import java.util.ArrayList;
import java.util.List;
import java.util.Map;

class Solution17 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String digits = "234";
    System.out.println(sol.letterCombinations(digits));
  }
}

class Solution {
  private List<String> res;
  private List<Character> subset;
  private static final Map<Character, String> digitToLetters = Map.of(
    '2', "abc",
    '3', "def",
    '4', "ghi",
    '5', "jkl",
    '6', "mno",
    '7', "pqrs",
    '8', "tuv",
    '9', "wxyz"
  );

  public List<String> letterCombinations(String digits) {
    res = new ArrayList<>();
    if (digits.length() == 0) return res;
    subset = new ArrayList<>();
    backtrack(digits, 0, 0);
    return res;
  }

  private void backtrack(String digits, int digitIdx, int letterIdx) {
    // if no more digits to process, add the current subset to result
    if (digitIdx == digits.length()) {
      res.add(listToString(subset));
      return;
    }
    // loop through letters of the digit
    String letters = digitToLetters.get(digits.charAt(digitIdx));
    for (int j = letterIdx; j < letters.length(); j++) {
      subset.add(letters.charAt(j));
      backtrack(digits, digitIdx + 1, letterIdx);
      subset.removeLast();
    }
  }

  private String listToString(List<Character> chars) {
    StringBuilder str = new StringBuilder();
    for (char c : chars) {
      str.append(c);
    }
    return str.toString();
  }
}