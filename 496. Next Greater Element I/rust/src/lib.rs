#![allow(dead_code)]

use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let mut res = Vec::with_capacity(nums1.len());
        let mut num_greater_map = HashMap::new();
        let mut monotonic_stack: Vec<i32> = Vec::with_capacity(nums2.len());

        for &n in nums2.iter() {
            while let Some(&top) = monotonic_stack.last()
                && top < n
            {
                monotonic_stack.pop();
                num_greater_map.insert(top, n);
            }
            monotonic_stack.push(n);
        }

        for n in nums1 {
            res.push(num_greater_map.get(&n).copied().unwrap_or(-1));
        }

        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_next_greater_element() {
        let test_cases = vec![
            (vec![4, 1, 2], vec![1, 3, 4, 2], vec![-1, 3, -1]),
            (vec![2, 4], vec![1, 2, 3, 4], vec![3, -1]),
        ];

        for (nums1, nums2, expected) in test_cases {
            assert_eq!(Solution::next_greater_element(nums1, nums2), expected);
        }
    }
}
