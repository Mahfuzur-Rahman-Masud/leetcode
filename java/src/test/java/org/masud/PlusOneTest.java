package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import java.util.Arrays;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class PlusOneTest extends PlusOne{

  @Test
  void name() {
    new Eval<int[], int[]>(this::plusOne)
      .put(new int[]{1,2,3}, new int[]{1,2,4})
      .put(new int[]{4,3,2,1}, new int[]{4,3,2,2})
      .put(new int[]{0}, new int[]{1})
      .put(new int[]{8}, new int[]{9})
      .put(new int[]{9}, new int[]{1,0})
      .kToString(Arrays::toString)
      .vToString(Arrays::toString)
      .toEqual(Arrays::equals)
      .evaluate();

//    int[] ints = {1, 2, 3,4};
//    System.out.printf("i: %s%n", Arrays.toString(ints));
  }
}