pub fn reverse_bits(x: u32) -> u32 {
    let mut n: u32 = x;
    n = (n & 0xffff0000) >> 16 | (n & 0x0000ffff) << 16;
    n = (n & 0xff00ff00) >> 8 | (n & 0x00ff00ff) << 8;
    n = (n & 0xf0f0f0f0) >> 4 | (n & 0x0f0f0f0f) << 4;
    n = (n & 0xcccccccc) >> 2 | (n & 0x33333333) << 2;
    n = (n & 0xaaaaaaaa) >> 1 | (n & 0x55555555) << 1;
    n
}




pub fn hamming_weight(n: u32) -> i32 {
    let mut weight =0;

    for i in 0..32 {
      if n>>i & 1 == 1{
          weight+=1;
      }
    }

    weight
}

pub fn hamming_weight2(n:u32) ->i32{
    let mut weight=0;
    let  mut x = n;
    for _ in 0..32 {
        if x&1 > 0 {
            weight+=1;
        }

        x>>=1;
    }

    weight
}


pub fn reverse_bits_iter(x: u32) -> u32 {
    let mut result = 0;
    let mut x = x;

    for _ in 0..32 {
        result <<= 1;
        result |= x & 1;
        x >>= 1;
    }

    result
}


#[allow(dead_code)]
pub fn read_binary_watch(turned_on: i32) -> Vec<String> {
    let ones = turned_on as u32;
    let mut result = Vec::with_capacity(240);

    for i in 0..=11_u32 {
        for j in 0..=59_u32 {
          if i.count_ones() + j.count_ones() == ones {
              result.push(format!("{}:{:0>2}", i, j))
          }
        }
    }

    result
}

#[allow(dead_code)]
pub fn hamming_distance(x: i32, y: i32) -> i32 {
    (x^y).count_ones() as i32
}

#[cfg(test)]
mod tests{
    use std::{assert_eq, vec};
    use crate::bitwise::{hamming_distance, hamming_weight2, read_binary_watch};

    use super::hamming_weight;
    use super::reverse_bits;
    use super::reverse_bits_iter;

    #[test]
    fn humming_difference_check(){
        println!("{:0b} {:0b} {:0b}", 5, 7 , 5^7);
        println!("{:0b} {:0b} {:0b}", 4, 1 , 4^1);

        assert_eq!(2, hamming_distance(1, 4));
        assert_eq!(1, hamming_distance(1, 3));
    }

    #[test]
    fn read_binary_watch_test(){
        assert_eq!(vec!["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"],read_binary_watch(1));
        assert_eq!(Vec::<String>::new(),read_binary_watch(9));
    }

    #[test]
    fn test_hamming_weight(){
        assert_eq!(hamming_weight(1), 1);
        assert_eq!(hamming_weight(2), 1);
        assert_eq!(hamming_weight(3), 2);
        assert_eq!(hamming_weight(5), 2);
    }

    #[test]
    fn test_hamming_weight2(){
        assert_eq!(hamming_weight2(1), 1);
        assert_eq!(hamming_weight2(2), 1);
        assert_eq!(hamming_weight2(3), 2);
        assert_eq!(hamming_weight2(5), 2);
    }

    #[test]
    fn test_reverse_bits(){
        assert_eq!(reverse_bits(0b00000000000000000000000000000000), 0b00000000000000000000000000000000);
        assert_eq!(reverse_bits(0b11111111111111111111111111111111), 0b11111111111111111111111111111111);
        assert_eq!(reverse_bits(0b01010101010101010101010101010101), 0b10101010101010101010101010101010);
        assert_eq!(reverse_bits(43261596), 964176192);
        assert_eq!(reverse_bits(4294967293), 3221225471)
    }

    #[test]
    fn test_reverse_bits_iter(){
        assert_eq!(reverse_bits_iter(0b00000000000000000000000000000000), 0b00000000000000000000000000000000);
        assert_eq!(reverse_bits_iter(0b11111111111111111111111111111111), 0b11111111111111111111111111111111);
        assert_eq!(reverse_bits_iter(0b01010101010101010101010101010101), 0b10101010101010101010101010101010);
        assert_eq!(reverse_bits_iter(43261596), 964176192);
        assert_eq!(reverse_bits_iter(4294967293), 3221225471)
    }
}