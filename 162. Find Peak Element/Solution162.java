class Solution162 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {1,2,3,1};
    System.out.println(sol.findPeakElement(nums));
  }
}

class Solution {
  public int findPeakElement(int[] nums) {
    int l = 0;
    int r = nums.length - 1;
    while (l < r) {
      int mid = l + (r - l) / 2;
      if (mid + 1 < nums.length && nums[mid] < nums[mid + 1]) {
        l = mid + 1;
      } else if (mid - 1 >= 0 && nums[mid] < nums[mid - 1]) {
        r = mid - 1;
      } else {
        return mid;
      }
    }
    return l;
  }
}