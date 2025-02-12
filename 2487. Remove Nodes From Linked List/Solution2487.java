class Solution2487 {
  public static void main(String[] args) {
    Solution sol = new Solution();
    ListNode temp;
    ListNode head = new ListNode(5);
    temp = head.next = new ListNode(2);
    temp = temp.next = new ListNode(13);
    temp = temp.next = new ListNode(3);
    temp = temp.next = new ListNode(8);
    ListNode res = sol.removeNodes(head);
    res.print();
  }
}

class Solution {
  public ListNode removeNodes(ListNode head) {
    ListNode headRev = reverse(head);
    int max = headRev.val;
    ListNode curr = headRev;
    ListNode prev = null;
    while (curr != null) {
      int val = curr.val;
      if (val < max) {
        prev.next = curr.next;
      } else {
        prev = curr;
      }
      max = Math.max(max, val);
      curr = curr.next;
    }
    return reverse(headRev);
  }

  private ListNode reverse(ListNode head) {
    ListNode curr = head;
    ListNode prev = null;
    while (curr != null) {
      ListNode temp = curr.next;
      curr.next = prev;
      prev = curr;
      curr = temp;
    }
    return prev;
  }
}

class ListNode {
  int val;
  ListNode next;
  ListNode() {}
  ListNode(int val) { this.val = val; }
  ListNode(int val, ListNode next) { this.val = val; this.next = next; }
  
  public void print() {
    ListNode curr = this;
    while (curr != null) {
      System.out.println(curr.val);
      curr = curr.next;
    }
  }
}
