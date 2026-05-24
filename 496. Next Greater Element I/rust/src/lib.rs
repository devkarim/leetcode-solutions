#![allow(dead_code)]

use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let mut res = vec![-1; nums1.len()];
        let mut num_idx_map = HashMap::new();

        for (idx, n) in nums1.iter().enumerate() {
            num_idx_map.insert(n, idx);
        }

        for (i, &n1) in nums2.iter().enumerate() {
            if let Some(&j) = num_idx_map.get(&n1) {
                for (_, &n2) in nums2.iter().enumerate().skip(i + 1) {
                    if n2 > n1 {
                        res[j] = n2;
                        break;
                    }
                }
            }
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
