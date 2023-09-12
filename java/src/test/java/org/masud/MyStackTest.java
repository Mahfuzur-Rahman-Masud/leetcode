package org.masud;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class MyStackTest extends MyStack {


  @Test
  void stackTest() {
    push(2);
    assert top() == 2;
    assert pop() == 2;
    assert top() == 0;
    assert empty();

    for (int i = 0; i < 560; i++ ){
      push(i);
    }
    System.out.printf("Size: %s%n", size());

    for (int i = 0; i < 560; i++ ){
      assert pop() == 559-i;
    }

  }
}