package org.masud;

/**
 * @author Mahfuzur Rahman
 */
class MyStack {
  private int[] c;
  private int mark = -1;

  public MyStack() {
    c = new int[64];
  }

  public void push(int x) {
    mark++;
    if (c.length < mark + 1) {
      int[] c = new int[mark * 2];
      System.arraycopy(this.c, 0, c, 0, this.c.length);
      this.c = c;
    }

    c[mark] = x;
  }

  public int pop() {
    if (mark < 0) {
      return 0;
    }

    int out = c[mark];
    mark--;
    return out;
  }


  public int size(){
    return mark +1;
  }
  public int top() {
    return mark < 0 ? 0 : c[mark];
  }

  public boolean empty() {
    return mark == -1;
  }
}
