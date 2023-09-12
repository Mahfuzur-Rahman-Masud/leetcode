package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class FindTheIndexOfTheFirstOccurrenceInAStringTest extends FindTheIndexOfTheFirstOccurrenceInAString {

  @Test
  void name() {
    new Eval<String[], Integer>(x->strStr2(x[0], x[1]))
      .put(new String[]{"sadbutsad", "sad"}, 0)
      .put(new String[]{"leetcode", "leeto"}, -1)
      .put(new String[]{"aaa", "aaaa"}, -1)
      .put(new String[]{"mississippi", "issipi"}, -1)
      .evaluate();
      ;
  }
}