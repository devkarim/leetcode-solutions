import java.util.Arrays;

class Solution338 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int n = 2;
    System.out.println(Arrays.toString(sol.countBits(n)));
  }
}

class Solution {
  public int[] countBits(int n) {
    int[] res = new int[n + 1];
    for (int i = 0; i <= n; i++) {
      res[i] = count_ones(i);
    }
    return res;
  }

  public int count_ones(int num) {
    int res = 0;
    while (num > 0) {
      res += 1 & num;
      num = num >> 1;
    }
    return res;
  }
}
