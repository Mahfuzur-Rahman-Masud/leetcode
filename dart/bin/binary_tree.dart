import 'dart:math';

class TreeNode {
  int val;
  TreeNode? left;
  TreeNode? right;

  TreeNode([this.val = 0, this.left, this.right]);
}

class Solution {
  int getMinimumDifference(TreeNode? root) {
    var list = <int>[];
     minDiff(root, list);
     list.sort((a,b)=>b-a);

     var min = 10000000;
     for (var i = 1; i< list.length; i++){
       var m = list[i-1] - list[i];
       if (m < min){
         min = m;
       }
    }

     return min;
  }

   minDiff(TreeNode? root, List<int> list) {
    if (root == null) return;
    list.add(root.val);
    minDiff(root.left, list);
    minDiff(root.right, list);
  }
}
