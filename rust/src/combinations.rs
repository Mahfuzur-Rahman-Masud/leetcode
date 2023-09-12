#[allow(dead_code)]
struct Solution();

#[allow(dead_code, unused_variables)]
impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        if n == 0{
            return Vec::new();
        }

        fn back_track( left: i32, right: i32, s: String) -> Vec<String> {
            let mut out = Vec::new();

            if left == 0 && right == 0 {
                out.push(s);
                return out;
            }




            if left > 0 {
                let mut s = s.clone();
                s.push_str("(");
                out.append(&mut back_track( left -1 , right +1 , s));
            }

            if right > 0 {
                let mut s = s;
                s.push_str(")");
                out.append(&mut back_track( left, right - 1, s));
            }

            out
        }

        back_track(n, 0, "".to_string())
    }
}

#[cfg(test)]
mod test {
    use crate::combinations::Solution;

    #[test]
    fn generate_parenthesis_test() {
        // println!("{:?}", Solution::generate_parenthesis(1));

        assert_eq!(Vec::<String>::new(), Solution::generate_parenthesis(0));
        assert_eq!(vec!["()".to_string()], Solution::generate_parenthesis(1));
        assert_eq!(vec!["(())", "()()"].into_iter()
                       .map(|x|x.to_string())
                       .collect::<Vec<String>>(), Solution::generate_parenthesis(2));

        assert_eq!(vec!["((()))", "(()())","(())()","()(())", "()()()"].into_iter()
                       .map(|x|x.to_string())
                       .collect::<Vec<String>>(), Solution::generate_parenthesis(3));


        assert_eq!(vec!["(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"].into_iter()
                       .map(|x|x.to_string())
                       .collect::<Vec<String>>(), Solution::generate_parenthesis(4));
    }
}