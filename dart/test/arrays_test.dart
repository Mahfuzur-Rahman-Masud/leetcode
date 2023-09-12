import 'package:test/expect.dart';
import 'package:test/scaffolding.dart';

import '../bin/arrays.dart';

void main(){
  group("easy", (){
      test("consecutive ones", (){
        expect (Solution().findMaxConsecutiveOnes([0, 1, 0, 1, 1,0,0,0,1]), 2);
      });

      test("last stone weight", (){
        expect(Solution().lastStoneWeight([2,7,4,1,8,1]), 1);
        expect(Solution().lastStoneWeight([1]), 1);
        expect(Solution().lastStoneWeight([]), 0);
        expect(Solution().lastStoneWeight([2,7]), 5);
      });


      test("next greatest", (){
        expect(Solution().nextGreaterElement([4,1,2], [1,3,4,2]), [-1,3,-1]);
        expect(Solution().nextGreaterElement([2,4], [1,3,4,2]), [-1,-1]);
        expect(Solution().nextGreaterElement([2,4], [1,2,3,4]), [3,-1]);
      });
  });
}