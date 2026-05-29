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
        return Self::build_tree(&nums, 0, nums.len() as i32);
    }

    pub fn build_tree(nums: &Vec<i32>, start: i32, end: i32) -> Option<Rc<RefCell<TreeNode>>> {
        let max_item = nums
            .iter()
            .skip(start as usize)
            .take((end - start) as usize)
            .enumerate()
            .max_by_key(|&(_index, &value)| value);
        let (idx, &max_val) = max_item.unwrap();
        let idx = idx as i32 + start as i32;

        let mut tree = TreeNode::new(max_val);
        if start < idx {
            tree.left = Self::build_tree(&nums, start, idx);
        }
        if idx + 1 < end {
            tree.right = Self::build_tree(&nums, idx + 1, end);
        }
        return Some(Rc::new(RefCell::new(tree)));
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
