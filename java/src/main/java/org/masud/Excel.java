package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class Excel {
  public String convertToTitle(int col) {
    final StringBuilder sb = new StringBuilder();

    int x;

    while (col > 0){
      x = (col-1) %26;
      sb.insert(0, (char) ('A' + x));
      col = (col-1)/26;
    }

    return sb.toString();
  }

  public int titleToNumber(String columnTitle) {
    int l = columnTitle.length();
    int out =0;
    for (int i = 0; i < l; i++) {
      int iReversed = l - 1 - i;
      int c = columnTitle.charAt(iReversed) - 64;
      System.out.printf("i: %s\t c: %s%n", i, c);
      out+= (int) (Math.pow(26, i) *c);
    }

    return out;
  }

  public int pow(int n, int p) {

    if(p ==0){
      return 1;
    }

    if(p==1 || n == 0 || n ==1){
      return n;
    }

    int out = n;

    for (int i = 2  ; i <= p; i++) {
      out *= n;
    }

    return out;
  }
}
