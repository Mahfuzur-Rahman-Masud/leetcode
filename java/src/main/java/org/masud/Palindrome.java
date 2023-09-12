package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class Palindrome {
  public boolean isPalindrome_slow(String s) {
    String sanitized = s.replaceAll("\\s", "")
      .replaceAll("[^\\w]", "")
      .replaceAll("_", "")
      .toLowerCase();

    int endI = sanitized.length() -1;

    for (int i = 0; i < sanitized.length()/2; i++) {
      if(sanitized.charAt(i) != sanitized.charAt(endI)) return false;
      endI --;
    }

    return true;
  }


  public boolean isPalindrome(String s){
    int markIn =0;
    int markOut = s.length() -1;

    while (markIn< markOut){
      int i = s.charAt(markIn);
      int o = s.charAt(markOut);

      if(i ==o){
        markIn++;
        markOut--;
        continue;
      }

      if(i <48 || (i>57 && i<65) || (i> 90 && i<97 ) || i >122) {
        markIn++;
        continue;
      }

      if(o <48 || (o>57 && o<65) || (o> 90 && o<97 ) || o >122 ){
        markOut--;
        continue;
      }

      if(i >64 && i < 91){
        i = i +  32;
      }

      if(o >64 && o < 91){
        o =  o+32;
      }

      if(i != o){
        return false;
      }

      markIn++;
      markOut--;
    }

    return true;
  }
}
