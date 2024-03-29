use core::num;
use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn max_subarray_length(nums: Vec<i32>, k: i32) -> i32 {
        let mut start = 0;
        let mut cur = 0;

        let mut mx = 0;
        let mut v;
        let mut c;
        let l = nums.len();

        let mut m: HashMap<i32, i32> = HashMap::new();

        while cur < l {
            v = nums[cur];
            c = match m.get(&v) {
                Some(x) => *x,
                None => 0,
            };

            m.insert(v, c + 1);

            if c == k {
                for i in start..cur {
                    let v2 = nums[i];
                    let x = *m.get(&v2).unwrap();
                    m.insert(v2, x - 1);

                    if v2 == v {
                        start = i + 1;
                        if l - start < mx {
                            return mx as i32;
                        }

                        break;
                    }
                }
            } else {
                let ll = cur - start + 1;
                if ll > mx {
                    mx = ll;
                }
            }

            cur += 1;
        }

        return mx as i32;
    }
}

impl Solution {
    pub fn count_subarrays(nums: Vec<i32>, k: i32) -> i64 {
        let mut mx = 0;
        for &n in &nums {
            if n > mx {
                mx = n;
            }
        }




        let mut mark = 0;
        let mut count = 0;
        let mut out = 0;

        for &n in &nums{
            if n == mx{
                count+=1;
            }


            while count >= k{
                if nums[mark] == mx{
                    count-=1;
                }
                mark+=1;
            }

            out+=mark as i64;
        }


        out 
    }

    pub fn count_subarrays0(nums: Vec<i32>, k: i32) -> i64 {
        let mut mx = 0;
        for &n in nums.iter() {
            if n > mx {
                mx = n
            }
        }

        let ln = nums.len();
        let mut cur = 0;
        let mut out = 0;
        let mut trace: Vec<usize> = vec![];
        let k = k as usize;

        while cur < ln {
            let n = nums[cur];
            if n == mx {
                trace.push(cur)
            }

            if trace.len() >= k {
                out += (trace[trace.len() - k] + 1) as i64
            }

            cur += 1;
        }

        out
    }
}

#[cfg(test)]
mod test {
    use crate::sliding_window::sliding_window::Solution;


    #[test]
    fn count_subarrays_test() {
        assert_eq!(6_i64, Solution::count_subarrays(vec![1, 3, 2, 3, 3], 2));
        assert_eq!(0_i64, Solution::count_subarrays(vec![1, 4, 2, 1], 3));
        assert_eq!(1_i64, Solution::count_subarrays(vec![1], 1));
        assert_eq!(2_i64, Solution::count_subarrays(vec![1, 2], 1));
    }


    #[test]
    fn count_subarrays_test0() {
        assert_eq!(6_i64, Solution::count_subarrays0(vec![1, 3, 2, 3, 3], 2));
        assert_eq!(0_i64, Solution::count_subarrays0(vec![1, 4, 2, 1], 3));
        assert_eq!(1_i64, Solution::count_subarrays0(vec![1], 1));
        assert_eq!(2_i64, Solution::count_subarrays0(vec![1, 2], 1));
    }

    #[test]
    fn max_subarray_length_test() {
        assert_eq!(
            6,
            Solution::max_subarray_length(vec![1, 2, 3, 1, 2, 3, 1, 2], 2)
        );
        assert_eq!(1, Solution::max_subarray_length(vec![1], 2));
        assert_eq!(1, Solution::max_subarray_length(vec![1], 1));
        assert_eq!(
            2,
            Solution::max_subarray_length(vec![1, 2, 1, 2, 1, 2, 1, 2], 1)
        );
        assert_eq!(
            4,
            Solution::max_subarray_length(vec![5, 5, 5, 5, 5, 5, 5], 4)
        );
    }
}
