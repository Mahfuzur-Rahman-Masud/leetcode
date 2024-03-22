package org.masud;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
public class MergeInBetweenTest {
  @Test
  void mergeInBetween() {
    for (int i = 0; i < 1000; i++) {
      ListNode l1 = ListNode.from(0, 1, 2, 3, 4, 5, 6);
      ListNode l2 = ListNode.from(1000, 1001, 1002, 1003, 1004);
      System.out.println(ListNode.mergeInBetween(l1, 2, 5, l2).toString());

      l1 = ListNode.from(0, 1, 2, 3, 4, 5, 6);
      l2 = ListNode.from(1000, 1001, 1002, 1003, 1004);
      System.out.println(ListNode.mergeInBetween(l1, 2, 2, l2).toString());

      l1 = ListNode.from(0, 1, 2, 3, 4, 5, 6);
      l2 = ListNode.from(1000, 1001, 1002, 1003, 1004);
      System.out.println(ListNode.mergeInBetween(l1, 5, 5, l2).toString());

      l1 = ListNode.from(0, 1, 2, 3, 4, 5, 6);
      l2 = ListNode.from(1000, 1001, 1002, 1003, 1004);
      System.out.println(ListNode.mergeInBetween(l1, 1, 1, l2).toString());

      l1 = ListNode.from(10, 1, 13, 6, 9, 5);
      l2 = ListNode.from(1000000, 1000001, 1000002);
      System.out.println(ListNode.mergeInBetween(l1, 3, 4, l2).toString());
    }
  }
}
