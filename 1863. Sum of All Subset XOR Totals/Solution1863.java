class Solution1863 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = { 5, 1, 6 };
    System.out.println(sol.subsetXORSum(nums));
  }
}

class Solution {
  public int subsetXORSum(int[] nums) {
    return subsetXORSum(nums, 0, 0);
  }

  private int subsetXORSum(int[] nums, int idx, int xorSum) {
    if (idx >= nums.length)
      return xorSum;
    // Decide to include the current element
    int withCurrElement = subsetXORSum(nums, idx + 1, xorSum ^ nums[idx]);
    // Decide to exclude the current element
    int withoutCurrElement = subsetXORSum(nums, idx + 1, xorSum);
    return withCurrElement + withoutCurrElement;
  }
}
