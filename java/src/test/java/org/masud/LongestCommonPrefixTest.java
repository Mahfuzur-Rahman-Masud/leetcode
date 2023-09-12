package org.masud;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class LongestCommonPrefixTest extends LongestCommonPrefix{

  @Test
  void name() {
    Eval<String[], String> e = new Eval<>(this::longestCommonPrefix);
    e
      .put(new String[]{"flower","flow","flight"}, "fl")
      .put(new String[]{"dog","racecar","car"}, "")
      .put(new String[]{"ab","a",}, "a")
      .evaluate();

  }
}