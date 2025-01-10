import java.util.List;
import java.util.ArrayList;

class Solution916 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String[] words1 = { "amazon", "apple", "facebook", "google", "leetcode" };
    String[] words2 = { "e", "o" };
    List<String> res = sol.wordSubsets(words1, words2);
    System.out.println(res);
  }
}

class Solution {
  public List<String> wordSubsets(String[] words1, String[] words2) {
    List<String> res = new ArrayList<>();
    int[] max = new int[26];
    // get maximum number of character occurrences in a single word in all of the
    // array
    for (String w2 : words2) {
      int[] freq = getFrequencyArray(w2);
      for (int i = 0; i < 26; i++) {
        max[i] = Math.max(max[i], freq[i]);
      }
    }
    // check if each word in words2 is universal
    for (String w1 : words1) {
      if (isUniversal(w1, max)) {
        res.add(w1);
      }
    }
    return res;
  }

  private int[] getFrequencyArray(String word) {
    int[] freq = new int[26];
    for (char c : word.toCharArray())
      freq[c - 'a']++;
    return freq;
  }

  private boolean isUniversal(String word, int[] max) {
    int[] freq = getFrequencyArray(word);
    for (int i = 0; i < 26; i++) {
      if (freq[i] < max[i])
        return false;
    }
    return true;
  }
}