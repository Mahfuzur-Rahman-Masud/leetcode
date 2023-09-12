
import 'package:test/expect.dart';
import 'package:test/scaffolding.dart';

import '../bin/strings.dart';


void main() {
  group('easy', (){
    test('substring', (){
      var s = Solution();
     expect(s.findLUSlength("aba", "cdc"), 3 );
     expect(s.findLUSlength("aaa", "bbb"), 3 );
     expect(s.findLUSlength("aaa", "aaa"), -1 );
     expect(s.findLUSlength("aaab", "aaa"), 4 );
     expect(s.findLUSlength("aabna", "aanaabbna"), 9 );


      expect(s.findLUSlength2("aba", "cdc"), 3 );
      expect(s.findLUSlength2("aaa", "bbb"), 3 );
      expect(s.findLUSlength2("aaa", "aaa"), -1 );
      expect(s.findLUSlength2("aaab", "aaa"), 4 );
      expect(s.findLUSlength2("aabna", "aanaabbna"), 9 );
    });


    test('detectCapitalUse', () {
      var s = Solution();
      expect(s.detectCapitalUse(""), true);
      expect(s.detectCapitalUse("a"), true);
      expect(s.detectCapitalUse("A"), true);
      expect(s.detectCapitalUse("Ab"), true);
      expect(s.detectCapitalUse("AA"), true);
      expect(s.detectCapitalUse("aa"), true);
      expect(!s.detectCapitalUse("aA"), true);
      expect(s.detectCapitalUse("abba"), true);
      expect(s.detectCapitalUse("ABBA"), true);
      expect(!s.detectCapitalUse("aXst"), true);
      expect(!s.detectCapitalUse("abbA"), true);
      expect(!s.detectCapitalUse("ABBa"), true);
      expect(s.detectCapitalUse("Abba"), true);
    });
  });
}
