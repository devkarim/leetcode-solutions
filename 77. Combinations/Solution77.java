import java.util.List;
import java.util.ArrayList;

class Solution77 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 4;
    int k = 2;
    System.out.println(sol.combine(n, k));
  }
}

class Solution {
  public List<List<Integer>> combine(int n, int k) {
    return backtrack(n, k, 1, new ArrayList<>());
  }

  private List<List<Integer>> backtrack(int n, int k, int start, List<Integer> subset) {
    List<List<Integer>> res = new ArrayList<>();
    if (subset.size() == k) {
      res.add(new ArrayList<>(subset));
    }
    for (int i = start; i <= n; i++) {
      subset.add(i);
      res.addAll(backtrack(n, k, i + 1, subset));
      subset.remove(subset.size() - 1);
    }
    return res;
  }
}
