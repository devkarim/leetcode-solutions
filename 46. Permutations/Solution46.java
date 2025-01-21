import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

class Solution46 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {1, 2, 3};
    System.out.println(sol.permute(nums));
  }
}

class Solution {
  public List<List<Integer>> permute(int[] nums) {
    return backtrack(convertToSet(nums), new ArrayList<>());
  }

  private List<List<Integer>> backtrack(HashSet<Integer> nums, List<Integer> subset) {
    List<List<Integer>> res = new ArrayList<>();
    if (nums.isEmpty()) {
      res.add(new ArrayList<>(subset));
    }
    HashSet<Integer> remainingNums = new HashSet<>(nums);
    for (int n : nums) {
      remainingNums.remove(n);
      subset.add(n);
      res.addAll(backtrack(remainingNums, subset));
      subset.removeLast();
      remainingNums.add(n);
    }
    return res;
  }

  private HashSet<Integer> convertToSet(int[] arr) {
    HashSet<Integer> res = new HashSet<>();
    for (int n : arr) {
      res.add(n);
    }
    return res;
  }
}