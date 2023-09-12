package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class LengthOfLastWord {
  public int lengthOfLastWord(String s) {

    int i = s.length() - 1;
    int li = -1;

    while (i > -1) {
      char c = s.charAt(i);
      if (c == ' ' && li != -1) {
        return li - i;
      } else if (li == -1 && c != ' ') {
        li = i;
      }

      i--;
    }

    return li+1;
  }
}
