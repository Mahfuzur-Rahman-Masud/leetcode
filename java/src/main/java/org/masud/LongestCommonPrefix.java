package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class LongestCommonPrefix {
  public String longestCommonPrefix(String[] strs) {
    if (strs.length == 0) {
      return "";
    }

    char c;

    int i = 0;
    for (; i < strs[0].length(); i++) {
      c = strs[0].charAt(i);
      for (int j = 1; j < strs.length; j++) {
          String target = strs[j];
          if (target.length() <= i || target.charAt(i) != c) {
            return strs[0].substring(0, i);
          }
      }
    }

    return strs[0].substring(0, i);
  }
}
