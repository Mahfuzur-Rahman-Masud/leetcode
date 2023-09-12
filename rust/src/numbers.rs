use std::sync::Mutex;

#[allow(dead_code)]
pub fn is_power_of_two_4(n: i32) -> bool {
    n > 0 && (n | (n - 1) == (n << 1) - 1)
}

#[allow(dead_code)]
pub fn is_power_of_two_3(n: i32) -> bool {
    n > 0 && ((n << 1) - 1 == (n - 1) | n)
}

#[allow(dead_code)]
pub fn is_power_of_two_2(n: i32) -> bool {
    n > 0 && n.count_ones() == 1
}

#[allow(dead_code)]
pub fn is_power_of_two(n: i32) -> bool {
    if n <= 0 {
        return false;
    }

    if n == 1 {
        return true;
    }

    let mut n2 = 1;

    for _ in 0..32 {
        n2 <<= 1;

        if n == n2 {
            return true;
        }
    }

    false
}

#[allow(dead_code)]
pub fn is_happy(n: i32) -> bool {
    let mut fast = n;
    let mut slow = n;

    loop {
        slow = sum_sqr_digits(slow);
        fast = sum_sqr_digits(sum_sqr_digits(fast));

        if fast == 1 || slow == 1 {
            return true;
        }

        if fast == slow {
            return false;
        }
    }
}

#[allow(dead_code)]
fn sum_sqr_digits(n: i32) -> i32 {
    let mut n = n;
    let mut sum = 0;
    while n > 0 {
        let d = n % 10;
        n /= 10;
        sum += d * d;
    }

    sum
}


#[allow(dead_code)]
pub fn add_digits(mut num: i32) -> i32 {
    let mut next;
    while num > 9 {
        next = num;
        num = 0;
        while next > 0 {
            num += next % 10;
            next /= 10;
        }
    }
    num
}


#[allow(dead_code)]
pub fn is_ugly(n: i32) -> bool {
    if n < 1 {
        return false;
    }

    if n == 1 || n == 2 || n == 3 || n == 5 {
        return true;
    }

    let mut x = n;

    if x % 2 == 0 { x /= 2; }
    if x % 3 == 0 { x /= 3; }
    if x % 5 == 0 { x /= 5; }

    if x != n {
        return is_ugly(x);
    }

    false
}


#[allow(dead_code)]
pub fn missing_number(nums: Vec<i32>) -> i32 {
    let l = (nums.len() + 1) as i32;
    let l2: i32 = nums.iter().sum();
    l * (l - 1) / 2 - l2
}

#[allow(dead_code)]
pub fn is_power_of_three_1(mut n: i32) -> bool {
    if n < 1 { return false; }
    while n % 3 == 0 {
        n /= 3;
    }
    n == 1
}

#[allow(dead_code)]
pub fn is_power_of_three(n: i32) -> bool {
    let n: i64 = n as i64;
    if n < 1 { return false; }
    if n == 1 { return true; }
    let mut x = 1;
    while x < n {
        x *= 3;
    }
    x == n
}

#[allow(dead_code)]
pub fn is_power_of_four(mut n: i32) -> bool {
    if n < 1 {
        return false;
    }

    while n % 4 == 0 {
        n >>= 2;
    }

    n == 1
}


/**
Observations:
    even number bits(n) = bits(n/2); i.e bits(6) = bits(3)
    odd number bits(n) = bits(n-1) +1;

00 00000000 --> 00
01 00000001 --> 01
02 00000010 --> 01
03 00000011 --> 02
04 00000100 --> 01
05 00000101 --> 02
06 00000110 --> 02
07 00000111 --> 03
08 00001000 --> 01
09 00001001 --> 02
10 00001010 --> 02
11 00001011 --> 03
12 00001100 --> 02
13 00001101 --> 03
14 00001110 --> 03
15 00001111 --> 04
16 00010000 --> 01
17 00010001 --> 02
18 00010010 --> 02
19 00010011 --> 03
20 00010100 --> 02
 */
