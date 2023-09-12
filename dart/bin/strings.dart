class Solution {
  int findLUSlength2(String a, String b) => a == b? -1: a.length>b.length? a.length: b.length;

  int findLUSlength(String a, String b) {
    var l = a.length;
    if (l != b.length) return l > b.length ? l : b.length;

    var max = -1;
    var markIn = 0;
    var markOut = 0;
    var out = l + 1;

    while (markOut < out) {
      var sub = a.substring(markIn, markOut);

      if (!b.contains(sub)) {
        max = sub.length;
      }

      markOut++;

      if (markOut == out) {
        markIn++;
        markOut = max == -1 ? markIn : markIn + max;
      }
    }

    return max;
  }

  bool detectCapitalUse(String word) {
    var l = word.length;
    if (l < 2) {
      return true;
    }

    var b = word.codeUnits;
    var start = b[0] < 97 ? 1 : 0;
    var isCapital = b[start] < 97;

    for (var i = start + 1; i < l; i++) {
      if (isCapital != (b[i] < 97)) return false;
    }

    return true;
  }
}
