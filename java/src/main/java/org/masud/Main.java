package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class Main {
  public static void main(String[] args) {
    System.out.println("Warming up...");
    for (int i = 0; i < 1000; i++) {
      ListNode.isPalindrome2(ListNode.from(new int[]{1, 2, 2, 1}));
    }

    System.out.println("Firing up...");
    long start = System.currentTimeMillis();
    for (int i = 0; i <5000_000; i++) {
      ListNode.isPalindrome2(ListNode.from(new int[]{1, 2, 2, 1}));
    }
    long end = System.currentTimeMillis();

    System.out.printf("Time taken: %s milli %s sec%n", end-start , (end-start)/1000);
  }
}