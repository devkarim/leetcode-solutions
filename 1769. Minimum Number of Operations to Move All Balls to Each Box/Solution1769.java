
import java.util.Arrays;

class Solution1769 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] ans = sol.minOperations("001011");
    System.out.println(Arrays.toString(ans));
  }
}

class Solution {
  public int[] minOperations(String boxes) {
      int n = boxes.length();
      int[] res = new int[n];
      int balls = 0;
      int moves = 0;
      for (int i = 0; i < n; i++) {
          res[i] = balls + moves;
          moves += balls;
          if (boxes.charAt(i) == '1') {
              balls++;
          }
      }
      balls = 0;
      moves = 0;
      for (int i = n - 1; i >= 0; i--) {
          res[i] += balls + moves;
          moves += balls;
          if (boxes.charAt(i) == '1') {
              balls++;
          }
      }
      return res;
  }
}