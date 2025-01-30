import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class Solution473 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] matchsticks = new int[]{1, 1, 2, 2, 2};
    System.out.println(sol.makesquare(matchsticks));
  }
}

class Solution {
  private final HashMap<Character, Integer> sides = new HashMap<>(Map.of(
    'T', 0,
    'B', 0,
    'L', 0,
    'R', 0
  ));

  public boolean makesquare(int[] matchsticks) {
    int totalSum = 0;
    for (int ms : matchsticks) {
      totalSum += ms;
    }
    if (totalSum % 4 != 0) return false;
    int target = totalSum / 4;
    Arrays.sort(matchsticks);
    reverse(matchsticks);
    return backtrack(matchsticks, 0, target);
  }

  private boolean backtrack(int[] matchsticks, int idx, int target) {
    // if out of boundaries, return
    if (idx == matchsticks.length)
      return true;
    // loop through all possible sides
    for (char side : sides.keySet()) {
      // check if current matchstick fits in the current side
      // by checking whether it is less than or equal the
      // target
      if (sides.get(side) + matchsticks[idx] <= target) {
        sides.put(side, sides.get(side) + matchsticks[idx]);
        boolean s = backtrack(matchsticks, idx + 1, target);
        sides.put(side, sides.get(side) - matchsticks[idx]);
        if (s) return true;
      }
    }
    return false;
  }

  private void reverse(int[] arr) {
    int l = 0;
    int r = arr.length - 1;
    while (l < r) {
      int temp = arr[l];
      arr[l++] = arr[r];
      arr[r--] = temp;
    }
  }
}