package org.masud;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Queue;
import java.util.Set;

/**
 * @author Mahfuzur Rahman
 */
public class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;

  TreeNode() {
  }

  TreeNode(int val) {
    this.val = val;
  }

  TreeNode(int val, TreeNode left, TreeNode right) {
    this.val = val;
    this.left = left;
    this.right = right;
  }

  private static class TN extends TreeNode {

  }

  public static TreeNode from(Integer[] val) {
    TreeNode t = new TreeNode(val[0]);

    List<TreeNode> nodeList = new ArrayList<>();
    nodeList.add(t);
    int mark = 0;
    boolean left = true;

    for (int i = 1; i < val.length; i++) {
      Integer v = val[i];
      TreeNode tree = nodeList.get(mark);

      if (left && tree.left == null && v != null) {
        tree.left = new TreeNode(v);
        left = false;
        nodeList.add(tree.left);
      } else if (left && v == null) {
        left = false;
      } else if (!left && v != null && tree.right == null) {
        tree.right = new TreeNode(v);
        left = true;
        nodeList.add(tree.right);
        mark++;
      } else if (!left) {
        left = true;
        mark++;
      }
    }

    return t;
  }

  public void print() {
    Queue<TreeNode> q = new ArrayDeque<>();
    q.add(this);
    List<String> out = new ArrayList<>();

    while (!q.isEmpty()) {
      TreeNode n = q.poll();
      if (!(n instanceof TN)) {

        q.add(n.left == null ? new TN() : n.left);
        q.add(n.right == null ? new TN() : n.right);
      }

      out.add(n instanceof TN ? "[N]" : n.val + "");
    }

    while (!out.isEmpty() && out.get(out.size() - 1).equals("[N]")) {
      out.remove(out.size() - 1);
    }
    System.out.println(out);
  }




  @Override
  public String toString() {
    return val + ", " + left + ", " + right;
  }


  public TreeNode add(int val) {
    if (left == null && right == null) {
      TreeNode next = new TreeNode(val);
      if (val < this.val) {
        left = next;
      } else {
        right = next;
      }
      return next;
    } else if (left == null) {
      left = new TreeNode(val);
      return left;
    } else if (right == null) {
      right = new TreeNode(val);
      return right;
    } else if (val < this.val) {
      return left.add(val);
    } else {
      return right.add(val);
    }
  }


  List<String> paths = new ArrayList<>();

  public void buildPath(TreeNode node, String pfx) {
    pfx += node.val;
    if (node.left == null && node.right == null) {
      paths.add(pfx);
      return;
    }

    pfx += "->";

    if (node.left != null) {
      buildPath(node.left, pfx);
    }

    if (node.right != null) {
      buildPath(node.right, pfx);
    }
  }

  public List<String> binaryTreePaths(TreeNode root) {
    buildPath(root, "");
    return paths;
  }


  public TreeNode invertTree(TreeNode root) {
    if (root == null) return null;

    TreeNode temp = root.left;
    root.left = root.right;
    root.right = temp;
    invertTree(root.left);
    invertTree(root.right);
    return root;
  }


  public int countNodes(TreeNode root) {
    return root == null ? 0 : 1 + countNodes(root.left) + countNodes(root.right);
  }


  public ListNode getIntersectionNodeOnePass(ListNode headA, ListNode headB) {
    ListNode n1 = headA;
    ListNode n2 = headB;


    while (n1 != n2) {
      n1 = n1 == null ? headB : n1.next;
      n2 = n2 == null ? headA : n2.next;
    }

    return n1;
  }

  public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
    Set<ListNode> nodes = new HashSet<>();
    ListNode current = headA;
    while (current != null) {
      nodes.add(current);
      current = current.next;
    }

    current = headB;
    while (current != null) {
      if (nodes.contains(current)) {
        return current;
      }
      current = current.next;
    }

    return null;
  }

  public List<Integer> postorderTraversal(TreeNode root) {
    ArrayList<Integer> out = new ArrayList<>();
    postOrderTraversal(root, out);
    return out;
  }

  private void postOrderTraversal(TreeNode root, ArrayList<Integer> out) {
    if (root == null) return;
    postOrderTraversal(root.left, out);
    postOrderTraversal(root.right, out);
    out.add(root.val);
  }

  public List<Integer> preorderTraversal(TreeNode root) {
    ArrayList<Integer> out = new ArrayList<>();
    preOrderTraversal(root, out);
    return out;
  }

  private void preOrderTraversal(TreeNode node, List<Integer> buffer) {
    if (node == null) return;
    buffer.add(node.val);
    preOrderTraversal(node.left, buffer);
    preOrderTraversal(node.right, buffer);
  }

  public boolean isBalanced(TreeNode node) {
    return node == null || balancedHeight(node) != -1;
  }

  private int balancedHeight(TreeNode node) {
    if (node == null) return 0;
    int l = balancedHeight(node.left);
    int r = balancedHeight(node.right);
    if (l == -1 || r == -1) return -1;
    if (Math.abs(l - r) > 1) return -1;

    return 1 + Math.max(balancedHeight(node.left), balancedHeight(node.right));
  }

  public int maxDepth(TreeNode root) {
    return depth(root);
  }

  private int depth(TreeNode node) {
    return node == null ? 0 : 1 + Math.max(depth(node.left), depth(node.right));
  }


  public boolean hasPathSum(TreeNode root, int targetSum) {
    return root != null && hasPathSum2(root, targetSum);
  }

  public boolean hasPathSum2(TreeNode root, int targetSum) {
    if (root == null && targetSum != 0) return false;
    if (root == null) return true;
    if (root.val == targetSum && root.left == null && root.right == null) return true;


    return hasPathSum(root.left, targetSum - root.val)
      || hasPathSum(root.right, targetSum - root.val);

  }

  public int minDepth(TreeNode root) {
    if (root == null) return 0;
    int l = minDepth(root.left);
    int r = minDepth(root.right);
    return l == 0 ? 1 + r : r == 0 ? 1 + l : 1 + Math.min(l, r);
  }


  public static boolean isSymmetric(TreeNode root) {
    return root == null || isMirror(root.left, root.right);
  }

  private static boolean isMirror(TreeNode a, TreeNode b) {
    if (a == null && b == null) return true;
    if (a == null || b == null) return false;
    if (a.val != b.val) return false;
    return isMirror(a.left, b.right) && isMirror(a.right, b.left);
  }


  public static TreeNode sortedArrayToBST(int[] nums) {
    return createTree(nums, 0, nums.length - 1);
  }

  public static TreeNode createTree(int[] nums, int l, int r) {
    if (l > r) return null;

    int mid = l + (r - l) / 2;

    TreeNode root = new TreeNode(nums[mid]);
    root.left = createTree(nums, l, mid - 1);
    root.right = createTree(nums, mid + 1, r);
    return root;
  }

  public static void add(TreeNode node, int val) {
    TreeNode next = new TreeNode(val);
    if (val < node.val && node.left == null) {
      node.left = next;

    } else if (val > node.val && node.right == null) {
      node.right = next;
    } else if (val < node.val && val < node.left.val && node.left.left == null && node.left.right == null) {
      next.right = node.left;
      node.left = next;
    } else if (val > node.val && val > node.right.val && node.right.left == null && node.right.right == null) {
      next.left = node.right;
      node.right = next;
    } else if (val < node.val) {
      add(node.left, val);
    } else {
      add(node.right, val);
    }
  }
}