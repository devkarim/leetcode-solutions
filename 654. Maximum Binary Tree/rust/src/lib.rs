#![allow(dead_code)]
// Definition for a binary tree node.
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

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn construct_maximum_binary_tree(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        let mut st: Vec<Rc<RefCell<TreeNode>>> = Vec::new();
        for num in nums {
            let mut left = None;
            while let Some(last) = st.last()
                && last.borrow().val < num
            {
                left = st.pop();
            }

            let tree = Rc::new(RefCell::new(TreeNode::new(num)));
            tree.borrow_mut().left = left;

            if let Some(last) = st.last() {
                last.borrow_mut().right = Some(Rc::clone(&tree));
            }

            st.push(tree);
        }
        Some(st.remove(0))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn node(
        val: i32,
        left: Option<Rc<RefCell<TreeNode>>>,
        right: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        let mut n = TreeNode::new(val);
        n.left = left;
        n.right = right;
        Some(Rc::new(RefCell::new(n)))
    }

    #[test]
    fn example1() {
        let expected = node(
            6,
            node(3, None, node(2, None, node(1, None, None))),
            node(5, node(0, None, None), None),
        );
        assert_eq!(
            Solution::construct_maximum_binary_tree(vec![3, 2, 1, 6, 0, 5]),
            expected
        );
    }
}
