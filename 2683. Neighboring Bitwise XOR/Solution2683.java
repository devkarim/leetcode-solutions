class Solution2683 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] derived = new int[] { 1, 1, 0 };
    System.out.println(sol.doesValidArrayExist(derived));
  }
}

class Solution {
  public boolean doesValidArrayExist(int[] derived) {
    int sumXor = 0;
    for (int n : derived) {
      sumXor ^= n;
    }
    return sumXor == 0;
    // Given: [1,1,0]
    // Original: [0,1,0]
    // To get 1 -> [0,1] OR [1,0]
    // To get 0 -> [0,0] OR [1,1]

    // Original: [x, y, z]
    // Given: [x^y, y^z, z^x]
    // x ^ y ^ y ^ z ^ z ^ x = 0

    // Given: [1,0]
    // 1 XOR 0 always gives 1 so 0 is not appropriate
  }
}
