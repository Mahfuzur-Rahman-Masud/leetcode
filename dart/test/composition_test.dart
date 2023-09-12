import 'package:test/expect.dart';
import 'package:test/scaffolding.dart';

import '../bin/composition.dart';

void main(){
  group("easy", (){
    test("keyboard row", (){
      var s = Solution();
      expect(s.findWords(["Hello","Alaska","Dad","Peace"]),  ["Alaska","Dad"]);
      expect(s.findWords(["omk"]),  []);
    });
  });
}