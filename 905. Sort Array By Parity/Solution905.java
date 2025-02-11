import java.util.Arrays;

class Solution905 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums = {3, 1, 2, 4};
    System.out.println(Arrays.toString(sol.sortArrayByParity(nums)));
  }
}

class Solution {
  public int[] sortArrayByParity(int[] nums) {
    int[] res = new int[nums.length];
    int l = 0;
    int r = nums.length - 1;
    for (int n : nums) {
      if (n % 2 == 0) {
        // if even
        res[l++] = n;
      } else {
        // if odd
        res[r--] = n;
      }
    }
    return res;
  }
}
