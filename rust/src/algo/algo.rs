

struct Solution {}

impl Solution {
    pub fn search0(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() == 0 {
            return -1;
        }

        let mut l:i32 = 0;
        let mut r:i32 = nums.len() as i32 -1;
        let mut c ;

        while l <= r {
            c = l + (r - l) / 2;
            let v = nums[c as usize];
            if v == target {
                return c as i32;
            }

            if v < target {
                l = c + 1;
            } else {
                r = c - 1;
            }
        }

        -1
    }
}


impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        if nums.len() == 0 {
            return -1;
        }

       let (mut l , mut r, mut i) = (0_usize, nums.len(), 0_usize);

       while l< r{
        i = (l+r)/2;
        match nums[i].cmp(&target) {
            std::cmp::Ordering::Less => l = i+1,
            std::cmp::Ordering::Equal => {return i as i32},
            std::cmp::Ordering::Greater => r = i,
        }
       }

        return -1
    }


}




#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn test_binary_search() {
        assert_eq!(4,  Solution::search(vec![-1,0,3,5,9,12], 9));
        assert_eq!(-1,  Solution::search(vec![-1,0,3,5,9,12], 2));
        assert_eq!(-1, Solution::search(vec![5], -5));
        assert_eq!(0, Solution::search(vec![5], 5));
    }
}
