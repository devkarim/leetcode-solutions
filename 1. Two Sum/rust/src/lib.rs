#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for (i, a) in nums.iter().enumerate() {
            for (j, b) in nums.iter().enumerate().skip(i + 1) {
                if a + b == target {
                    return vec![i as i32, j as i32];
                }
            }
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

        for (nums, target, expected) in test_cases {
            assert_eq!(expected, Solution::two_sum(nums, target));
        }
    }
}
