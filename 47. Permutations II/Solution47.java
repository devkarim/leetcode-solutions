import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution47 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {1, 1, 2};
    System.out.println(sol.permuteUnique(nums));
  }
}

class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
      Map<Integer, Integer> countMap = createHashMapCount(nums); 
      return backtrack(nums, countMap, new ArrayList<>());
    }

    public List<List<Integer>> backtrack(int[] nums, Map<Integer, Integer> countMap, List<Integer> perm) {
      List<List<Integer>> res = new ArrayList<>();
      if (perm.size() == nums.length) {
        res.add(new ArrayList<>(perm));
        return res;
      }
      for (int n : countMap.keySet()) {
        if (countMap.get(n) > 0) {
          countMap.put(n, countMap.get(n) - 1);
          perm.add(n);

          res.addAll(backtrack(nums, countMap, perm));

          perm.removeLast();
          countMap.put(n, countMap.get(n) + 1);
        }
      }
      return res;
    }

    private Map<Integer, Integer> createHashMapCount(int[] nums) {
      Map<Integer, Integer> countMap = new HashMap<>();
      for (int n : nums) {
        countMap.put(n, countMap.getOrDefault(n, 0) + 1);
      }
      return countMap;
    }
}