package org.masud;

import java.util.ArrayDeque;
import org.junit.jupiter.api.Test;

/**
 * @author Mahfuzur Rahman
 */
class TreeNodeTest extends TreeNode {


  @Test
  void hasPathSomeTest() {
    TreeNode root = new TreeNode(5);
    TreeNode eleven = root.add(4)
      .add(11);
    eleven.add(7);
    eleven.add(2);

    boolean b = hasPathSum(root, 22);
    System.out.printf("Ans: %s | expected %s%n", b, true);

  } @Test
  void hasPathSomeTest2() {
    TreeNode root = new TreeNode(-2);
    TreeNode eleven = root.add(-3);

    boolean b = hasPathSum(root, -5);
    System.out.printf("Ans: %s | expected %s%n", b, true);

  }

  @Test
  void buildTest() {
    int[] ints = {0, 1, 2, 3, 4, 5};
    TreeNode node = sortedArrayToBST(ints);

    ArrayDeque<TreeNode> queue = new ArrayDeque<>();
    queue.add(node);

    while (!queue.isEmpty()) {
      TreeNode poll = queue.poll();

      System.out.printf("%s, ", poll.val);
      if (poll.left != null) {
        queue.add(poll.left);
      }

      if (poll.right != null) {
        queue.add(poll.right);
      }
    }


//    expected : [3,1,5,0,2,4]
  }




  @Test
  void preOrderTraversal() {
    TreeNode node = new TreeNode(3, new TreeNode(9), new TreeNode(20, new TreeNode(15), new TreeNode(7)));
//    System.out.println(node);
    assert node.left.val ==9;
    assert node.right.val ==20;
    assert node.right.left.val ==15;
    assert node.right.right.val ==7;
    node.print();
    TreeNode.from(new Integer[]{3,9,20, null, null, 15, 7}).print();
  }
}