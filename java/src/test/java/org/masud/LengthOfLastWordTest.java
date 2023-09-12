package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class LengthOfLastWordTest extends LengthOfLastWord{

  @Test
  void name() {
    new Eval<String ,Integer>(this::lengthOfLastWord)
      .put("Hello World", 5)
      .put("   fly me   to   the moon  ", 4)
      .put("luffy is still joyboy", 6)
      .put("", 0)
      .put(" ", 0)
      .put("q", 1)
      .put("qq", 2)
      .evaluate();
  }
}