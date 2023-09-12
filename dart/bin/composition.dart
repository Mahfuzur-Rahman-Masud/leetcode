class Solution {
  List<String> findWords(List<String> words) {
    var m = <int, int>{};
    m[113] = 0;
    m[119] = 0;
    m[101] = 0;
    m[114] = 0;
    m[116] = 0;
    m[121] = 0;
    m[117] = 0;
    m[105] = 0;
    m[111] = 0;
    m[112] = 0;

    m[97] = 1;
    m[115] = 1;
    m[100] = 1;
    m[102] = 1;
    m[103] = 1;
    m[104] = 1;
    m[106] = 1;
    m[107] = 1;
    m[108] = 1;


    print('');
    print('');
    print('');
    m.forEach((k,v)=>
    print("m.insert($k, $v);"));



    words.removeWhere((w) {
      int? prev;
      for (var value in w.runes) {
        value = value< 97? value+32:value;
        var v = m[value] ?? 2;
        prev ??= v;
        if (v != prev) {
          return true;
        }
      }
      return false;
    });

    return words;
  }
}
