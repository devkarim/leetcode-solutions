import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution90 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = { 1, 2, 2 };
    System.out.println(sol.subsetsWithDup(nums));
  }
}

class Solution {
  public List<List<Integer>> subsetsWithDup(int[] nums) {
    Arrays.sort(nums);
    List<List<Integer>> res = new ArrayList<>();
    backtrack(res, nums, 0, new ArrayList<>());
    return res;
  }

  private void backtrack(List<List<Integer>> res, int[] nums, int idx, List<Integer> subset) {
    if (idx >= nums.length) {
      res.add(new ArrayList<>(subset));
      return;
    }
    // include current element
    subset.add(nums[idx]);
    backtrack(res, nums, idx + 1, subset);
    // remove current element
    subset.remove(subset.size() - 1);
    // skip duplicates of current num
    while (idx + 1 < nums.length && nums[idx + 1] == nums[idx]) {
      idx++;
    }
    backtrack(res, nums, idx + 1, subset);
  }
}