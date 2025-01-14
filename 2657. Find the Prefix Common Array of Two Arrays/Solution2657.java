import java.util.Arrays;

class Solution2657 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int[] A = { 1, 3, 2, 4 };
    int[] B = { 3, 1, 2, 4 };
    int[] C = sol.findThePrefixCommonArray(A, B);
    System.out.println(Arrays.toString(C));
  }
}

class Solution {
  public int[] findThePrefixCommonArray(int[] A, int[] B) {
    int n = A.length;
    int[] res = new int[n];
    int[] freq = new int[n + 1];
    int count = 0;
    for (int i = 0; i < n; i++) {
      int a = A[i];
      int b = B[i];
      if (++freq[a] == 2)
        count++;
      if (++freq[b] == 2)
        count++;
      res[i] = count;
    }
    return res;
  }
}
