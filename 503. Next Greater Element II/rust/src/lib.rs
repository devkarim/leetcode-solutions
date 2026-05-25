#![allow(dead_code)]

struct Solution {}

impl Solution {
    pub fn next_greater_elements(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut res = vec![-1; n];
        let mut st = Vec::with_capacity(n);

        for i in 0..(2 * n) {
            let idx = i % n;
            let num = nums[idx];
            while let Some(&j) = st.last()
                && nums[j] < num
            {
                res[j] = num;
                st.pop();
            }
            if i < n {
                st.push(idx);
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
            (vec![5, 4, 3, 2, 1], vec![-1, 5, 5, 5, 5]),
        ];

        for (nums, expected) in test_cases {
            assert_eq!(Solution::next_greater_elements(nums), expected);
        }
    }
}
