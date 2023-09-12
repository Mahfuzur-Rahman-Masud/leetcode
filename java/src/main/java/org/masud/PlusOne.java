package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class PlusOne {



  public int[] plusOne(int[] digits_original) {
    int length = digits_original.length;
    int[] digits = new int[length];
    System.arraycopy(digits_original, 0, digits, 0, length);
    int carry = 1;
    int end = length -1;
    int sum;
    for (int i = end; i >=0; i--) {
      sum = digits[i]+carry;
      if(sum < 10) {
        digits[i] = sum;
        carry =0;
        break;
      }

      digits[i] = 0;
    }

    if(carry ==1){
      int[] a = new int[digits.length +1];
      a[0] =1;
      System.arraycopy(digits, 0, a, 1, digits.length);
      return a;
    }

    return digits;
  }
}