#[allow(dead_code)]
pub fn count_bits2(n: i32) -> Vec<i32> {
    let l = (n + 1) as usize;
    let mut out: Vec<i32> = vec![0; l];
    for i in 0..l {
        out[i] = if i % 2 == 0 {
            out[i / 2]
        } else {
            out[i - 1] + 1
        }
    }

    out
}

#[allow(dead_code)]
pub fn count_bits3(n: i32) -> Vec<i32> {
    let l = (n + 1) as usize;
    let mut out: Vec<i32> = vec![0; l];

    for x in 1..l {
        out[x] = if x & 1 == 1 {
            out[x - 1] + 1
        } else {
            out[x / 2]
        }
    }


    out
}


#[allow(dead_code)]
pub fn count_bits(n: i32) -> Vec<i32> {
    (0..=n).map(|x| x.count_ones() as i32).collect()
}

#[allow(dead_code)]
pub fn is_perfect_square(num: i32) -> bool {
    let target = num as i64;
    let mut l: i64 = 0;
    let mut h: i64 = num as i64;
    let mut mid;
    let mut r;

    while l < h {
        mid = l + (h - l) / 2;
        r = mid * mid;
        if r == target {
            return true;
        } else if r > target {
            h = mid - 1;
        } else {
            l = mid + 1;
        }
    }

    l * l == target
}

static SEED: Mutex<i32> = Mutex::new(0);

#[allow(dead_code)]
fn set_seed(num: i32) -> i32 {
    *SEED.lock().unwrap() = num;
    num
}

#[allow(dead_code)]
unsafe fn guess(num: i32) -> i32 {
    let r = *SEED.lock().unwrap();
    if num > r {
        return -1;
    } else if num < r {
        return 1;
    }

    0
}

#[allow(dead_code)]
unsafe fn guess_number(n: i32) -> i32 {
    let (mut l, mut h) = (0, n);
    while l < h {
        let m = l + (h - l) / 2;
        match guess(m) {
            -1 => h = m - 1,
            1 => l = m + 1,
            _ => return m,
        }
    }

    l
}


/**
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
 */
#[allow(dead_code)]
pub fn is_subsequence(s: String, t: String) -> bool {
    let mut i = 0;
    let a = s.as_bytes();
    let b = t.as_bytes();


    'outer: for x in a {
        while i < t.len() {
            if b[i] == *x {
                continue 'outer;
            } else {
                i += 1;
            }
        }

        return false;
    }

    true
}

#[allow(dead_code)]
pub fn is_subsequence2(s: String, t: String) -> bool {
    let mut itr = t.chars();
    s.chars()
        .all(|c| itr.find(|c2| c == *c2).is_some())
}

#[allow(dead_code)]
pub fn to_hex(num: i32) -> String {
    if num == 0 {
        return "0".to_string();
    }
    let mut nn = num as u32;
    let mut out = Vec::with_capacity(64);
    let code: [u8; 16] = [48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 101, 102];

    while nn > 0 {
        let n = nn % 16;
        out.push(code[n as usize]);
        nn /= 16;
    }

    out.reverse();

    unsafe { String::from_utf8_unchecked(out) }
}


#[allow(dead_code)]
fn fizz_buzz(n: i32) -> Vec<String> {
    let mut ans = Vec::with_capacity(n as usize);

    for x in 1..=n {
        let t = x % 3 == 0;
        let f = x % 5 == 0;

        if t && f {
            ans.push("FizzBuzz".to_string());
        } else if t { ans.push("Fizz".to_string()) } else if f { ans.push("Buzz".to_string()) } else { ans.push(x.to_string()) }
    }

    ans
}


