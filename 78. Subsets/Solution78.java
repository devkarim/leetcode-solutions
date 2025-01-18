import java.util.List;
import java.util.Stack;
import java.util.ArrayList;

class Solution78 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = { 1, 2, 3 };
    System.out.println(sol.subsets(nums));
  }
}

class Solution {
  public List<List<Integer>> subsets(int[] nums) {
    List<List<Integer>> res = new ArrayList<>();
    Stack<Integer> subset = new Stack<Integer>();
    subsets(res, nums, 0, subset);
    return res;
  }

  private void subsets(List<List<Integer>> res, int[] nums, int idx, Stack<Integer> subset) {
    if (idx >= nums.length) {
      res.add(new ArrayList<>(subset));
      return;
    }
    // Include current element
    subset.push(nums[idx]);
    subsets(res, nums, idx + 1, subset);
    // Exclude current element
    subset.pop();
    subsets(res, nums, idx + 1, subset);
  }
}
