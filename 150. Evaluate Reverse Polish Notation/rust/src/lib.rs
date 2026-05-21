#![allow(dead_code)]
struct Solution {}

impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack: Vec<i32> = Vec::new();
        for token in tokens {
            let res = match token.as_str() {
                "+" | "-" | "*" | "/" => {
                    let b = stack.pop().unwrap();
                    let a = stack.pop().unwrap();
                    match token.as_str() {
                        "+" => a + b,
                        "-" => a - b,
                        "/" => a / b,
                        "*" => a * b,
                        _ => unreachable!(),
                    }
                }

                _ => token.parse().unwrap(),
            };
            stack.push(res);
        }
        stack.pop().unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn tokens(v: &[&str]) -> Vec<String> {
        v.iter().map(|s| s.to_string()).collect()
    }

    #[test]
    fn test_eval_rpn() {
        let test_cases = vec![
            (tokens(&["2", "1", "+", "3", "*"]), 9),
            (tokens(&["4", "13", "5", "/", "+"]), 6),
        ];

        for (input, answer) in test_cases {
            assert_eq!(Solution::eval_rpn(input), answer);
        }
    }
}
