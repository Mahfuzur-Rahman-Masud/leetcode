class Solution {
  int findMaxConsecutiveOnes(List<int> nums) {
    var mx = 0;
    var count = 0;
    for (var n in nums) {
      count = count * n + n;
      mx = count > mx ? count : mx;
    }

    return mx;
  }

  int lastStoneWeight(List<int> stones) {
    if (stones.isEmpty) {
      return 0;
    }

    if (stones.length == 1) {
      return stones[0];
    }

    stones.sort((a, b) => a - b);
    var n = stones.removeLast() - stones.removeLast();
    stones.add(n);

    return lastStoneWeight(stones);
  }

  List<int> nextGreaterElement(List<int> nums1, List<int> nums2) {
    List<int> list = [];
    Map<int, int> idxMap = {};
    var i = 0;
    for (var n in nums2) {
      idxMap[n] = i++;
    }

    for (var n in nums1) {
      list.add(-1);

      for (var i = idxMap[n]!; i < nums2.length; i++) {
        if (nums2[i] > n) {
          list[list.length - 1] = nums2[i];
          break;
        }
      }
    }

    return list;
  }
}
