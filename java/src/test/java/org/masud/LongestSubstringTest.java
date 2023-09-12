package org.masud;

import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class LongestSubstringTest extends LongestSubstring{

  @Test
  void testSolution() {
    Map<String, Integer> cases = new HashMap<>();
    cases.put("abcabcbb", 3);
    cases.put("bbbbb", 1);
    cases.put("pwwkew", 3);
    cases.put(" ", 1);
    cases.put("tmmzuxt", 5);

    for (Entry<String, Integer> entry : cases.entrySet()) {
      System.out.printf("K: %s v: %s r: %s%n", entry.getKey(), entry.getValue(), lengthOfLongestSubstring(entry.getKey()));
      assert lengthOfLongestSubstring(entry.getKey()) == entry.getValue();
    }
  }
}