package org.masud;

import java.util.Stack;

/**
 * @author Mahfuzur Rahman
 */
public class MyQueue {
  private final Stack<Integer> stack = new Stack<>();

  public MyQueue() {

  }

  public void push(int x) {
    stack.insertElementAt(x, 0);
  }

  public int pop() {
    return stack.isEmpty()? 0:  stack.pop();
  }

  public int peek() {
    return stack.isEmpty()? 0 : stack.get(stack.size() -1);
  }

  public boolean empty() {
    return stack.isEmpty();
  }
}

