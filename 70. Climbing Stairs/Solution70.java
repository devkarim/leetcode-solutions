class Solution70 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 45;
    int res = sol.climbStairs(n);
    System.out.println(res);
  }
}

class Solution {
  public int climbStairs(int n) {
    int[] memory = new int[n + 1];
    return fib(n + 1, memory);
  }

  private int fib(int n, int[] memory) {
    if (n <= 1)
      return n;
    if (memory[n - 1] == 0) {
      memory[n - 1] = fib(n - 1, memory) + fib(n - 2, memory);
    }
    return memory[n - 1];
  }
}