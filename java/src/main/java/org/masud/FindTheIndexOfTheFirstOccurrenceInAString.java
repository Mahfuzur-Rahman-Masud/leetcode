package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class FindTheIndexOfTheFirstOccurrenceInAString {

  public int strStr(String haystack, String needle) {

    return haystack.indexOf(needle);
  }

  public int strStr2(String haystack, String needle) {
    if(needle.length()> haystack.length()){
      return -1;
    }

    int last = needle.length()-1;
    int marker ;
    int hlen = haystack.length();

    for (int i = 0; i < haystack.length(); i++) {
      marker =0;
      while (i+marker < hlen && haystack.charAt(i+marker)== needle.charAt(marker)   ){
        if(marker == last ){
          return i;
        }
        marker++;
      }
    }
   return -1;
  }
}
