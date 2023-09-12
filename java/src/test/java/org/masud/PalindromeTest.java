package org.masud;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class PalindromeTest extends Palindrome{
  @Test
  void isPalindromeTest() {
    new Eval<>(this::isPalindrome)
      .put("panama", false)
      .put("a", true)
      .put("aa", true)
      .put("aba", true)
      .put("A man, a plan, a canal: Panama", true)
      .put("99", true)
      .put("181", true)
      .put("race a car", false)
      .put(" ", true)
      .put("", true)
      .put("ab_a", true)
      .put("ab-a", true)
      .put("ab+a", true)
      .put("ab.a", true)
      .evaluate();
    ;


    isPalindrome_slow("A man, a plan, a canal: Panama");
  }
}