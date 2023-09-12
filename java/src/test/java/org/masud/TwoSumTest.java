package org.masud;

import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class TwoSumTest extends TwoSum {

  private final Map<int[], Integer> problems = new HashMap<>();

  @Test
  void testTwoSum() {
    problems.put(new int[]{2, 7, 11, 15}, 9);

    for (Entry<int[], Integer> entry : problems.entrySet()) {
      int[] key = entry.getKey();
      Integer value = entry.getValue();

      int[] ints = twoSum(key, value);
      assert key[ints[0]] + key[ints[1]] == value;
    }

    System.out.println("Success!");
  }
}