import java.util.HashSet;
import java.util.Set;

class Solution929 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String[] emails = {"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com", "testemail@leetcode.com"};
    System.out.println(sol.numUniqueEmails(emails));
  }
}

class Solution {
  public int numUniqueEmails(String[] emails) {
    Set<String> unique = new HashSet<>();
    for (String email : emails) {
      unique.add(getEmail(email));
    }
    return unique.size();
  }

  private static String getEmail(String email) {
    StringBuilder sb = new StringBuilder();
    int splitIdx = email.indexOf("@");
    for (int i = 0; i < splitIdx; i++) {
      char c = email.charAt(i);
      if (c == '+') break;
      if (c == '.') continue;
      sb.append(c);
    }
    sb.append(email.substring(splitIdx));
    return sb.toString();
  }
}
