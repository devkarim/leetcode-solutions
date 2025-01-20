import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Stack;

class Solution40 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    // int[] candidates = new int[]{10, 1, 2, 7, 6, 1, 5};
    int[] candidates = new int[]{2, 5, 2, 1, 2};
    int target = 5;
    System.out.println(sol.combinationSum2(candidates, target));
  }
}

class Solution {
  public List<List<Integer>> combinationSum2(int[] candidates, int target) {
    Arrays.sort(candidates);
    List<List<Integer>> res = new ArrayList<>();
    Stack<Integer> subset = new Stack<>();
    backtrack(res, candidates, target, 0, 0, subset);
    return res;
  }

  private void backtrack(List<List<Integer>> res, int[] candidates, int target, int sum, int idx, Stack<Integer> subset) {
    if (target == sum) {
      res.add(new ArrayList<>(subset));
      return;
    }
    if (idx >= candidates.length || target < sum) return;
    // include current element
    subset.push(candidates[idx]);
    backtrack(res, candidates, target, sum + candidates[idx], idx + 1, subset);
    // exclude current element and similar neighbours
    while (idx + 1 < candidates.length && candidates[idx] == candidates[idx + 1]) {
      idx++;
    }
    subset.pop();
    backtrack(res, candidates, target, sum, idx + 1, subset);
  }
}