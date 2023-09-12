package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class ExcelTest extends  Excel{

  @Test
  void convertToTitleTest() {
    new Eval<>(this::convertToTitle)
      .put(1, "A")
      .put(5, "E")
      .put(26, "Z")
      .put(28, "AB")
      .put(701, "ZY")
      .evaluate();
    ;
  }

  @Test
  void name() {
    new Eval<>(this::titleToNumber)
      .put( "A", 1)
      .put( "E", 5)
      .put( "Z", 26)
      .put( "AB",28)
      .put( "ZY", 701)
      .evaluate();
  }

  @Test
  void testPow() {
    new Eval2<>(this::pow)
      .put(1, 0, 1)
      .put(1, 1, 1)
      .put(0, 1, 0)
      .put(1000, 0, 1)
      .put(1000, 1, 1000)
      .put(2, 2, 4)
      .put(2, 3, 8)
      .evaluate();


    long start = System.currentTimeMillis();
    for (int i = 0; i < 10000; i++) {
      int pow = pow(500, 10);
    }

    long check1 = System.currentTimeMillis();
    for (int i = 0; i < 10000; i++) {
      double pow = Math.pow(500, 10);
    }
    long check2 = System.currentTimeMillis();

    System.out.printf("pow: %s\nMath.pow:%s\n", check1 - start, check2 - check1);
  }
}