#[allow(dead_code)]
pub fn arrange_coins(mut n: i32) -> i32 {
    let mut step = 0;
    while n > 0 {
        step += 1;
        n -= step;
    }


    let ans = if n == 0 {
        step
    } else {
        step - 1
    };

    ans
}

#[allow(dead_code)]
pub fn arrange_coins2(n: i32) -> i32 {
    let target = n as i64;
    let mut high = n as i64;
    let mut low = 0;
    let mut ans = 0;

    while low <= high {
        let mid = low + (high - low) / 2;
        let res = mid * (mid + 1) / 2;
        if res <= target {
            low = mid + 1;
            ans = mid;
        } else {
            high = mid - 1;
        }
    }

    ans as i32
}

#[allow(dead_code)]
pub fn reverse(x: i32) -> i32 {
    let mut ans = 0;
    let mut x = x;
    // ans *10 + d <= MAX
    // (MAX - d)/10 >= ans
    // overflow ans > (MAX -d ) /10
    // overflow ans < (MIN -d)/10

    if x > 0 {
        while x != 0 {
            let d = x % 10;
            if ans > (i32::MAX - d) / 10 {
                return 0;
            }
            x /= 10;
            ans = ans * 10 + d;
        }
    } else {
        while x != 0 {
            let d = x % 10;
            if ans < (i32::MIN - d) / 10 {
                return 0;
            }
            x /= 10;
            ans = ans * 10 + d;
        }
    }

    ans
}


#[allow(dead_code)]
pub fn my_atoi(s: String) -> i32 {
    let mut sign = 1;
    let mut ans = 0;
    let mut signed = false;


    for c in s.chars() {
        if signed && (c == '-' || c == '+' || c == ' ') {
            break;
        }
        if c == '-' {
            sign *= -1;
            signed = true;
            continue;
        }

        if c == '+' {
            signed = true;
            continue;
        }

        if c == ' ' {
            continue;
        }

        if c < '0' || c > '9' {
            break;
        }

        signed = true;

        let mut d = c as i32 - 48;
        d *= sign;

        // println!("{} {}-> {}", ans, d, ans as i64 * 10 + d as i64);


        if d >= 0 && ans > (i32::MAX - d) / 10 {
            return i32::MAX;
        }

        if d <= 0 && ans < (i32::MIN - d) / 10 {
            // ans *10+d >= MIN
            // ans *10 >=MIN-d
            // ans >= (MIN-d)/10
            return i32::MIN;
        }

        // println!("{} {}-> {}", ans, d, ans * 10 + d);
        ans = ans * 10 + d
    }

    ans
}


#[allow(dead_code)]
pub fn btn_phone_number_to_letter_combinations(digits: String) -> Vec<String> {
    if digits.len() == 0 {
        return Vec::new();
    }

    let keys: Vec<Vec<u8>> = vec![vec![32], vec![0], vec![97, 98, 99], vec![100, 101, 102], vec![103, 104, 105], vec![106, 107, 108], vec![109, 110, 111], vec![112, 113, 114, 115], vec![116, 117, 118], vec![119, 120, 121, 122]];
    let keys_pressed = digits.chars().map(|d| d as usize - 48).collect::<Vec<usize>>();
    let possible_combs = keys_pressed.iter().map(|x| keys[*x].len()).product();


    let mut combination_buffer = vec![vec![0_u8; digits.len()]; possible_combs];
    let mut spread_size = 1;

    for (key_seq, key) in keys_pressed.into_iter().enumerate() {
        let letter_count = keys[key].len();

        let mut comb = 0;
        while comb < possible_combs {
            for letter in 0..letter_count{
                for _ in 0..spread_size {
                    combination_buffer[comb][key_seq] = keys[key][letter];
                    comb +=1;
                }
            }
        }
        spread_size *= letter_count;
    }

    unsafe { combination_buffer.into_iter().map(|v| String::from_utf8_unchecked(v)).collect() }
}


