import java.util.ArrayList;
import java.util.List;

class Solution131 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "aab";
    System.out.println(sol.partition(s));
  }
}

class Solution {
    private List<List<String>> res;
    private List<String> subset;

    public List<List<String>> partition(String s) {
        res = new ArrayList<>();
        subset = new ArrayList<>();
        backtrack(s, 0, 0);
        return res;
    }

    private void backtrack(String s, int start, int end) {
        // if out of boundaries, return
        if (end >= s.length()) {
            if (start == end) {
                res.add(new ArrayList<>(subset));
            }
            return;
        }
        if (isPalindrome(s, start, end)) {
            subset.add(s.substring(start, end + 1));
            backtrack(s, end + 1, end + 1);
            subset.removeLast();
        }
        backtrack(s, start, end + 1);
    }

    private boolean isPalindrome(String s, int l, int r) {
        while (l < r) {
            if (s.charAt(l) != s.charAt(r)) return false;
            l++;
            r--;
        }
        return true;
    }
}