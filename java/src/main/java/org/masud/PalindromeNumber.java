package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class PalindromeNumber {

  public boolean isPalindrome(int x) {
    if(x == 0 ){
      return true;
    }

    if(x<0){
      return  false;
    }

    String str = String.valueOf(x);
    int length = str.length();
    for (int i = 0; i < length /2; i++) {
      int endIndex = length-1 -i;
      if(str.charAt(i) != str.charAt(endIndex)){
        return false;
      }
    }


    return true;
  }

  public static void main(String[] args) {
    System.out.printf("by: %s<> %s%n", 3/2, 5);
  }
}
