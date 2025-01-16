class Solution2425 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] nums1 = { 2, 1, 3 };
    int[] nums2 = { 10, 2, 5, 0 };
    int res = sol.xorAllNums(nums1, nums2);
    System.out.println(res);
  }
}

class Solution {
  public int xorAllNums(int[] nums1, int[] nums2) {
    int res = 0;
    int m = nums1.length;
    int n = nums2.length;
    if (n % 2 == 1) {
      for (int num : nums1) {
        res ^= num;
      }
    }
    if (m % 2 == 1) {
      for (int num : nums2) {
        res ^= num;
      }
    }
    return res;
    // nums1 = [a, b, c]
    // nums2 = [x, y, z]
    // nums3 = [a^x, a^y, a^z, b^x, b^y, b^z, c^x, c^y, c^z]
    // res = a ^ x ^ a ^ y ^ a ^ z ^ b ^ x ^ b ^ y ^ b ^ z ^ c ^ x ^ c ^ y ^ c ^ z
    // res = a ^ b ^ c ^ x ^ y ^ z
    // 1^1 = 0
    // 1^0 = 1

    // [a, b]
    // [x, y]
    // [a^x,a^y,b^x,b^y]
    // res = a ^ x ^ a ^ y ^ b ^ x ^ b ^ y

    // [a,b]
    // [x,y,z]
    // [a^x,a^y,a^z,b^x,b^y,b^z]
    // res = a ^ x ^ a ^ y ^ a ^ z ^ b ^ x ^ b ^ y ^ b ^ z
  }
}
