package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class SearchInsertPosition {

  public int searchInsert(int[] nums, int target) {
    int l = nums.length;
    int i = l/2;
    int ub = l-1;
    int lb = 0;
    int v;

    while (true) {
      v = nums[i];
      if (v == target) {
        return i;
      } else if (v > target) {
        ub = i - 1;
      } else {
        lb = i + 1;
      }
      if (lb> ub) {
        return lb;
      } else if (ub< 0) {
        return 0;
      }

      i = (ub - lb) / 2 + lb;
    }

  }
}
