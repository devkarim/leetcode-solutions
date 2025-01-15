class Solution2429 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    int num1 = 3;
    int num2 = 5;
    int res = sol.minimizeXor(num1, num2);
    System.out.println(res);
  }
}

class Solution {
  public int minimizeXor(int num1, int num2) {
    int x = num1;

    int count1 = countOnes(num1);
    int count2 = countOnes(num2);

    // number of ones in num1 < number of ones in num2
    while (count1 < count2) {
      x = x | (x + 1); // set rightmost unset bit
      count2--;
    }

    // number of ones in num1 > number of ones in num2
    while (count1 > count2) {
      x = x & (x - 1); // unset rightmost bit
      count1--;
    }

    return x;

    // 4 -> 0010 -> num1
    // 5 -> 0101 -> num2
    // x has same 1s as num2 AND x XOR num1 is minimal

    // 1001 XOR
    // 0010
    // =1011 -> 13

    // 0101 XOR
    // 0010
    // =0111 -> 7

    // 0110 XOR
    // 0010
    // =0100 -> 4

    // 0011 XOR
    // 0010
    // =0001 -> 1

    // 0011 | 0100 = 0111 = 7
    // 0100 & 0011 = 0000 = 0
  }

  private int countOnes(int num) {
    int res = 0;
    while (num > 0) {
      res += 1 & num;
      num = num >> 1;
    }
    return res;
  }
}
