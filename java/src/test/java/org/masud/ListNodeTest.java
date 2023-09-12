package org.masud;

import static org.masud.ListNode.isPalindrome;
import static org.masud.ListNode.isPalindrome2;
import static org.masud.ListNode.reverseList;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class ListNodeTest {


  @Test
  void removeElementTest() {
    ListNode node = new ListNode(1);
    node.add(2).add(2).add(1);
    ListNode result = node.removeElements(node, 2);
    System.out.printf("Result: %s%n", result);
  }


  @Test
  void reverseListTest() {
    ListNode node = ListNode.from(1, 2, 3, 4, 5);
    assert  reverseList(node).equals(ListNode.from(5,4,3,2,1));
    assert reverseList(ListNode.from(1)).equals(ListNode.from(1));
    assert reverseList(ListNode.from(2,1)).equals(ListNode.from(1,2));

    System.out.printf("%s%n", node);
    ListNode reversed = reverseList(node);
    System.out.printf("%s%n", reversed);
  }


  @Test
  void isPalindromeTest() {
    assert isPalindrome( ListNode.from(1,2,2,1));
    assert isPalindrome( ListNode.from(1,2,1));
    assert isPalindrome( ListNode.from(1));
    assert isPalindrome( ListNode.from(1,2,3,4,3,2,1));
    assert isPalindrome( ListNode.from(1,2,3,4,4,3,2,1));
    assert !isPalindrome( ListNode.from(1,2,3,4,4,5,3,2,1));
    assert !isPalindrome( ListNode.from(1,2));
  }


  @Test
  void isPalindromeTest2() {
    assert isPalindrome2( ListNode.from(1,2,2,1));
    assert isPalindrome2( ListNode.from(1,2,1));
    assert isPalindrome2( ListNode.from(1));
    assert isPalindrome2( ListNode.from(1,2,3,4,3,2,1));
    assert isPalindrome2( ListNode.from(1,2,3,4,4,3,2,1));
    assert !isPalindrome2( ListNode.from(1,2,3,4,4,5,3,2,1));
    assert !isPalindrome2( ListNode.from(1,2));
  }
}