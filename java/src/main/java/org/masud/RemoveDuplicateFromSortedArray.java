package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class RemoveDuplicateFromSortedArray {


  public int removeDuplicates(int[] nums) {
    int marker =0;
    int edge = 1;

    while (edge < nums.length){
      if(nums[marker] != nums[edge]){
        marker++;
        if(marker!= edge){
          nums[marker] = nums[edge];
        }
      }

      edge++;
    }

    return marker+1;
  }

}
