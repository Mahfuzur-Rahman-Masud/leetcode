package org.masud;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class RomanToIntegerTest extends RomanToInteger {

  @Test
  void name() {
    Eval<String, Integer> eval = new Eval<>(this::romanToInt);
    eval.put("III", 3);
    eval.put("LVIII", 58);
    eval.put("MCMXCIV", 1994);

    eval.evaluate();

  }
}