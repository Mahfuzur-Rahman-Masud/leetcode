package org.masud;

import java.util.HashSet;
import java.util.Objects;
import java.util.Set;

/**
 * @author Mahfuzur Rahman
 */
public class ListNode {
  int val;
  ListNode next;

  ListNode(int x) {
    val = x;
    next = null;
  }

  public ListNode add(int val){
    next = new ListNode(val);
    return next;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    ListNode node = (ListNode) o;
    return val == node.val && Objects.equals(next, node.next);
  }

  @Override
  public int hashCode() {
    return Objects.hash(val, next);
  }


  public static boolean isPalindrome2(ListNode head) {
    if (head == null) {
      return false;
    }

    if (head.next == null) {
      return true;
    }

    // go to middle

    ListNode fast = head;
    ListNode slow = head;

    while (fast != null && fast.next != null) {
      fast = fast.next.next;
      slow = slow.next;
    }


    // Break the list in the middle and reverse 2nd half of it
    ListNode previous = null;
    ListNode next;
    ListNode current = slow;

    while (current != null) {
      next = current.next;
      current.next = previous;
      previous = current;
      current = next;
    }

    ListNode tail = previous;

    while (head != null && tail != null) {
      if (head.val != tail.val) return false;
      head = head.next;
      tail = tail.next;
    }
    return true;
  }


  public static boolean isPalindrome(ListNode head) {
    if (head == null) {
      return false;
    }

    if (head.next == null) {
      return true;
    }

    int len = 0;
    int[] c = new int[64];

    while (head != null) {
      if (len == c.length) {
        int[] c2 = new int[c.length * 2];
        System.arraycopy(c, 0, c2, 0, len);
        c = c2;
      }

      c[len] = head.val;
      len++;
      head = head.next;
    }

    int exit = len / 2 + 1;

    for (int i = 0; i < exit; i++) {
      if (c[i] != c[len - i - 1]) {
        return false;
      }
    }

    return true;
  }


  public static ListNode from (int... values){
    ListNode head = new ListNode(values[0]);

    ListNode c =head;
    for (int i = 1; i < values.length; i++) {
      c = c.add(values[i]);
    }

    return head;
  }


  public static ListNode reverseList(ListNode head) {
    if(head == null){
      return null;
    }

    ListNode h = head;
    ListNode temp;
    ListNode c = head.next;
    h.next= null;

    while (c != null){
      temp = c;
      c = c.next;
      temp.next = h;
      h = temp;
    }

    return h;
  }

  @Override
  public String toString() {
    return val + (next == null ? "" : ", " + next);
  }





  public ListNode removeElements(ListNode head, int val) {
    while (head != null &&  head.val == val ){
      head = head.next;
    }

    ListNode current  = head;
    while (current != null){
      if(current.next != null && current.next.val == val){
        current.next = current.next.next;
        continue;
      }

      current = current.next;
    }

    return head;
  }

  public boolean hasCycle(ListNode head) {
    if(head == null) return false;
    Set<ListNode> set = new HashSet<>();
    ListNode current = head;

    while (current != null){
      if(!set.add(current)){
        return true;
      }
      current = current.next;
    }

    return false;
  }


  public boolean hasCycleFastSlow(ListNode head) {
    ListNode fast = head;
    ListNode slow = head;

    while(fast != null && fast.next != null){
      fast= fast.next.next;
      slow = slow.next .next;
      if (fast == slow){
        return true;
      }
    }

    return false;
  }


}
