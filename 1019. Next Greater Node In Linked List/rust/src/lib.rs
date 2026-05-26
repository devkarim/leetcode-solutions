#![allow(dead_code)]

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution {}

impl Solution {
    pub fn next_larger_nodes(head: Option<Box<ListNode>>) -> Vec<i32> {
        let mut res = Vec::new();
        let mut curr = head;
        let mut st: Vec<[i32; 2]> = Vec::new();
        let mut counter = 0;
        while let Some(node) = curr {
            while let Some(&idx_num) = st.last()
                && idx_num[1] < node.val
            {
                res[idx_num[0] as usize] = node.val;
                st.pop();
            }
            st.push([counter, node.val]);
            res.push(0);
            curr = node.next;
            counter += 1;
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn to_list(vals: &[i32]) -> Option<Box<ListNode>> {
        vals.iter()
            .rev()
            .fold(None, |next, &val| Some(Box::new(ListNode { val, next })))
    }

    #[test]
    fn test_next_larger_nodes() {
        assert_eq!(
            Solution::next_larger_nodes(to_list(&[2, 1, 5])),
            vec![5, 5, 0]
        );
    }
}
