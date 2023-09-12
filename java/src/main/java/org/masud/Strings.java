package org.masud;

import java.util.HashMap;
import java.util.Map;
import java.util.Map.Entry;

/**
 * @author Mahfuzur Rahman
 */
public class Strings {

  public String reverseVowels(String s) {
    char[] c = s.toCharArray();

    int l = 0, h = s.length() - 1;
    while (l < h) {
      switch (Character.toLowerCase(c[l])) {
        case 'a', 'e', 'i', 'o', 'u':
          switch (Character.toLowerCase(c[h])) {
            case 'a', 'e', 'i', 'o', 'u':
              char cc = c[l];
              c[l] = c[h];
              c[h] = cc;
              l++;
              h--;
              break;
            default:
              h--;
          }
          break;
        default:
          l++;
      }
    }

    return new String(c);
  }

  public boolean wordPattern(String pattern, String s) {
    String[] words = s.split(" ");
    if (pattern.length() != words.length) {
      return false;
    }

    Map<String, Byte> wordToChar = new HashMap<>();

    byte[] bytes = pattern.getBytes();
    String[] buffer = new String[26];
    for (int i = 0; i < words.length; i++) {
      if (buffer[bytes[i] - 97] == null) {
        if (wordToChar.put(words[i], bytes[i]) != null) {
          return false;
        }
        buffer[bytes[i] - 97] = words[i];
      } else if (!buffer[bytes[i] - 97].equals(words[i])) {
        return false;
      }
    }

    return true;
  }


  public boolean isIsomorphic(String s, String t) {
    if (s.length() != t.length()) {
      return false;
    }

    int[] is = new int[200];
    int[] it = new int[200];
    byte[] bs = s.getBytes();
    byte[] bt = t.getBytes();

    for (int i = 0; i < bs.length; i++) {
      int ss = bs[i];
      int tt = bt[i];

      if (is[ss] != it[tt]) {
        return false;
      }

      int x = i + 1;
      is[ss] = x;
      it[tt] = x;

    }

    return true;
  }

  public boolean isAnagram(String s, String t) {
    if (s.length() != t.length()) {
      return false;
    }

    Map<Integer, Integer> a = new HashMap<>();
    Map<Integer, Integer> b = new HashMap<>();
    for (int i = 0; i < s.length(); i++) {
      a.compute((int) s.charAt(i), (k, old) -> old == null ? 1 : old + 1);
      b.compute((int) t.charAt(i), (k, old) -> old == null ? 1 : old + 1);
    }

    for (Entry<Integer, Integer> entry : a.entrySet()) {
      if (!entry.getValue().equals(b.get(entry.getKey()))) {
        return false;
      }
    }

    return true;
  }
}
