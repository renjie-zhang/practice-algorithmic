import java.util.List;
import java.util.ArrayList;

/*
 * @Descripttion: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 * @version:  链表的中序遍历
 * @LastEditTime: 2019-10-23 19:21:25
 * @Date: 2019-10-22 17:20:05
 * @LastEditTime: 2019-10-22 17:28:36
 */
class traversal{
    public static void main(String[] args) {
        TreeNode tree = new TreeNode(1);
        TreeNode tree2 = new TreeNode(2);
        TreeNode tree3 = new TreeNode(3);
        tree.right = tree2;
        tree2.left = tree3;
        List<Integer> a = inorderTraversal(tree);
        System.out.print(a.toString());
    }
    //使用递归
   public static List <Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<Integer>();
        helper(root, res);
        return res;
    }

    public static void helper(TreeNode root, List<Integer> res) {
        if (root != null) {
            if (root.left != null) {
                helper(root.left, res);
            }
            res.add(root.val);
            if (root.right != null) {
                helper(root.right, res);
            }
        }
    }
}
class TreeNode {
      int val;
      TreeNode left;
      TreeNode right;
      TreeNode(int x) { val = x; }
}