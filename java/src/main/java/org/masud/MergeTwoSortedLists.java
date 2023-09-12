package org.masud;

import java.util.Objects;

/**
 * @author Mahfuzur Rahman
 */
public class MergeTwoSortedLists {


  public static class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
      this.val = val;
    }

    ListNode(int val, ListNode next) {
      this.val = val;
      this.next = next;
    }

    @Override
    public boolean equals(Object o) {
      if (this == o) return true;
      if (o == null || getClass() != o.getClass()) return false;
      ListNode listNode = (ListNode) o;
      System.out.printf("%s->%s%n", val, listNode.val);
      return val == listNode.val && Objects.equals(next, listNode.next);
    }

    @Override
    public String toString() {
      return val +
        ", " + next ;
    }

    public static ListNode of(int ...v){
      ListNode head = new ListNode();
      ListNode tail = head;

      for (int i : v) {
        tail.next = new ListNode(i);
        tail = tail.next;
      }

      return head.next;
    }


  }


  public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
//    System.out.printf("%s\t%s%n", list1, list2);
    ListNode head = new ListNode();
    ListNode tail = head;


     while (list1 !=null || list2 != null){
       if(list1 == null){
         tail.next = new ListNode(list2.val);
         tail = tail.next;
         list2 = list2.next;
       } else if (list2 == null) {
         tail.next = new ListNode(list1.val);
         tail = tail.next;
         list1 = list1.next;
       }else if(list2.val >= list1.val) {
         tail.next = new ListNode(list1.val);
         tail = tail.next;
         list1 = list1.next;
       }else{
         tail.next = new ListNode(list2.val);
         tail = tail.next;
         list2 = list2.next;
       }
     }

     return head.next;
  }
}
