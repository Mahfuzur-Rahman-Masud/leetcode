package org.masud;

/**
 * @author Mahfuzur Rahman
 */
public class Api {
  public int bad = 0;

  boolean isBadVersion(int version) {
    return version >= bad;
  }


  public int firstBadVersion(int n) {
    int high = n, low = 0;

    while (low <= high) {
      int mid = low + (high - low) / 2;
      if (isBadVersion(mid)) {
        high = mid - 1;
      } else {
        low = mid + 1;
      }
    }

    return low;
  }


}
