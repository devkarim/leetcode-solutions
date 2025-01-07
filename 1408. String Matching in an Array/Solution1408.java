
import java.util.ArrayList;
import java.util.List;

class Solution1408 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String[] test = {"mass", "as", "hero", "superhero"};
    List<String> res = sol.stringMatching(test);
    System.out.println(res);
  }
}

class Solution {
  public List<String> stringMatching(String[] words) {
      int n = words.length;
      List<String> res = new ArrayList<>();
      for (int i = 0; i < n; i++) {
          String w1 = words[i];
          for (int j = 0; j < n; j++) {
              String w2 = words[j];
              if (i != j && w2.contains(w1)) {
                  res.add(w1);
                  break;
              }
          }
      }
      return res;
  }
}