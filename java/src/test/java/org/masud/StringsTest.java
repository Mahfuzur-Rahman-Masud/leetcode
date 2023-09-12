package org.masud;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class StringsTest extends Strings {

  @Test
  void reverseVowelsTest() {
    Assertions.assertEquals("leotcede", reverseVowels("leetcode"));
  }

  @Test
  void isIsoMorphicTest() {
    assert isIsomorphic("egg", "add");
    assert !isIsomorphic("foo", "bar");
    assert isIsomorphic("paper", "title");
  }

  @Test
  void isAnagramTest() {
    assert isAnagram("anagram", "nagaram");
    assert !isAnagram("cat", "rat");
    assert isAnagram("cat", "tac");
  }

  @Test
  void wordPatternTest() {
    assert wordPattern("abba", "dog cat cat dog");
    assert !wordPattern("abba", "dog dog dog dog");
    assert !wordPattern("aaaa", "dog cat cat dog");
  }
}