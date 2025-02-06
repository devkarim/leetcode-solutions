class Solution374 {
  public static void main(String[] args) {
    Solution sol = new Solution(6);
    int n = 10;
    System.out.println(sol.guessNumber(n));
  }
}

class Solution {
  private int pick;

  Solution(int pick) {
    this.pick = pick;
  }
  
  public int guess(int num) {
    if (num > pick) return -1;
    if (num < pick) return 1;
    return 0;
  }
  
  public int guessNumber(int n) {
    int l = 0;
    int r = n;
    while (l <= r) {
      int mid = l + (r - l) / 2;
      int g = guess(mid);
      if (g == -1) {
        r = mid - 1;
      } else if (g == 1) {
        l = mid + 1;
      } else {
        return mid;
      }
    }
    return -1;
  }
}
