import java.util.Stack;

class Solution2116 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    String s = "))()))(())((";
    String locked = "010100000000";
    System.out.println(sol.canBeValid(s, locked));
  }
}

class Solution {
  public boolean canBeValid(String s, String locked) {
    int n = s.length();
    if (n % 2 == 1)
      return false;
    Stack<Integer> lockedStack = new Stack<>();
    Stack<Integer> unlockedStack = new Stack<>();
    for (int i = 0; i < n; i++) {
      char c = s.charAt(i);
      boolean isLocked = locked.charAt(i) == '1';
      if (!isLocked) {
        unlockedStack.push(i);
      } else if (c == '(') {
        lockedStack.push(i);
      } else { // c == ')' and is locked
        if (lockedStack.isEmpty() && unlockedStack.isEmpty())
          return false;
        if (!lockedStack.isEmpty()) {
          lockedStack.pop();
        } else {
          unlockedStack.pop();
        }
      }
    }
    // check if there are unlocked entites that precede remaining locked entites
    while (!lockedStack.isEmpty() && !unlockedStack.isEmpty() && lockedStack.peek() < unlockedStack.peek()) {
      unlockedStack.pop();
      lockedStack.pop();
    }
    // if there are locked entites that could not match with unlocked entites
    if (!lockedStack.isEmpty())
      return false;
    return unlockedStack.size() % 2 == 0;

    // [X ) X ) X X]

    // [X ( X ( X X]

    // [X ) ( X ( X X]
  }

  private char getOppositeParentheses(char c) {
    return c == ')' ? '(' : ')';
  }
}