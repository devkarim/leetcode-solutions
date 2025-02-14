import java.util.HashSet;
import java.util.Set;

class Solution416 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {1, 5, 11, 5};
    //int[] nums = {1, 2, 3, 5};
    System.out.println(sol.canPartition(nums));
  }
}

class Solution {
  public boolean canPartition(int[] nums) {
    int sum = 0;
    for (int num : nums) {
      sum += num;
    }
    if (sum % 2 != 0) return false;
    int target = sum / 2;
    Set<Integer> dp = new HashSet<>();
    dp.add(0);
    for (int n : nums) {
      Set<Integer> newDP = new HashSet<>();
      for (int t : dp) {
        newDP.add(n + t);
        newDP.add(t);
      }
      dp = newDP;
    }
    return dp.contains(target);
  }
}

// class Solution {
//   private Boolean[][] cache;
// 
//   public boolean canPartition(int[] nums) {
//     int sum = 0;
//     for (int n : nums) {
//       sum += n;
//     }
//     if (sum % 2 != 0) return false;
//     int target = sum / 2;
//     cache = new Boolean[nums.length][target + 1];
//     return dfs(nums, 0, target);
//   }
// 
//   private boolean dfs(int[] nums, int idx, int target) {
//     if (idx >= nums.length) {
//       return target == 0;
//     }
//     if (target < 0) return false;
//     if (cache[idx][target] != null) return cache[idx][target];
//     cache[idx][target] = dfs(nums, idx + 1, target) || dfs(nums, idx + 1, target - nums[idx]);
//     return cache[idx][target];
//   }
// }

