package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class ValidParenthesesTest extends ValidParentheses{

  @Test
  void name() {
    new Eval<>(this::isValid)
      .put("()", true)
      .put("(){}[]", true)
      .put("()[]{}", true)
      .put("(]{}", false)
      .put("({})", true)
      .put("({[]})", true)
      .evaluate();
  }
}