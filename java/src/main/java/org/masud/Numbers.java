package org.masud;

import java.util.HashSet;
import java.util.Set;

/**
 * @author Mahfuzur Rahman
 */
public class Numbers {


  public boolean isPowerOfThree(int n) {
    if(n<1)return false;
    while(n%3==0)n/=3;
    return n==1;
  }

  /**
   * Write an algorithm to determine if a number n is happy.
   * <p>
   * A happy number is a number defined by the following process:
   * <p>
   * Starting with any positive integer, replace the number by the sum of the squares of its digits.
   * Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
   * Those numbers for which this process ends in 1 are happy.
   * Return true if n is a happy number, and false if not.
   */
  public boolean isHappy(int n) {
    Set<Integer> seen = new HashSet<>();

    while (n != 1 && !seen.contains(n)) {
      seen.add(n);
      n = sumOfPowOfDigits(n);
    }

    return n == 1;
  }

  public int missingNumber(int[] nums) {
    long sum =0;
    long expect = 0;

    for (int i = 0; i < nums.length; i++) {
      expect +=i+1;
      sum +=nums[i];
    }

    return (int) (expect - sum);
  }


  public boolean isUgly(int n) {
    if (n < 1) return false;
    if (n == 1 || n == 2 || n == 3 || n == 5) return true;

    int x = n;
    if (x % 2 == 0) x /= 2;
    if (x % 3 == 0) x /= 3;
    if (x % 5 == 0) x /= 5;

    return x != n && isUgly(x);
  }

  public int addDigits(int num) {
    while (num > 9) {
      int sum = 0;
      int next = num;
      while (next > 0) {
        sum += next % 10;
        next /= 10;
      }
      num = sum;
    }
    return num;
  }

  public boolean isPowerOfTwo(int n) {
    return ((n << 1) - 1) == (n | (n - 1));
  }

  public int sumOfPowOfDigits(int n) {
    int sum = 0;
    int rem;
    while (n > 0) {
      rem = n % 10;
      n /= 10;
      sum += rem * rem;
    }
    return sum;
  }


  public boolean isHappyFS(int n) {
    int slow = n;
    int fast = n;

    do {
      slow = sumOfPowOfDigits(slow);
      fast = sumOfPowOfDigits(sumOfPowOfDigits(fast));
    } while (slow != fast);

    return slow == 1;
  }

  public boolean isHappyFS2(int n) {
    int slow = n;
    int fast = n;

    do {
      slow = sumOfPowOfDigits(slow);
      fast = sumOfPowOfDigits(sumOfPowOfDigits(fast));

    } while (slow != fast);

    return slow == 1;
  }
}
