package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class RomanToInteger {



  private int romanToInt(char c){
    switch (c){
      case 'I':
        return 1;
      case 'V':
          return 5;
      case 'X':
        return 10;
      case 'L':
        return 50;
      case 'C':
        return 100;
      case 'D':
        return 500;
      case 'M':
        return 1000;
    }
    return 0;
  }
  public int romanToInt(String s) {
    int last =0;
    int total =0;
    for (int i = s.length() -1; i >=0; i--) {
      int c  = romanToInt(s.charAt(i));
      if (last > c){
        total -=c;
      }else{
        total +=c;
      }
      last = c;
    }
    return total;
  }

}
