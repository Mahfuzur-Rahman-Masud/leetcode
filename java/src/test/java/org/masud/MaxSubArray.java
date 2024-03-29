package org.masud;

import java.util.HashMap;
import java.util.Map;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
public class MaxSubArray {


  @Test
  void countSubarraysTest() {
    assert  6 == countSubarrays(new int[]{1, 3, 2, 3, 3},2 );
    assert  0 == countSubarrays(new int[]{1, 4, 2, 1}, 3);
    assert  1 == countSubarrays(new int[]{1}, 1);
    assert  2 == countSubarrays(new int[]{1,2}, 1);
  }

  public long countSubarrays(int[] nums, int k) {
    int max = 0;
    for (int n : nums) {
      if (n> max){
        max =n;
      }
    }

    long out=0;
    int mark = 0;
    int count = 0;


    for (int n : nums) {
      if (n== max){
        count++;
      }

      while (count >=k){
        if (nums[mark]==max){
          count--;
        }
        mark++;
      }

      out+=mark;
    }



    return out;
  }
  
  @Test
  void maxSubarrayLength_test() {
    assert 6 == maxSubarrayLength(new int[]{1,2,3,1,2,3,1,2}, 2);
    assert 2 == maxSubarrayLength(new int[]{1,2,1,2,1,2,1,2}, 1);
    assert 4 == maxSubarrayLength(new int[]{5,5,5,5,5,5,5}, 4);
    assert 1 == maxSubarrayLength(new int[]{1}, 1);
    assert 1 == maxSubarrayLength(new int[]{1}, 10);
  }

  
  
  public int maxSubarrayLength(int[] nums, int k) {

    int start = 0;
    int cur = 0;
    int ml = 0;
    int ln = 0;
    int l = nums.length;
    int v = 0;


    Map<Integer, Integer> m = new HashMap<>();

    while (cur< l){
      v = nums[cur];
      int c = m.getOrDefault(v, 0);
      m.put(v, c+1);

      if (c!= k){
        ln = cur-start+1;
        if (ln>ml){
          ml=ln;
        }
      }else{
        int v2 = 0;
        for (int i = start; i < cur ; i++) {
          v2= nums[i];
          m.put(v2, m.get(v2) -1);
          if (v2 == v){
            start = i+1;
            if(l-start < ml){
              return ml;
            }

            break;
          }
        }

      }

      cur++;
    }


    return ml;
  }
}
