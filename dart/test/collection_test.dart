import 'package:test/expect.dart';
import 'package:test/scaffolding.dart';

import '../bin/collections_impl.dart';

void main(){
  group("easy", (){
    test("smallest infinite set", (){
      var set = SmallestInfiniteSet();
      expect(1,set.popSmallest());
      expect(2,set.popSmallest());
      expect(3,set.popSmallest());
      set.addBack(2);
      expect(2,set.popSmallest());
      expect(4,set.popSmallest());
      set.addBack(3);
      set.addBack(2);
      set.addBack(2);
      expect(2,set.popSmallest());
      expect(3,set.popSmallest());
      expect(5,set.popSmallest());
      expect(6,set.popSmallest());
      expect(7,set.popSmallest());
    });
  });
}