public class Solution540 {
  public static void main(String[] args) {
    int[][] testCases = { { 1, 1, 2, 3, 3, 4, 4, 8, 8 }, { 3, 3, 7, 7, 10, 11, 11 } };
    int[] answer = { 2, 10 };

    Solution sol = new Solution();

    for (int i = 0; i < testCases.length; i++) {
      int[] t = testCases[i];
      int n = sol.singleNonDuplicate(t);
      if (answer[i] == n) {
        System.out.println("Accepted at Test Case " + (i + 1));
      } else {
        System.out.println("Wrong Answer at Test Case " + (i + 1));
      }
    }
  }
}

class Solution {
  public int singleNonDuplicate(int[] nums) {
    if (nums.length == 1)
      return nums[0];
    int low = 0;
    int high = nums.length - 1;
    while (high > low) {
      int mid = (low + high) / 2;
      if ((mid - 1 >= 0 && mid % 2 == 0 && nums[mid - 1] == nums[mid]) ||
          (mid + 1 < nums.length && mid % 2 == 1 && nums[mid + 1] == nums[mid])) {
        high = mid - 1;
      } else if ((mid + 1 < nums.length && mid % 2 == 0 && nums[mid + 1] == nums[mid]) ||
          (mid - 1 >= 0 && mid % 2 == 1 && nums[mid - 1] == nums[mid])) {
        low = mid + 1;
      } else {
        return nums[mid];
      }
    }
    return nums[low];
  }
}