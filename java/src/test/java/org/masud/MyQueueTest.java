package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class MyQueueTest extends MyQueue{
  @Test
  void test() {
    assert  empty();
    push(5);
    assert  !empty();
    assert peek() == 5;
    assert pop() == 5;
    assert peek() ==0;
    assert empty();


    push(3);
    push(2);
    assert peek() == 3;
    assert pop() ==3;
    assert pop() ==2;
    assert  empty();

  }
}