#[allow(dead_code)]
pub fn check_perfect_number(num: i32) -> bool {
    if num < 2{
        return false;
    }

    let mut n = 2;

    let mut accum = 1;

    while n *n  < num{
        if num%n == 0{
            // println!("Divisor {}  next {}", n, num/n);
            accum+= num/n;
            accum+=n;
        }

        n+=1;
    }

    accum == num
}


#[cfg(test)]
mod test {
    use std::{assert_eq, println, vec};

    use crate::numbers::{add_digits, arrange_coins, arrange_coins2, btn_phone_number_to_letter_combinations, check_perfect_number, count_bits, count_bits2, fizz_buzz, guess_number, is_perfect_square, is_power_of_three, is_power_of_two, is_power_of_two_2, is_subsequence, is_subsequence2, is_ugly, missing_number, my_atoi, reverse, set_seed, to_hex};

    use super::is_happy;



    #[test]
    fn check_perfect_number_test(){
        assert!(check_perfect_number(28));
        assert!(!check_perfect_number(1));
        assert!(!check_perfect_number(1));
        assert!(!check_perfect_number(33));
        assert!(!check_perfect_number(4));
    }

    #[test]
    fn btn_phone_number_to_letter_combinations_test() {
        // println!("{:?}", btn_phone_number_to_letter_combinations("2379".to_string()));
        assert_eq!( vec!["aa", "ba", "ca", "ab", "bb", "cb", "ac", "bc", "cc"], btn_phone_number_to_letter_combinations("22".to_string()));
        assert_eq!( vec!["ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"], btn_phone_number_to_letter_combinations("23".to_string()));
        assert_eq!( vec!["adp", "bdp", "cdp", "aep", "bep", "cep", "afp", "bfp", "cfp", "adq", "bdq", "cdq", "aeq", "beq", "ceq", "afq", "bfq", "cfq", "adr", "bdr", "cdr", "aer", "ber", "cer", "afr", "bfr", "cfr", "ads", "bds", "cds", "aes", "bes", "ces", "afs", "bfs", "cfs"], btn_phone_number_to_letter_combinations("237".to_string()));
        assert_eq!( vec!["adpw", "bdpw", "cdpw", "aepw", "bepw", "cepw", "afpw", "bfpw", "cfpw", "adqw", "bdqw", "cdqw", "aeqw", "beqw", "ceqw", "afqw", "bfqw", "cfqw", "adrw", "bdrw", "cdrw", "aerw", "berw", "cerw", "afrw", "bfrw", "cfrw", "adsw", "bdsw", "cdsw", "aesw", "besw", "cesw", "afsw", "bfsw", "cfsw", "adpx", "bdpx", "cdpx", "aepx", "bepx", "cepx", "afpx", "bfpx", "cfpx", "adqx", "bdqx", "cdqx", "aeqx", "beqx", "ceqx", "afqx", "bfqx", "cfqx", "adrx", "bdrx", "cdrx", "aerx", "berx", "cerx", "afrx", "bfrx", "cfrx", "adsx", "bdsx", "cdsx", "aesx", "besx", "cesx", "afsx", "bfsx", "cfsx", "adpy", "bdpy", "cdpy", "aepy", "bepy", "cepy", "afpy", "bfpy", "cfpy", "adqy", "bdqy", "cdqy", "aeqy", "beqy", "ceqy", "afqy", "bfqy", "cfqy", "adry", "bdry", "cdry", "aery", "bery", "cery", "afry", "bfry", "cfry", "adsy", "bdsy", "cdsy", "aesy", "besy", "cesy", "afsy", "bfsy", "cfsy", "adpz", "bdpz", "cdpz", "aepz", "bepz", "cepz", "afpz", "bfpz", "cfpz", "adqz", "bdqz", "cdqz", "aeqz", "beqz", "ceqz", "afqz", "bfqz", "cfqz", "adrz", "bdrz", "cdrz", "aerz", "berz", "cerz", "afrz", "bfrz", "cfrz", "adsz", "bdsz", "cdsz", "aesz", "besz", "cesz", "afsz", "bfsz", "cfsz"], btn_phone_number_to_letter_combinations("2379".to_string()));
    }


