
import java.util.Arrays;

class Solution698 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {4,3,2,3,5,2,1};
    int k = 4;
    System.out.println(sol.canPartitionKSubsets(nums, k));
  }
}

class Solution {
  private int[] used;

  public boolean canPartitionKSubsets(int[] nums, int k) {
    int sum = 0;
    for (int n : nums) {
      sum += n;
    }
    if (sum % k != 0) return false;
    int target = sum / k;
    used = new int[nums.length];
    sortDesc(nums);
    return backtrack(nums, k, target, 0, 0);
  }

  private boolean backtrack(int[] nums, int k, int target, int idx, int subsetSum) {
    if (k == 0) return true;
    if (target == subsetSum) return backtrack(nums, k - 1, target, 0, 0);
    if (target < 0 || idx == nums.length) return false;
    for (int i = idx; i < nums.length; i++) {
      if (used[i] == 1 || subsetSum + nums[i] > target) continue;
      used[i] = 1;
      boolean foundSubset = backtrack(nums, k, target, i + 1, subsetSum + nums[i]);
      if (foundSubset) return true;
      used[i] = 0;
    }
    return false;
  }

  private void sortDesc(int[] nums) {
    Arrays.sort(nums);
    int l = 0;
    int r = nums.length - 1;
    while (l < r) {
      int temp = nums[l];
      nums[l] = nums[r];
      nums[r] = temp;
      l++;
      r--;
    }
  }
}