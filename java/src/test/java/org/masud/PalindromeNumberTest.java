package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class PalindromeNumberTest extends PalindromeNumber {

  @Test
  void name() {
    Map<Integer, Boolean> cases = new HashMap<>();


    cases.put(121, true);
    cases.put(-121, false);
    cases.put(10, false);
    cases.put(11, true);
    cases.put(111, true);

    for (Entry<Integer, Boolean> entry : cases.entrySet()) {
      System.out.printf("K: %s, V: %s, r: %s%n", entry.getKey(), entry.getValue(), isPalindrome(entry.getKey()));
      assert entry.getValue() == isPalindrome(entry.getKey());
    }
  }
}