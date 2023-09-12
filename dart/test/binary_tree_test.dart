import 'package:test/expect.dart';
import 'package:test/scaffolding.dart';

import '../bin/binary_tree.dart';

void main() {
  group("easy", () {
    test("bst get abs min diff", () {
      var s = Solution();
      TreeNode n = TreeNode(236,
        TreeNode(104),
        TreeNode(701,
        TreeNode(227, TreeNode(911))
        )
      );

      expect(s.getMinimumDifference(n), 9);
    });
  });
}