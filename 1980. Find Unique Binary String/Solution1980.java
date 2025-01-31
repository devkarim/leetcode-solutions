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
  private StringBuilder str;

  public String findDifferentBinaryString(String[] nums) {
    str = new StringBuilder();
    numSet = new HashSet<>();
    numSet.addAll(Arrays.asList(nums));
    backtrack(nums.length);
    return str.toString();
  }

  private boolean backtrack(int n) {
    if (n == 0) return !numSet.contains(str.toString());
    str.append('0');
    boolean found = backtrack(n - 1);
    if (found) return true;
    str.deleteCharAt(str.length() - 1);
    str.append('1');
    found = backtrack(n - 1);
    if (found) return true;
    str.deleteCharAt(str.length() - 1);
    return false;
  }
}
