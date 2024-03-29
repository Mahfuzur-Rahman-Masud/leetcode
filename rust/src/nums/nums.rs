use std::process::Output;

struct Solution {}

impl Solution {
    pub fn self_dividing_numbers(left: i32, right: i32) -> Vec<i32> {
        let mut out = vec![];

        for n in left..=right {
            if n == 0 || n % 10 == 0 {
                continue;
            }

            let mut b = n;
            let mut sd = true;
            while b > 0 {
                if b % 10 == 0 || n % (b % 10) != 0 {
                    sd = false;
                    break;
                }

                b /= 10;
            }

            if sd {
                out.push(n)
            }
        }

        return out;
    }
}

#[cfg(test)]
mod test {
    use crate::nums::nums::Solution;

    #[test]
    fn test_x() {
        assert_eq!(vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22],Solution::self_dividing_numbers(1, 22));
        assert_eq!(vec![48,55,66,77],Solution::self_dividing_numbers(47, 85));
    }
}
