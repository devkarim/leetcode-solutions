#![allow(dead_code)]

use std::collections::HashMap;
struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut num_index_map = HashMap::new();
        for (i, &n) in nums.iter().enumerate() {
            let required = target - n;
            if let Some(&j) = num_index_map.get(&required) {
                return vec![i as i32, j as i32];
            }
            num_index_map.insert(n, i);
        }
        unreachable!();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        let test_cases = vec![
            (vec![2, 7, 11, 15], 9, vec![0, 1]),
            (vec![3, 2, 4], 6, vec![1, 2]),
            (vec![3, 3], 6, vec![0, 1]),
        ];

        for (nums, target, mut expected) in test_cases {
            assert_eq!(expected.sort(), Solution::two_sum(nums, target).sort());
        }
    }
}
