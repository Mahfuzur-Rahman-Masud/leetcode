package org.masud;

import org.junit.jupiter.api.Test;
import org.masud.MergeTwoSortedLists.ListNode;

/**
 * @author Mahfuzur Rahman
 */
class MergeTwoSortedListsTest extends MergeTwoSortedLists{


  @Test
  void name() {
    new Eval<ListNode[], ListNode>(a->mergeTwoLists(a[0], a[1]))
      .put(new ListNode[]{ListNode.of(1,2,4), ListNode.of(1,3,4)}, ListNode.of(1,1,2,3,4,4))
      .put(new ListNode[]{ListNode.of(), ListNode.of()}, ListNode.of())
      .put(new ListNode[]{ListNode.of(), ListNode.of(0)}, ListNode.of(0))
      .evaluate();

    ;
  }

  @Test
  void testEqual() {
    ListNode l1 = new ListNode(5);
    l1.next = new ListNode(6);
    l1.next.next = new ListNode(7);

    ListNode l2 = new ListNode(5);
    l2.next = new ListNode(6);
    l2.next.next = new ListNode(7);
    assert l1.equals(l2);

    assert ListNode.of(5,6,7).equals(l2);
  }
}