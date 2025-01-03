class Solution2270 {
  public static void main(String[] args) {
    int[][] testCases = { { 10, 4, -8, 7 }, { 2, 3, 1, 0 }, { 10, 4 } };
    int[] answer = { 2, 2, 1 };

    Solution sol = new Solution();

    for (int i = 0; i < testCases.length; i++) {
      int[] t = testCases[i];
      int n = sol.waysToSplitArray(t);
      if (answer[i] == n) {
        System.out.println("Accepted at Test Case " + (i + 1));
      } else {
        System.out.println("Wrong Answer at Test Case " + (i + 1));
      }
    }
  }
}

class Solution {
  public int waysToSplitArray(int[] nums) {
    int n = nums.length;
    long[] prefixSum = new long[nums.length + 1];
    prefixSum[0] = 0;
    for (int i = 0; i < n; i++) {
      prefixSum[i + 1] = prefixSum[i] + nums[i];
    }
    int numberOfWays = 0;
    for (int i = 0; i < n - 1; i++) {
      if (prefixSum[i + 1] >= (prefixSum[n] - prefixSum[i + 1])) {
        numberOfWays++;
      }
    }
    return numberOfWays;
  }
}