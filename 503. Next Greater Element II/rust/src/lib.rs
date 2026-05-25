#![allow(dead_code)]

struct Solution {}

impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let len = nums.len();
        let mut res: Vec<i32> = vec![-1; len];

        for (i, &n) in nums.iter().enumerate() {
            for j in 1..len {
                let n2 = nums[(i + j) % len];
                if n < n2 {
                    res[i] = n2;
                    break;
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
    fn test_next_greater_elements() {
        let test_cases = vec![
            (vec![1, 2, 1], vec![2, -1, 2]),
            (vec![1, 2, 3, 4, 3], vec![2, 3, 4, -1, 4]),
            (vec![1, 1], vec![-1, -1]),
            (vec![1], vec![-1]),
        ];

        for (nums, expected) in test_cases {
            assert_eq!(Solution::next_greater_elements(nums), expected);
        }
    }
}
