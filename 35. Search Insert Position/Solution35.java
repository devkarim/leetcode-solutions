class Solution35 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] test = {1,3,5,6};
    int target = 2;
    int pos = sol.searchInsert(test, target);
    System.out.println(pos);
  }
}

class Solution {
  public int searchInsert(int[] nums, int target) {
      int l = 0;
      int r = nums.length - 1;
      while (r >= l) {
          int mid = (l + r) / 2;
          if (target > nums[mid]) {
              l = mid + 1;
          } else if (target < nums[mid]) {
              r = mid - 1;
          } else {
              return mid;
          }
      }
      return l;
  }
}