
pub fn max_count(m: i32, n: i32, ops: Vec<Vec<i32>>) -> i32 {
    let mut m = m;
    let mut n = n;
    for v in ops {
        if v[0] < m {
            m = v[0];
        }

        if v[1] < n {
            n = v[1];
        }
    }

    m * n
}

#[cfg(test)]
mod test {
    use crate::matrix::max_count;

    #[test]
    fn max_count_test() {
        assert_eq!(4, max_count(3, 3, vec![vec![2, 2], vec![3, 3]]));
        assert_eq!(9, max_count(3, 3, vec![]))
    }
}
