use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}


pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
    match &root {
        None => {}
        Some(ref node) => {
            let mut n = node.borrow_mut();
            let left = invert_tree(n.right.take());
            let right = invert_tree(n.left.take());
            n.left = left;
            n.right = right;
        }
    };

    root
}


pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    match root {
        None => 0,
        Some(node) => {
            let mut n = node.borrow_mut();
            1 + count_nodes(n.left.take()) + count_nodes(n.right.take())
        }
    }
}

#[allow(dead_code)]
struct Solution();

#[allow(dead_code)]
impl Solution {
    pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        match root {
            None => 0,
            Some(n) => {
                let mut n = n.borrow_mut();
                1 + Self::count_nodes(n.left.take()) + Self::count_nodes(n.right.take())
            }
        }
    }
}

#[allow(dead_code)]
pub fn binary_tree_paths(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<String> {
    fn build(node: &Option<Rc<RefCell<TreeNode>>>, path: &str, v: &mut Vec<String>) {
        if let Some(n) = node {
            let l = &n.borrow().left;
            let r = &n.borrow().right;
            if l.is_none() && r.is_none() {
                v.push(format!("{path}{}", n.borrow().val))
            } else {
                let next_path = format!("{path}{}->", n.borrow().val);
                build(l, &next_path, v);
                build(r, &next_path, v);
            }
        }
    }

    let mut paths = Vec::new();
    build(&root, "", &mut paths);
    paths
}

#[allow(dead_code)]
pub fn binary_tree_paths_iter_1(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<String> {
    let mut v = vec![];
    binary_tree_paths_build(root, &mut v, "".to_string());
    return v;
}

#[allow(dead_code)]
pub fn binary_tree_paths_build(node: Option<Rc<RefCell<TreeNode>>>, v: &mut Vec<String>, p: String) {
    let mut p2 = format!("{}", node.as_ref().unwrap().borrow().val);
    if p != "" {
        p2 = format!("{}->{}", p, p2);
    }

    if node.as_ref().unwrap().borrow().left.is_none() && node.as_ref().unwrap().borrow().right.is_none() {
        v.push(p2.clone());
        return;
    }

    if node.as_ref().unwrap().borrow().left.is_some() {
        binary_tree_paths_build(node.as_ref().unwrap().borrow().left.clone(), v, p2.clone());
    }

    if node.as_ref().unwrap().borrow().right.is_some() {
        binary_tree_paths_build(node.as_ref().unwrap().borrow().right.clone(), v, p2.clone());
    }
}


#[allow(dead_code)]
pub fn diameter_of_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    use std::cmp::max;
    fn max_path_and_v(root: &Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
        match root {
            None => (0, 0),
            Some(n) => {
                let l = max_path_and_v(&n.borrow().left);
                let r = max_path_and_v(&n.borrow().right);
                (max(l.0, r.0)+1, max(l.0+r.0, max(l.1,r.1)))
            }
        }
    }
    max_path_and_v(&root).1
}

#[allow(dead_code)]
pub fn diameter_of_binary_tree2(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    use std::cmp::max;
    fn max_path_and_v(root: Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
        match root {
            None => (0, 0),
            Some(n) => {
                let l = max_path_and_v(n.borrow().left.clone());
                let r = max_path_and_v(n.borrow().right.clone());
                (max(l.0, r.0)+1, max(l.0+r.0, max(l.1,r.1)))
            }
        }
    }
    max_path_and_v(root).1
}




#[cfg(test)]
mod test {
    use std::{assert_eq, println};
    use std::cell::{Cell, RefCell};
    use std::rc::Rc;

    use crate::b_tree::{binary_tree_paths, binary_tree_paths_iter_1, count_nodes, diameter_of_binary_tree, invert_tree, TreeNode};

    #[test]
    fn ref_cell_test(){
        let x = Rc::new(RefCell::new(5));
        println!("{:?}", x);
        let y = x.clone();
        println!("{:?}", y);
        // add 1 to the value of x
        {
            let mut x_mut = x.borrow_mut();
            *x_mut += 1;
        }

        // x.set(Rc::new(RefCell::new(5)));
        let z = Rc::new(Cell::new(5));
        z.set(4);
        println!("z:{:?}", z);

        // values after modifying y , but both x and y are
        println!("{:?}", x);
        println!("{:?}", y);

    }

    #[test]
    fn diameter_of_binary_tree_test() {
        assert_eq!(0,diameter_of_binary_tree(None) );
        assert_eq!(3,diameter_of_binary_tree(Some(Rc::new(RefCell::new(TreeNode{
            val:1,
            right: Some(Rc::new(RefCell::new(TreeNode{val:3, left:None, right:None}))),
            left: Some(Rc::new(RefCell::new(TreeNode{
            val:2,
            right: Some(Rc::new(RefCell::new(TreeNode{val:5, left:None, right:None}))),
            left: Some(Rc::new(RefCell::new(TreeNode{val:4, left:None, right:None}))),
        })))
        })))) );

        assert_eq!(1,diameter_of_binary_tree(Some(Rc::new(RefCell::new(TreeNode{
            val:1,
            right: None,
            left: Some(Rc::new(RefCell::new(TreeNode{
                val:2,
                right: None,
                left: None,
            })))
        })))) );


    }

    #[test]
    fn binary_tree_paths_test() {
        let t = TreeNode {
            val: 1,
            left: Some(Rc::new(RefCell::new(TreeNode {
                val: 2,
                left: None,
                right: Some(Rc::new(RefCell::new(TreeNode { val: 5, left: None, right: None }))),
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode { val: 3, left: None, right: None }))),
        };

        let p = binary_tree_paths_iter_1(
            Some(Rc::new(RefCell::new(TreeNode::new(1)))));
        println!("{:?}", p);
        println!("{:?}", binary_tree_paths(Some(Rc::new(RefCell::new(t)))));
    }

    #[test]
    fn invert_tree_test() {
        let node = TreeNode::new(10);
        invert_tree(Some(Rc::new(RefCell::new(node))));
    }

    #[test]
    fn count_nodes_test() {
        assert_eq!(count_nodes(Some(Rc::new(RefCell::new(TreeNode::new(1))))), 1)
    }
}