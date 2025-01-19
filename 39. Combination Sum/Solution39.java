import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

class Solution39 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] candidates = new int[]{2, 3, 6, 7};
    int target = 7;
    List<List<Integer>> res = sol.combinationSum(candidates, target);
    System.out.println(res);
  }
}

class Solution {
  public List<List<Integer>> combinationSum(int[] candidates, int target) {
    List<List<Integer>> res = new ArrayList<>();
    Stack<Integer> subset = new Stack<>();
    backtrack(res, candidates, target, 0, 0, subset);
    return res;
  }

  private void backtrack(List<List<Integer>> res, int[] candidates, int target, int sum, int idx, Stack<Integer> subset) {
    if (idx >= candidates.length) return;
    if (sum > target) return;
    if (sum == target) {
      res.add(new ArrayList<>(subset));
      return;
    }
    // include same element
    subset.add(candidates[idx]);
    backtrack(res, candidates, target, sum + candidates[idx], idx, subset);
    subset.pop();
    backtrack(res, candidates, target, sum, idx + 1, subset);
  }
}