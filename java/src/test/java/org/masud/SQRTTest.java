package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class SQRTTest extends SQRT{

  @Test
  void name() {
    new Eval<>(this::mySqrt)
      .put(4, 2)
      .put(8, 2)
      .put(1, 1)
      .put(0, 0)
      .put(11, 3)
      .put(10, 3)
      .evaluate();;
  }

  @Test
  void intdiv() {
    int a = 11;
    int b = 3;
    System.out.printf("val: %s%n", a/b);
  }
}