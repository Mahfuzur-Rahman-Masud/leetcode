package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import java.util.List;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class RemoveDuplicateFromSortedArrayTest extends RemoveDuplicateFromSortedArray{


  @Test
  void name() {
    int[] a = new int[]{1,1,2};
    int[] a1 = new int[]{1,2,2};
    int al = 2;


    int[] b = new int[]{0,0,1,1,1,2,2,3,3,4};
    int[] b1 = new int[]{0,1,2,3,4};
    int bl = 5;


    int a1l = removeDuplicates(a);
    assert al == a1l;
    equal(a, a1, a1l);


    int b1l = removeDuplicates(b);
    assert bl == b1l;
    equal(b, b1, b1l);

  }

  public static void equal(int[] a, int[] b, int len){

    assert a.length >= len;
    assert b.length >=len;
    System.out.printf("len: %s\tin: %s\tout: %s%n", len, List.of(a), List.of(b));

    for (int i = 0; i <len; i++) {
      assert a[i] == b[i];
    }
  }
}