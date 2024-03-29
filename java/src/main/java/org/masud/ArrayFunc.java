package org.masud;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

/**
 * @author Mahfuzur Rahman
 */
public class ArrayFunc {

  public static class NumArray {
    private final int[] runningSum;

    public NumArray(int[] nums) {
      runningSum = new int[nums.length];
      int sum = 0;
      for (int i = 0; i < nums.length; i++) {
        sum += nums[i];
        runningSum[i] = sum;
      }
    }

    public int sumRange(int left, int right) {
      return left == 0 ? runningSum[right]
          : runningSum[right] - runningSum[left - 1];
    }
  }

  public void moveZeroes(int[] nums) {
    int zeroes = 0;
    for (int i = 0; i < nums.length; i++) {
      if (nums[i] == 0) {
        zeroes++;
      } else if (zeroes > 0) {
        int t = nums[i];
        nums[i] = 0;
        nums[i - zeroes] = t;
      }
    }
  }

  public void move_zeroes(int[] nums) {

  }

  public void moveZeroes_first_iter(int[] nums) {
    int l = nums.length;
    int zeroes = 0;

    for (int i = 0; i < l; i++) {
      while (nums[i] == 0 && i < l - zeroes - 1) {
        System.arraycopy(nums, i + 1, nums, i, l - i - zeroes - 1);
        zeroes++;
        nums[l - zeroes] = 0;
      }
    }
  }

  public static boolean containsDuplicate(int[] nums) {
    final Set<Integer> container = new HashSet<>();
    for (int num : nums)
      if (!container.add(num))
        return true;
    return false;
  }

  public boolean containsNearbyDuplicate(int[] nums, int k) {
    int l = nums.length;

    for (int i = 0; i < l - 1; i++)
      for (int j = i + 1; j < i + k + 1 && j < l; j++)
        if (nums[i] == nums[j])
          return true;

    return false;
  }

  public static List<String> summaryRange(int[] nums) {
    List<String> out = new ArrayList<>();

    for (int i = 0; i < nums.length; i++) {
      int start = nums[i];
      while (i + 1 < nums.length && nums[i] + 1 == nums[i + 1])
        i++;
      out.add(start == nums[i] ? start + "" : start + "->" + nums[i]);
    }

    return out;
  }

  public static List<String> summaryRanges(int[] nums) {
    int start = 0;
    int previous = 0;
    int end = nums.length - 1;
    List<String> out = new ArrayList<>();

    for (int i = 0; i < nums.length; i++) {
      int current = nums[i];
      if (i == 0) {
        start = current;
        previous = start;
      }

      if ((long) current - previous > 1) {
        addRange(out, start, previous);
        start = current;
      }

      if (i == end) {
        addRange(out, start, current);
      }

      previous = current;
    }

    return out;
  }

  private static void addRange(List<String> out, int start, int end) {
    out.add(start == end ? start + "" : start + "->" + end);
  }

  public boolean containsNearByDuplicateByMap(int[] nums, int k) {
    final Map<Integer, Integer> valToIndex = new HashMap<>();
    for (int i = 0; i < nums.length; i++) {
      Integer i2 = valToIndex.put(nums[i], i);
      if (i2 != null && i - i2 <= k) {
        return true;
      }
    }

    return false;
  }

  public static boolean hasValueWithinRange(int[] nums, int k) {
    Arrays.sort(nums);
    for (int i = 1; i < nums.length; i++)
      if (nums[i] - nums[i - 1] <= k) {
        return true;
      }
    return false;
  }

  public int findDuplicate(int[] nums) {
    int[] x = new int[500001];

    for (int n : nums) {
      if (x[n] > 0) {
        return n;
      }

      x[n]++;
    }

    return -1;
  }
}
