package org.masud;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Mahfuzur Rahman
 */
public class LongestSubstring {


  // wpwkew
    public int lengthOfLongestSubstring(String s) {
      final Map<Character, Integer> m = new HashMap<>();
      int max = 0;
      int len = 0;
      int i = 0;
      int start = 0;

      for (; i < s.length(); i++) {
        char c = s.charAt(i);
        Integer val = m.get(c);
        m.put(c, i);


        if(val == null || val < start){
          len++;
          if(len > max){
            max = len;
          }
          continue;
        }

        start = val + 1;
        len = i-val;
      }

      return max;
    }
}