    #[test]
    fn my_atoi_test() {
        println!("{}", '0' as i32);
        assert_eq!(100, my_atoi("100".to_string()));
        assert_eq!(536454, my_atoi("536454".to_string()));
        assert_eq!(-536454, my_atoi("-536454".to_string()));
        assert_eq!(-536454, my_atoi("-536454.532".to_string()));
        assert_eq!(i32::MIN, my_atoi((i32::MIN).to_string()));
        assert_eq!(i32::MIN, my_atoi((i32::MIN as i64 - 1).to_string()));
        assert_eq!(i32::MAX, my_atoi((i32::MAX).to_string()));
        assert_eq!(i32::MAX, my_atoi((i32::MAX as i64 + 1).to_string()));
        assert_eq!(0, my_atoi("0".to_string()));
        assert_eq!(0, my_atoi("words and 987".to_string()));
        assert_eq!(-42, my_atoi("   -42".to_string()));
        assert_eq!(42, my_atoi("   +42".to_string()));
        assert_eq!(0, my_atoi("+-12".to_string()));
        assert_eq!(2147483647, my_atoi("21474836460".to_string()));
        assert_eq!(0, my_atoi("00000-42a1234".to_string()));
    }

    #[test]
    fn reverse_number_digits_test() {
        assert_eq!(321, reverse(123));
        assert_eq!(5, reverse(500));
        assert_eq!(1, reverse(1));
        assert_eq!(9, reverse(9));
        assert_eq!(0, reverse(2147483647));
        assert_eq!(0, reverse(-2147483648));
        assert_eq!(-412, reverse(-214));
    }

    #[test]
    fn arrange_coins_test() {
        assert_eq!(2, arrange_coins(5));
        assert_eq!(1, arrange_coins(1));
        assert_eq!(3, arrange_coins(8));
        assert_eq!(0, arrange_coins(0));
        assert_eq!(10635, arrange_coins(56564124));

        assert_eq!(1, arrange_coins2(1));
        assert_eq!(2, arrange_coins2(5));
        assert_eq!(3, arrange_coins2(8));
        assert_eq!(0, arrange_coins2(0));
        assert_eq!(10635, arrange_coins2(56564124));
        assert_eq!(60070, arrange_coins2(1804289383));
    }

    #[test]
    fn fizz_buzz_test() {
        assert_eq!(vec!["1".to_string(), "2".to_string(), "Fizz".to_string()], fizz_buzz(3));
        assert_eq!(vec!["1".to_string(), "2".to_string(), "Fizz".to_string(), "4".to_string(), "Buzz".to_string()], fizz_buzz(5));
    }

    #[test]
    fn to_hex_test() {
        assert_eq!("f", to_hex(15));
        assert_eq!("1", to_hex(1));
        assert_eq!("a", to_hex(10));
        assert_eq!("b", to_hex(11));
        assert_eq!("10", to_hex(16));
        assert_eq!("11", to_hex(17));
        assert_eq!("ffffffff", to_hex(-1));
        assert_eq!("0", to_hex(0));
    }


    #[test]
    fn is_subsequence_test() {
        assert!(is_subsequence("abc".to_string(), "ahbgdc".to_string()));
        assert!(!is_subsequence("axc".to_string(), "ahbgdc".to_string()));

        assert!(is_subsequence2("abc".to_string(), "ahbgdc".to_string()));
        assert!(!is_subsequence2("axc".to_string(), "ahbgdc".to_string()));
    }

