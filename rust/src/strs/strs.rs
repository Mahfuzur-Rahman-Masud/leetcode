struct Solution{}

impl Solution {
    pub fn num_jewels_in_stones(jewels: String, stones: String) -> i32 {
        let mut js = [0;123];

        jewels.bytes().into_iter().map(|b|b as usize)
        .for_each(|b|js[b]=1);


        stones.bytes().into_iter()
        .map(|b|js[b as usize])
        .sum()
    }
}



impl Solution {
    pub fn rotate_string(s: String, goal: String) -> bool {
        if s.len() != goal.len(){
            return false;
        }

        let l = s.len();
        let g = goal.into_bytes();
        let b = s.into_bytes();


        for i in 0..l{
            let mut m = true;
            for j in 0..l{
                let xx = j+i;
                let x = match (xx).cmp(&l) {
                    std::cmp::Ordering::Less => xx,
                    std::cmp::Ordering::Equal => xx-l,
                    std::cmp::Ordering::Greater => xx-l,
                };

                if b[x] != g[j]{
                    m = false;
                    break;
                }
            }
            
            if m == true{
                return  true;
            }
        }

     false   
    }
}



#[cfg(test)]
mod test{
    use crate::strs::strs::Solution;

    #[test]
    fn num_jewesls_in_stones_test(){
        assert_eq!(3, Solution::num_jewels_in_stones("aA".to_owned(), "aAAbbbb".to_owned()));
    }


    #[test]
    fn roatate_string_test(){
        assert!(Solution::rotate_string("abcde".to_owned(), "cdeab".to_owned()));
        assert!(Solution::rotate_string("a".to_owned(), "a".to_owned()));
        assert!(Solution::rotate_string("ab".to_owned(), "ab".to_owned()));
        assert!(!Solution::rotate_string("abcde".to_owned(), "abced".to_owned()));
    }
}