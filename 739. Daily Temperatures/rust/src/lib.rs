#![allow(dead_code)]

struct Solution {}

impl Solution {
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut res: Vec<i32> = vec![0; temperatures.len()];
        let mut monotonic_stack: Vec<usize> = vec![];
        for (i, &t) in temperatures.iter().enumerate() {
            while let Some(&j) = monotonic_stack.last()
                && temperatures[j] < t
            {
                res[j] = (i - j) as i32;
                monotonic_stack.pop();
            }
            monotonic_stack.push(i);
        }
        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_daily_temp() {
        let test_cases = vec![
            (
                vec![73, 74, 75, 71, 69, 72, 76, 73],
                vec![1, 1, 4, 2, 1, 1, 0, 0],
            ),
            (vec![30, 40, 50, 60], vec![1, 1, 1, 0]),
            (vec![30, 60, 90], vec![1, 1, 0]),
        ];
        for (temps, expected) in test_cases {
            assert_eq!(Solution::daily_temperatures(temps), expected);
        }
    }
}