    #[test]
    fn guess_number_test() {
        unsafe { assert_eq!(set_seed(6), guess_number(10)); }
        unsafe { assert_eq!(set_seed(0), guess_number(10)); }
        unsafe { assert_eq!(set_seed(9), guess_number(10)); }
        unsafe { assert_eq!(set_seed(10), guess_number(10)); }
        unsafe { assert_eq!(set_seed(5), guess_number(10)); }
        unsafe { assert_eq!(set_seed(7), guess_number(10)); }
        unsafe { assert_eq!(set_seed(7), guess_number(11)); }
        unsafe { assert_eq!(set_seed(6), guess_number(11)); }
        unsafe { assert_eq!(set_seed(5), guess_number(11)); }
        unsafe { assert_eq!(set_seed(8), guess_number(11)); }
    }

    #[test]
    fn divide() {
        println!("{}", 17 / 4);
    }

    #[test]
    fn is_perfect_square_test() {
        assert!(is_perfect_square(4));
        assert!(is_perfect_square(9));
        assert!(is_perfect_square(16));
        assert!(is_perfect_square(25));
        assert!(is_perfect_square(36));
        assert!(is_perfect_square(49));
        assert!(is_perfect_square(808201));
        assert!(!is_perfect_square(3));
        assert!(!is_perfect_square(8));
        assert!(!is_perfect_square(6));
        assert!(!is_perfect_square(7));
    }

    #[test]
    fn count_bits_test() {
        assert_eq!(count_bits(2), [0, 1, 1]);
        assert_eq!(count_bits(5), [0, 1, 1, 2, 1, 2]);
        assert_eq!(count_bits2(2), [0, 1, 1]);
        assert_eq!(count_bits2(5), [0, 1, 1, 2, 1, 2]);
    }

    #[test]
    fn print_bit_count_test() {
        for i in 0..21_i8 {
            println!("{:0>2} {:0>08b} --> {:0>2}", i, i, i.count_ones())
        }
    }

    #[test]
    fn missing_number_test() {
        assert_eq!(missing_number(vec![0, 1, 3]), 2);
    }

    #[test]
    fn is_power_of_3() {
        let x = is_power_of_three(2147483647);
        println!("{}", x);
    }

    #[test]
    fn is_ugly_test() {
        assert!(is_ugly(6));
        assert!(is_ugly(1));
        assert!(!is_ugly(7));
        assert!(!is_ugly(11));
        assert!(!is_ugly(13));
        assert!(is_ugly(15));
    }

    #[test]
    fn add_digits_test() {
        assert_eq!(add_digits(38), 2);
        assert_eq!(add_digits(0), 0);
        assert_eq!(add_digits(11), 2);
        assert_eq!(add_digits(55), 1);
        assert_eq!(add_digits(555), 6);
    }

    #[test]
    fn count_one() {
        println!("{}", 2_i32.count_ones());
        println!("{}", 4_i32.count_ones());
        println!("{}", 5_i32.count_ones());
        println!("{}", (3 | 4));
    }

    #[test]
    fn is_power_of_two_test() {
        let m = 2_i32.pow(30);
        println!("{}", m);
        assert!(!is_power_of_two(0));
        assert!(is_power_of_two(1));
        assert!(is_power_of_two(2));
        assert!(is_power_of_two(4));
        assert!(!is_power_of_two(-2147483648));
        assert!(is_power_of_two(2_i32.pow(30)));
    }


    #[test]
    fn is_power_of_two_test2() {
        let m = 2_i32.pow(30);
        println!("{}", m);
        assert!(!is_power_of_two_2(0));
        assert!(is_power_of_two_2(1));
        assert!(is_power_of_two_2(2));
        assert!(is_power_of_two_2(4));
        assert!(!is_power_of_two_2(-2147483648));
        assert!(is_power_of_two_2(2_i32.pow(30)));
    }


    #[test]
    fn is_happy_test() {
        assert_eq!(is_happy(1), true);
        assert_eq!(is_happy(19), true);
        assert_eq!(is_happy(2), false);
    }
}