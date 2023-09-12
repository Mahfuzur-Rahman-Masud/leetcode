package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class SearchInsertPositionTest extends SearchInsertPosition{

  @Test
  void name() {
    assert searchInsert(new int[]{1,3,5,6}, 5) ==2;
    assert searchInsert(new int[]{1,3,5,6}, 2) ==1;
    assert searchInsert(new int[]{1,3,5,6}, 7) ==4;
  }
}