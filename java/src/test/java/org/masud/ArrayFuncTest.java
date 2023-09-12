package org.masud;

import java.util.Arrays;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class ArrayFuncTest extends ArrayFunc {

  @Test
  void moveZeroesTest() {
    int[] a = new int[]{0,1,0,3,12};
     moveZeroes(a);
    System.out.println(Arrays.toString(a));
    Assertions.assertArrayEquals(a, new int[]{1,3,12,0,0});

    a = new int[]{0};
    moveZeroes(a);
    Assertions.assertArrayEquals(a, new int[]{0});


  }

  @Test
  void summaryRangesTest() {
    System.out.printf("%s%n", summaryRanges(new int[]{-2147483648,-2147483647,2147483647}));
    Assertions.assertArrayEquals(summaryRanges(new int[]{-2147483648,-2147483647,2147483647}).toArray( String[]::new)
      , new String[]{"-2147483648->-2147483647","2147483647"});

    Assertions.assertArrayEquals(summaryRange(new int[]{-2147483648,-2147483647,2147483647}).toArray( String[]::new)
      , new String[]{"-2147483648->-2147483647","2147483647"});
  }

  @Test
  void containsCloseNumberTest() {
    assert  ArrayFunc.hasValueWithinRange(new int[]{1,2,3,1}, 3 );
    assert  ArrayFunc.hasValueWithinRange(new int[]{1,4}, 3 );
    assert  !ArrayFunc.hasValueWithinRange(new int[]{-2,-6}, 3 );
    assert  ArrayFunc.hasValueWithinRange(new int[]{-2,-5}, 3 );
    assert  !ArrayFunc.hasValueWithinRange(new int[]{-2}, 3 );
  }

  @Test
  void containsNearbyDuplicateTest(){
    assert containsNearbyDuplicate(new int[]{1,2,3,1}, 3);
    assert !containsNearbyDuplicate(new int[]{1,2,3,1}, 2);
    assert containsNearbyDuplicate(new int[]{1,2,3,1}, 4);
    assert !containsNearbyDuplicate(new int[]{1,2}, 3);
    assert containsNearbyDuplicate(new int[]{1,1}, 3);
    assert !containsNearbyDuplicate(new int[]{1,2,3,1}, 2);
  }


  @Test
  void containsNearbyDuplicateWithMapTest(){
    assert containsNearByDuplicateByMap(new int[]{1,2,3,1}, 3);
    assert !containsNearByDuplicateByMap(new int[]{1,2,3,1}, 2);
    assert containsNearByDuplicateByMap(new int[]{1,2,3,1}, 4);
    assert !containsNearByDuplicateByMap(new int[]{1,2}, 3);
    assert containsNearByDuplicateByMap(new int[]{1,1}, 3);
    assert !containsNearByDuplicateByMap(new int[]{1,2,3,1}, 2);
  }



}