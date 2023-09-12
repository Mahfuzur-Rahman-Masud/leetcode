package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class NumbersTest extends Numbers{

  @Test
  void missingNumberTest() {
    assert missingNumber(new int[]{3,0,1}) ==2;
    assert missingNumber(new int[]{0,1}) ==2;
    assert missingNumber(new int[]{9,6,4,2,3,5,7,0,1}) ==8;
  }

  @Test
  void addDigitsTest() {
    assert addDigits(38) == 2;
    assert addDigits(0) == 0;
    assert addDigits(11) == 2;
    assert addDigits(55) == 1;
    assert addDigits(555) == 6;
  }

  @Test
  void isHappyTest() {
    new Eval<>(this::isHappy)
      .put(19, true)
      .put(2, false)
      .evaluate();


    new Eval<>(this::isHappyFS)
      .put(19, true)
      .put(2, false)
      .evaluate();


    new Eval<>(this::isHappyFS2)
      .put(19, true)
      .put(2, false)
      .evaluate();
  }


  @Test
  void isHappyBench() {
    long c1 = System.currentTimeMillis();
    for (int i = 0; i < 5000000; i++) {
      isHappy(i *5324 + 4961);
    }

    long c2 = System.currentTimeMillis();
    for (int i = 0; i < 5000000; i++) {
      isHappyFS(i *5324 + 4961);
    }
    long c3 = System.currentTimeMillis();
    for (int i = 0; i < 5000000; i++) {
      isHappyFS2(i *5324 + 4961);
    }
    long c4 = System.currentTimeMillis();


    System.out.printf("Happy:\t%s%nHappyFS:\t%s%nHappyFS2:\t%s%n", c2-c1, c3-c2, c4-c3);
  }
}