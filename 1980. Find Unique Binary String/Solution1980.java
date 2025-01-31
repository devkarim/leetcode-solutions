import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

class Solution1980 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // String[] nums = new String[]{"01", "10"};
    String[] nums = new String[]{"111", "011", "001"};
    System.out.println(sol.findDifferentBinaryString(nums));
  }
}

class Solution {
  private Set<String> numSet;
  private final char[] chars = new char[]{'0', '1'};

  public String findDifferentBinaryString(String[] nums) {
    numSet = new HashSet<>();
    numSet.addAll(Arrays.asList(nums));
    return backtrack(nums.length, new StringBuilder());
  }

  private String backtrack(int n, StringBuilder str) {
    if (str.length() == n) {
      if (!numSet.contains(str.toString())) return str.toString();
      return ""; 
    }
    for (char c : chars) {
      str.append(c);
      String res = backtrack(n, str);
      if (res.length() != 0) return res;
      str.deleteCharAt(str.length() - 1);
    }
    return "";
  }
}
