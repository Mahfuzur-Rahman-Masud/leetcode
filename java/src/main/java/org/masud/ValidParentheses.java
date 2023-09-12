package org.masud;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Mahfuzur Rahman
 */
public class ValidParentheses {
  public boolean isValid(String s) {
    List<Character> expected = new ArrayList<>();
    for (int i = 0; i < s.length(); i++) {
      char c = s.charAt(i);

      if (c == '(') {
        expected.add(')');
      } else if (c == '{') {
        expected.add('}');
      } else if (c == '[') {
        expected.add(']');
      } else if (expected.isEmpty() || expected.remove(expected.size() - 1) != c) {
        return false;
      }
    }

    return expected.isEmpty();
  }
}
