use std::collections::{HashMap, HashSet};

#[allow(dead_code)]
pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
    let mut out = vec![];
    let mut i = 0;

    while i < nums.len() {
        let start = nums[i];
        while i + 1 < nums.len() && nums[i] + 1 == nums[i + 1] {
            i += 1;
        }

        if start == nums[i] {
            out.push(start.to_string());
        } else {
            out.push(format!("{}->{}", start, nums[i]))
        }
        i += 1;
    }

    out
}

#[allow(dead_code)]
pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
    let mut m = HashMap::with_capacity(nums.len());

    for (i, n) in nums.iter().enumerate() {
        match m.insert(n, i) {
            Some(i2)if i - i2 <= k as usize => {
                return true;
            }

            _ => {}
        }
    }

    false
}

#[allow(dead_code)]
pub fn contains_duplicate2(nums: Vec<i32>) -> bool {
    let mut exists = HashSet::new();
    !nums.into_iter().all(|x| exists.insert(x))
}

#[allow(dead_code)]
pub fn contains_duplicate(nums: Vec<i32>) -> bool {
    let mut set: HashSet<i32> = HashSet::new();

    for n in nums {
        if !set.insert(n) {
            return true;
        }
    }

    false
}


#[allow(dead_code)]
pub fn move_zeroes_snow_ball(nums: &mut Vec<i32>) {
    let mut z = 0;
    for i in 0..nums.len() {
        let n = nums[i];
        if n == 0 {
            z += 1;
        } else if z > 0 {
            nums[i] = 0;
            nums[i - z] = n;
        }
    }
}


#[allow(dead_code)]
pub fn move_zeroes(nums: &mut Vec<i32>) {
    let mut z = Vec::new();
    nums.retain(|x| {
        if *x != 0 {
            true
        } else {
            z.push(0);
            false
        }
    });

    nums.append(&mut z);
}

#[allow(dead_code)]
struct NumArray {
    val: Vec<i32>,
}


/**
 * `&self` means the method takes an immutable reference
 * If you need a mutable reference, change it to `&mut self` instead.
 */
#[allow(dead_code)]
impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        let mut sum = 0;
        let mut rs = Vec::with_capacity(nums.len());
        for (_, &n) in nums.iter().enumerate() {
            sum += n;
            rs.push(sum);
        }

        Self {
            val: rs
        }
    }

    fn sum_range(&self, left: i32, right: i32) -> i32 {
        if left == 0 {
            return self.val[right as usize];
        }

        self.val[right as usize] - self.val[left as usize - 1]
    }
}


#[allow(dead_code)]
pub fn reverse_string(s: &mut Vec<char>) {
    let l = s.len() - 1;
    let m = s.len() / 2;

    for i in 0..m {
        let i2 = l - i;
        let c = s[i];
        s[i] = s[i2];
        s[i2] = c;
    }
}

#[allow(dead_code)]
pub fn reverse_vowels(mut s: String) -> String {
    let b = unsafe { s.as_bytes_mut() };
    let (mut l, mut h) = (0_usize, b.len() - 1);

    let is_vowel = |x: u8| -> bool{
        let x = if x < 91 {
            x + 32
        } else {
            x
        };
        matches!(x, 97 | 101 |105 | 111 | 117)
    };

    while l < h {
        if !is_vowel(b[l]) {
            l += 1;
        } else {
            if !is_vowel(b[h]) {
                h -= 1;
            } else {
                let t = b[l];

                b[l] = b[h];
                b[h] = t;

                l += 1;
                h -= 1;
            }
        }
    }
    s
}


#[allow(dead_code)]
pub fn reverse_vowels2(s: String) -> String {
    let mut b = s.into_bytes();

    let (mut l, mut h) = (0_usize, b.len() - 1);


    while l < h {
        let mut c = b[l];
        if c < 91 {
            c += 32
        }

        if c == 97 || c == 101 || c == 105 || c == 111 || c == 117 {
            let mut c2 = b[h];
            if c2 < 91 {
                c2 += 32
            }

            if c2 == 97 || c2 == 101 || c2 == 105 || c2 == 111 || c2 == 117 {
                let cc = b[l];
                b[l] = b[h];
                b[h] = cc;

                l += 1;
                h -= 1;
            } else {
                h -= 1;
            }
        } else {
            l += 1;
        }
    }

    String::from_utf8(b).unwrap()
}

#[allow(dead_code)]
pub fn intersection2(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    let mut seen = HashSet::new();
    nums1.into_iter()
        .filter(move |x| seen.insert(*x))
        .filter(|x| nums2.contains(x))
        .collect::<Vec<i32>>()
}


#[allow(dead_code)]
pub fn intersection(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    HashSet::<i32>::from_iter(nums1)
        .intersection(&HashSet::<i32>::from_iter(nums2))
        .cloned()
        .collect::<Vec<i32>>()
}

#[allow(dead_code)]
pub fn intersection_all(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    let mut a_map: std::collections::HashMap<i32, i32> = HashMap::new();
    let mut out = Vec::new();

    for n in nums2 {
        *a_map.entry(n).or_insert(0) += 1;
    }

    for n in nums1 {
        match a_map.get_mut(&n) {
            Some(v) if v > &mut 0 => {
                *v -= 1;
                out.push(n);
            }
            _ => {}
        }
    }

    out
}

#[allow(dead_code)]
pub fn intersection_all2(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
    let mut out = Vec::with_capacity(nums1.len());
    let mut counter: Vec<i32> = vec![0; 1001];

    nums2.iter()
        .for_each(|x| counter[*x as usize] += 1);

    for x in nums1 {
        if counter[x as usize] != 0 {
            counter[x as usize] -= 1;
            out.push(x);
        }
    }

    out
}


#[allow(dead_code)]
pub fn third_max(mut nums: Vec<i32>) -> i32 {
    nums.sort_by(|a, b| b.cmp(a));

    let max = nums[0];
    let mut third = nums[0];
    let mut counter = 1;

    for i in 1..nums.len() {
        let n = nums[i];
        if n < third {
            third = n;
            counter += 1;
            if counter == 3 {
                break;
            }
        }
    }

    if counter < 3 {
        return max;
    }

    third
}


// satisfied if greed > value, each value can be used once
pub fn find_greed_satisfaction(mut greed: Vec<i32>, mut values: Vec<i32>) -> i32 {
    greed.sort();
    values.sort();
    let mut satisfaction = 0;
    let mut value_index = -1_i32;
    let value_count = values.len() as i32;

    for g in greed {
        value_index += 1;
        while value_index < value_count {
            if values[value_index as usize] >= g {
                satisfaction += 1;
                break;
            }

            value_index += 1;
        }

        if value_index == value_count - 1 {
            break;
        }
    }


    satisfaction
}

#[allow(dead_code)]
pub fn find_disappeared_numbers(mut nums: Vec<i32>) -> Vec<i32> {
    nums.sort();
    let mut out = Vec::with_capacity(nums.len());
    let mut expect = 1;
    let l = nums.len() as i32;

    // println!("{:?}", nums);

    for n in nums {
        if n > expect {
            for x in expect..n {
                out.push(x);
            }
        }
        expect = n + 1;
    }

    for n in expect..=l {
        out.push(n);
    }

    out
}

#[allow(dead_code)]
pub fn find_disappeared_numbers_in_place(mut nums: Vec<i32>) -> Vec<i32> {
    for i in 0..nums.len() {
        // 1 based index
        //[4,3,2,7,8,2,3,1]
        //[4,3,2,-7,8,2,3,1]
        //[4,3,-2,-7,8,2,3,1]
        //[4,-3,-2,-7,8,2,3,1]
        //[4,-3,-2,-7,8,2,-3,1]
        //[4,-3,-2,-7,8,2,-3,-1]
        //[-4,-3,-2,-7,8,2,-3,-1]
        //by positive index [-4,-3,-2,-7,8,2,-3,-1] [5, 6]

        // [2,2,3,4,5,5,5,5]
        // [2,-2,-3,-4,-5,5,5,5]
        // by positive index [1, 6, 7,8]

        // find appropriate zeroBased index for the number if it is a sorted array
        let flag_index = nums[i].abs() as usize - 1; // abs=> value can be previously flagged by making negative
        nums[flag_index] = -nums[flag_index].abs() // mark the location to indicate number is seen
    }

    nums.iter().enumerate()
        .filter_map(|(i, n)| if n > &0 { Some(i as i32 + 1) } else { None })
        .collect()
}

// #[allow(dead_code)]
// struct Node {
//     x: i32,
//     v: i32,
//     p: i32,
// }

#[allow(dead_code)]
pub fn max_area(height: Vec<i32>) -> i32 {
    let mut max = 0;

    for (i, h) in height.iter().enumerate() {
        match height.iter().enumerate().filter(|(ii, hh)| ii != &i && hh >= &h)
            .map(|(ii, _)| h * (ii as i32 - i as i32).abs())
            .max() {
            Some(m) if m > max => max = m,
            _ => {}
        }
    }


    max
}


#[allow(dead_code)]
pub fn max_area2(height: Vec<i32>) -> i32 {
    struct Node {
        x: usize,
        v: usize,
        p: usize,
        h: usize,
    }

    let mut max = 0;
    let l = height.len() - 1;
    let mut nodes = std::collections::HashMap::<usize, Node>::new();

    for (i, h) in height.iter().enumerate() {
        let (x, v, p, h) = (i, 0 as usize, (l - i) * (*h as usize), *h as usize);
        if h == 0 {
            continue;
        }

        let mut has_potential = true;
        let mut removables = vec![];

        for (_, mut node) in nodes.iter_mut() {
            node.v = (i - node.x) * std::cmp::min(node.h, h);
            if node.v > max {
                max = node.v;
            }

            if max > node.p {
                removables.push(node.x)
            }

            if !has_potential {
                continue;
            }

            if p < max {
                has_potential = false;
                continue;
            }

            if p < node.p && h >= node.h {
                has_potential = false;
                continue;
            }
        }

        for r in removables {
            nodes.remove(&r);
        }

        if has_potential {
            nodes.insert(x, Node { x, v, p, h });
        }
    }


    max as i32
}


#[allow(dead_code)]
fn max_area3(height: Vec<i32>) -> i32 {
    let mut max = 0;
    let (mut l, mut r) = (0, height.len() - 1);

    while l < r {
        let h = std::cmp::min(height[l], height[r]);
        max = std::cmp::max(max, h * (r - l) as i32);
        if height[l] < height[r] {
            l += 1;
        } else {
            r -= 1;
        }
    }


    max
}


#[allow(dead_code)]
pub fn three_sum_timeout(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let l = nums.len();
    let mut out = std::collections::HashSet::with_capacity(l / 3);


    for x in 0..l {
        for y in (x + 1)..l {
            for z in (y + 1)..l {
                if nums[x] + nums[y] + nums[z] == 0 {
                    let mut tsum = vec![nums[x], nums[y], nums[z]];
                    tsum.sort();
                    out.insert(tsum);
                }
            }
        }
    }

    out.iter().cloned().collect()
}


#[allow(dead_code)]
pub fn three_sum_three_pointer_fail(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let l = nums.len();
    if l < 3 {
        return Vec::new();
    }


    let mut out = std::collections::HashSet::with_capacity(l / 3);
    let (mut l, mut r) = (0, l - 1);

    while l < r {
        for m in (l + 1)..r {
            if nums[l] + nums[m] + nums[r] == 0 {
                let mut v = vec![nums[l], nums[m], nums[r]];
                v.sort();
                out.insert(v);
                break;
            }
        }

        if nums[l] < nums[r] {
            l += 1;
        } else {
            r -= 1;
        }
    }


    out.iter().cloned().collect()
}

#[allow(dead_code)]
pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut nums = nums;
    nums.sort();
    let mut res = Vec::new();

    for i in 0..nums.len() {
        let mut l = i + 1;
        let mut r = nums.len() - 1;

        if i > 0 && nums[i] == nums[i - 1] {
            continue;
        }

        while l < r {
            let sum = nums[i] + nums[l] + nums[r];
            if sum == 0 {
                res.push(vec![nums[i], nums[l], nums[r]]);
                l += 1;
                while l < r && nums[l] == nums[l - 1] {
                    l += 1;
                }
            } else if sum < 0 {
                l += 1;
            } else { r -= 1; }
        }
    }
    res
}


#[allow(dead_code)]
pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
    let mut best_proximity = i32::MAX;
    let mut closest = 0;
    let mut nums = nums;
    nums.sort();

    for i in 0..nums.len() {
        if i > 0 && nums[i] == nums[i - 1] {
            continue;
        }

        let mut l = i + 1;
        let mut r = nums.len() - 1;

        while l < r {
            let sum = nums[i] + nums[l] + nums[r];
            if sum == target {
                return sum;
            }

            let proximity = (target - sum).abs();

            if proximity < best_proximity {
                best_proximity = proximity;
                closest = sum;
            } else if sum > target {
                r -= 1;
            } else {
                l += 1;
            }
        }
    }


    closest
}

#[allow(dead_code)]
pub fn four_sum(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    let mut out = Vec::new();
    let len = nums.len();

    if len < 4 {
        return out;
    }
    let mut nums = nums;
    nums.sort();

    for i in 0..len - 3 {
        if nums[i] as i64 * 4 > target as i64 { break; }
        if i > 0 && nums[i] == nums[i - 1] { continue; };

        for j in i + 1..len - 2 {
            if nums[j] as i64 * 3 + nums[i] as i64 > target as i64 { continue; }
            if j > i + 1 && nums[j] == nums[j - 1] { continue; }

            let (mut left, mut right) = (j + 1, len - 1);

            while left < right {
                let sum: i64 = nums[i] as i64 + nums[j] as i64 + nums[left] as i64 + nums[right] as i64;
                if sum == target as i64 {
                    out.push(vec![nums[i], nums[j], nums[left], nums[right]]);
                    left += 1;
                    right -= 1;

                    while left < right && nums[left] == nums[left - 1] { left += 1; }
                    while left < right && nums[right] == nums[right + 1] { right -= 1; }
                } else if sum < target as i64 {
                    left += 1;
                } else {
                    right -= 1;
                }
            }
        }
    }

    out
}

#[allow(dead_code)]
fn four_sum2(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    let mut out = Vec::new();
    let len = nums.len();
    if len < 4 {
        return out;
    }

    let mut nums = nums.into_iter().map(|x| x as i64).collect::<Vec<i64>>();
    nums.sort();
    let target = target as i64;

    let mut i = 0;
    let t_end = len - 1;
    let j_end = len - 2;
    let i_end = len - 3;

    while i < i_end {
        let a = nums[i];
        let mut j = i + 1;

        while j < j_end {
            let b = nums[j];
            let (mut l, mut r) = (j + 1, t_end);

            while l < r {
                let (c, d) = (nums[l], nums[r]);
                let sum = a + b + c + d;
                if sum == target {
                    out.push(vec![a as i32, b as i32, c as i32, d as i32]);
                    while l < r && nums[l] == c { l += 1; }
                    while l < r && nums[r] == d { r -= 1; }
                } else if sum < target {
                    while l < r && nums[l] == c { l += 1; }
                } else {
                    while l < r && nums[r] == d { r -= 1; }
                }
            }
            while j < j_end && nums[j] == b { j += 1; }
        }
        while i < i_end && nums[i] == a { i += 1; }
    }

    out
}


#[allow(dead_code)]
pub fn find_words(words: Vec<String>) -> Vec<String> {
    let mut m: std::collections::HashMap<u8, u8> = std::collections::HashMap::with_capacity(123);
    m.insert(113, 0);
    m.insert(119, 0);
    m.insert(101, 0);
    m.insert(114, 0);
    m.insert(116, 0);
    m.insert(121, 0);
    m.insert(117, 0);
    m.insert(105, 0);
    m.insert(111, 0);
    m.insert(112, 0);

    m.insert(97, 1);
    m.insert(115, 1);
    m.insert(100, 1);
    m.insert(102, 1);
    m.insert(103, 1);
    m.insert(104, 1);
    m.insert(106, 1);
    m.insert(107, 1);
    m.insert(108, 1);

    let mut out = Vec::with_capacity(words.len());

    for w in words.into_iter() {
        println!();
        println!("{}", w);
        let mut prev = Some(&4);
        for mut c in w.bytes() {
            if c < 97 { c += 32; }
            let v = m.get(&c);


            if prev == Some(&4) {
                prev = v;
                println!("c: {}, current: {:?} prev: {:?}", c, v, prev);

                continue;
            }

            println!("c: {}, current: {:?} prev: {:?}", c, v, prev);

            if v != prev {
                prev = Some(&2);
                break;
            }
        }

        if prev != Some(&2) {
            println!("OK {:?}", prev);
            out.push(w);
        } else {
            println!("Exit {:?}", prev);
        }
    }

    out
}


#[allow(dead_code)]
pub fn find_relative_ranks(score: Vec<i32>) -> Vec<String> {
    let mut out = vec![String::new(); score.len()];
    let mut heap = std::collections::BinaryHeap::new();
    score.into_iter().enumerate()
        .for_each(|a| heap.push((a.1, a.0)));

    let mut rank = 0;
    while let Some((_, i)) = heap.pop() {
        rank += 1;

        let rank_name = match rank {
            1 => "Gold Medal".to_string(),
            2 => "Silver Medal".to_string(),
            3 => "Bronze Medal".to_string(),
            other => other.to_string()
        };

        out[i] = rank_name;
    }

    out
}

#[allow(dead_code)]
pub fn find_lhs(nums: Vec<i32>) -> i32 {
    let mut nums = nums;
    nums.sort();
    let len = nums.len();
    let mut max = 0;
    let mut mark_in = 0;
    let mut mark_out = 1;
    let mut is_first = false;
    let mut next_in = 1;

    while mark_in < mark_out && mark_out < len {
        let v1 = nums[mark_in];
        let v2 = nums[mark_out];
        let diff = v2 - v1;

        if diff == 0 {
            is_first = false;
            mark_out += 1;
        } else if diff == 1 {
            if !is_first {
                is_first = true;
                next_in = mark_out;
            }

            let spread = mark_out - mark_in + 1;
            if spread > max {
                max = spread;
            }
            mark_out += 1;
        } else {
            mark_in = next_in;
            mark_out = mark_in + 1;
            next_in = mark_out;
            is_first = false;
        }
    }

    max as i32
}

#[allow(dead_code)]
pub fn find_restaurant(list1: Vec<String>, list2: Vec<String>) -> Vec<String> {
    let mut out = vec![];
    let mut sum = i32::MAX as usize;

    for (i1, s1) in list1.iter().enumerate() {
        for (i2, s2) in list2.iter().enumerate() {
            let local_sum = i1 + i2;
            if local_sum > sum {
                break;
            }

            if s1 != s2 {
                continue;
            }

            if local_sum == sum {
                out.push(s1.clone())
            } else if local_sum < sum {
                sum = local_sum;
                out = vec![s1.clone()]
            }
        }
    }

    out
}


#[allow(dead_code)]
pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
    let len = flowerbed.len();
    let end = flowerbed.len() - 1;

    let can_plant = |x| {
        if flowerbed[x] == 1 { return false; }
        if x != 0 && flowerbed[x - 1] == 1 { return false; }
        if x != end && flowerbed[x + 1] == 1 { return false; }
        true
    };

    let mut count = 0;
    let mut i = 0;
    while i < len {
        if can_plant(i) {
            count += 1;
            if count == n {
                return true;
            }

            i += 2;
        } else {
            i += 1;
        }
    }

    count >= n
}


#[allow(dead_code)]
pub fn maximum_product(nums: Vec<i32>) -> i32 {
    let mut nums = nums;
    let end = nums.len() - 1;

    nums.sort_by(|a, b| {
        return b.cmp(a);
    });

    if nums[0] < 0 {
        return nums[0] * nums[1] * nums[2];
    }

    let m2 = nums[end] * nums[end - 1];
    let m3 = nums[1] * nums[2];

    nums[0] * core::cmp::max(m2, m3)
}


#[allow(dead_code)]
pub fn find_max_average_old(nums: Vec<i32>, k: i32) -> f64 {
    let k = k as usize;
    let mut i = k  -1;
    let mut max = f64::MIN;
    let len = nums.len();


    while i < len{
        if i >= k && nums[i] < nums[i-k]{
            i+=1;
            continue
        }

        let mut sum = 0;
        for j in 0..k {
           sum+= nums[i-j];
        }

        let avg = sum as f64 /k as f64;
        if avg > max{
            max = avg
        }
        i+=1
    }


    max
}


#[allow(dead_code)]
pub fn find_max_average(nums: Vec<i32>, k: i32) -> f64{
    nums.windows(k as usize)
        .map(|w|w.iter().sum::<i32>())
        .max()
        .unwrap() as f64 / k as f64
}

#[allow(dead_code)]
pub fn find_error_nums(nums: Vec<i32>) -> Vec<i32> {
    let mut contains = vec![false; nums.len()];
    let mut out = vec![0, 2];

    for n in nums{
        let i = n as usize -1;
        if contains[i]{
            out[0] = n;
        }else{
            contains[i] = true
        }
    }

    let miss = contains.into_iter().enumerate()
        .filter(|(_, v)|v==&false)
        .find_map(|(i, _)|Some(i as i32 +1))
        .unwrap();

    out[1] = miss;


    out
}


#[allow(dead_code)]
pub fn find_length_of_lcis(nums: Vec<i32>) -> i32 {
    if nums.len() == 0 {
        return 0
    }

    let mut prev = nums[0]-1;
    let mut count = 0;
    let mut max = 0;

    for n in nums{
        if n > prev{
            count+=1;
            if count > max{
                 max = count;
            }
        } else{
            count = 1;
        }

        prev = n;
    }

    max
}



#[cfg(test)]
mod test {
    use std::{assert_eq, println, vec};

    use crate::arrays::{can_place_flowers, contains_duplicate, find_disappeared_numbers, find_disappeared_numbers_in_place, find_error_nums, find_greed_satisfaction, find_length_of_lcis, find_lhs, find_max_average, find_max_average_old, find_relative_ranks, find_restaurant, find_words, four_sum, four_sum2, intersection, intersection_all, intersection_all2, max_area, max_area2, max_area3, move_zeroes, NumArray, reverse_string, reverse_vowels, reverse_vowels2, third_max, three_sum, three_sum_closest};




    #[test]
    fn find_length_of_lcis_test(){
        assert_eq!(3, find_length_of_lcis(vec![1,3,5,4,7]));
        assert_eq!(1, find_length_of_lcis(vec![2,2,2,2,2]));
        assert_eq!(2, find_length_of_lcis(vec![2,2,1,2,2]));
        assert_eq!(3, find_length_of_lcis(vec![2,2,1,2,3,2]));
        assert_eq!(3, find_length_of_lcis(vec![2,2,1,2,10,2]));
    }

    #[test]
    fn find_error_nums_test(){
        assert_eq!(vec![2,3] , find_error_nums(vec![1,2,2,4]));
        assert_eq!(vec![1,2] , find_error_nums(vec![1,1]));
    }
    #[test]
    fn find_max_average_test(){
        assert_eq!(12.75000, find_max_average_old(vec![1,12,-5,-6,50,3],  4));
        assert_eq!(12.75000, find_max_average(vec![1,12,-5,-6,50,3],  4));
    }
    #[test]
    fn can_place_flower_test() {
        assert!(can_place_flowers(vec![1, 0, 0, 0, 1], 1));
        assert!(!can_place_flowers(vec![1, 0, 0, 0, 1], 2));
        assert!(!can_place_flowers(vec![0, 0, 1, 0, 1], 2));
        assert!(can_place_flowers(vec![0, 0, 1, 0, 1], 1));
        assert!(can_place_flowers(vec![0, 0, 1, 0, 0], 2));
        assert!(!can_place_flowers(vec![1, 0, 0, 0, 0, 1], 2));
    }

    #[test]
    fn find_restaurant_test() {
        assert_eq!(vec!["Shogun".to_string()], find_restaurant(
            vec!["Shogun".to_string(), "Tapioca Express".to_string(), "Burger King".to_string(), "KFC".to_string()],
            vec!["Piatti".to_string(), "The Grill at Torrey Pines".to_string(), "Hungry Hunter Steakhouse".to_string(), "Shogun".to_string()]));

        assert_eq!(vec!["Shogun".to_string()], find_restaurant(vec!["Shogun".to_string(), "Tapioca Express".to_string(), "Burger King".to_string(), "KFC".to_string()], vec!["KFC".to_string(), "Shogun".to_string(), "Burger King".to_string()]));

        assert_eq!(vec!["sad".to_string(), "happy".to_string()], find_restaurant(vec!["happy".to_string(), "sad".to_string(), "good".to_string()], vec!["sad".to_string(), "happy".to_string(), "good".to_string()]));
    }

    #[test]
    fn find_lhs_test() {
        assert_eq!(5, find_lhs(vec![1, 3, 2, 2, 5, 2, 3, 7]));
        assert_eq!(2, find_lhs(vec![1, 2, 3, 4]));
        assert_eq!(0, find_lhs(vec![1, 1, 1, 1]));
    }


    #[test]
    fn find_relative_ranks_test() {
        assert_eq!(vec!["Gold Medal".to_string(), "Silver Medal".to_string(), "Bronze Medal".to_string(), "4".to_string(), "5".to_string()], find_relative_ranks(vec![5, 4, 3, 2, 1]));
        assert_eq!(vec!["Gold Medal".to_string(), "5".to_string(), "Bronze Medal".to_string(), "Silver Medal".to_string(), "4".to_string()], find_relative_ranks(vec![10, 3, 8, 9, 4]));
    }


    #[test]
    fn find_words_test() {
        // assert_eq!(vec!["ASDFJKL".to_string(),  "eoiurowqeuo".to_string()] ,
        //            find_words(vec!["ASDFJKL".to_string(), "ABCED".to_string(), "eoiurowqeuo".to_string()]));

        assert_eq!(vec!["a".to_string()], find_words(vec!["abdfs".to_string(), "cccd".to_string(), "a".to_string(), "qwwewm".to_string()]));
    }

    #[test]
    fn four_sum_test() {
        assert_eq!(vec![vec![-3, -2, 2, 3], vec![-3, -1, 1, 3], vec![-3, 0, 0, 3], vec![-3, 0, 1, 2], vec![-2, -1, 0, 3], vec![-2, -1, 1, 2],
                        vec![-2, 0, 0, 2], vec![-1, 0, 0, 1]], four_sum(vec![-3, -2, -1, 0, 0, 1, 2, 3], 0));
        assert_eq!(vec![vec![-2, -1, 1, 2], vec![-1, -1, 1, 1]], four_sum(vec![-2, -1, -1, 1, 1, 2, 2], 0));
        assert_eq!(vec![vec![-2, -1, 1, 2], vec![-2, 0, 0, 2], vec![-1, 0, 0, 1]], four_sum(vec![1, 0, -1, 0, -2, 2], 0));
        assert_eq!(vec![vec![2, 2, 2, 2]], four_sum(vec![2, 2, 2, 2, 2], 8));
        assert_eq!(vec![vec![-1, -1, 0, 2]], four_sum(vec![-1, 2, -1, 0, 0, 0], 0));
        assert_eq!(Vec::<Vec<i32>>::new(), four_sum(vec![1, 2, 1, 0, 0, 0], 5));
        assert_eq!(vec![vec![0, 0, 0, 2], vec![0, 0, 1, 1]], four_sum(vec![1, 2, 1, 0, 0, 0], 2));
        assert_eq!(vec![vec![-2, -1, 2, 4], vec![-2, 0, 1, 4], vec![-1, 0, 0, 4], vec![-1, 1, 1, 2], vec![0, 0, 1, 2]], four_sum(vec![1, 2, 1, 0, 0, 0, 4, -1, -2], 3));


        assert_eq!(vec![vec![-3, -2, 2, 3], vec![-3, -1, 1, 3], vec![-3, 0, 0, 3], vec![-3, 0, 1, 2], vec![-2, -1, 0, 3], vec![-2, -1, 1, 2],
                        vec![-2, 0, 0, 2], vec![-1, 0, 0, 1]], four_sum2(vec![-3, -2, -1, 0, 0, 1, 2, 3], 0));
        assert_eq!(vec![vec![-2, -1, 1, 2], vec![-1, -1, 1, 1]], four_sum2(vec![-2, -1, -1, 1, 1, 2, 2], 0));
        assert_eq!(vec![vec![-2, -1, 1, 2], vec![-2, 0, 0, 2], vec![-1, 0, 0, 1]], four_sum2(vec![1, 0, -1, 0, -2, 2], 0));
        assert_eq!(vec![vec![2, 2, 2, 2]], four_sum2(vec![2, 2, 2, 2, 2], 8));
        assert_eq!(vec![vec![-1, -1, 0, 2]], four_sum2(vec![-1, 2, -1, 0, 0, 0], 0));
        assert_eq!(Vec::<Vec<i32>>::new(), four_sum2(vec![1, 2, 1, 0, 0, 0], 5));
        assert_eq!(vec![vec![0, 0, 0, 2], vec![0, 0, 1, 1]], four_sum2(vec![1, 2, 1, 0, 0, 0], 2));
        assert_eq!(vec![vec![-2, -1, 2, 4], vec![-2, 0, 1, 4], vec![-1, 0, 0, 4], vec![-1, 1, 1, 2], vec![0, 0, 1, 2]], four_sum2(vec![1, 2, 1, 0, 0, 0, 4, -1, -2], 3));
    }

    #[test]
    fn three_sum_test() {
        assert_eq!(Vec::<Vec<i32>>::new(), three_sum(vec![0, 1, 1]));
        assert_eq!(vec![vec![-1, 0, 1], vec![-1, -1, 2]], three_sum(vec![-1, 0, 1, 2, -1, -4]));
        assert_eq!(vec![vec![0, 0, 0]], three_sum(vec![0, 0, 0]));
    }


    #[test]
    fn tree_sum_closest_test() {
        assert_eq!(2, three_sum_closest(vec![-1, 2, 1, -4], 1));
        assert_eq!(0, three_sum_closest(vec![0, 0, 0], 1));
        assert_eq!(0, three_sum_closest(vec![0, 0, 0], 0));
    }

    #[test]
    fn max_area_test() {
        assert_eq!(49, max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]));
        assert_eq!(1, max_area(vec![1, 1]));
        assert_eq!(0, max_area(vec![0, 1]));
        assert_eq!(0, max_area(vec![0, 0]));
        assert_eq!(0, max_area(vec![0, 0, 5]));
        assert_eq!(6, max_area(vec![2, 0, 0, 5]));
        assert_eq!(6, max_area(vec![2, 0, 5, 5]));
        assert_eq!(8, max_area(vec![2, 4, 5, 5]));


        // assert_eq!(402471897, max_area(v));

        assert_eq!(49, max_area2(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]));
        assert_eq!(1, max_area2(vec![1, 1]));
        assert_eq!(0, max_area2(vec![0, 1]));
        assert_eq!(0, max_area2(vec![0, 0]));
        assert_eq!(0, max_area2(vec![0, 0, 5]));
        assert_eq!(6, max_area2(vec![2, 0, 0, 5]));
        assert_eq!(6, max_area2(vec![2, 0, 5, 5]));
        assert_eq!(8, max_area2(vec![2, 4, 5, 5]));
        // assert_eq!(402471897, max_area2(v));
        // assert_eq!(v2.len(), 100000);
        // assert_eq!(999990000, max_area2(v2));


        assert_eq!(49, max_area3(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]));
        assert_eq!(1, max_area3(vec![1, 1]));
        assert_eq!(0, max_area3(vec![0, 1]));
        assert_eq!(0, max_area3(vec![0, 0]));
        assert_eq!(0, max_area3(vec![0, 0, 5]));
        assert_eq!(6, max_area3(vec![2, 0, 0, 5]));
        assert_eq!(6, max_area3(vec![2, 0, 5, 5]));
        assert_eq!(8, max_area3(vec![2, 4, 5, 5]));
    }

    #[test]
    fn fun() {
        let x = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
        let y = "OK".to_string();
        println!("{:?} {}", x, y);
        println!("{:?} {}", x, y);
        // print(&x);
        // print(&x);
    }

    #[allow(dead_code)]
    fn print(x: &Vec<i32>) {
        println!("{:?}", x)
    }

    #[test]
    fn find_disappeared_numbers_test() {
        assert_eq!(vec![5, 6], find_disappeared_numbers(vec![4, 3, 2, 7, 8, 2, 3, 1]));
        assert_eq!(vec![2], find_disappeared_numbers(vec![1, 1]));
        assert_eq!(vec![1], find_disappeared_numbers(vec![2, 2]));
        assert_eq!(vec![1, 6, 7, 8], find_disappeared_numbers(vec![2, 2, 3, 4, 5, 5, 5, 5]));

        assert_eq!(vec![5, 6], find_disappeared_numbers_in_place(vec![4, 3, 2, 7, 8, 2, 3, 1]));
        assert_eq!(vec![2], find_disappeared_numbers_in_place(vec![1, 1]));
        assert_eq!(vec![1], find_disappeared_numbers_in_place(vec![2, 2]));
        assert_eq!(vec![1, 6, 7, 8], find_disappeared_numbers_in_place(vec![2, 2, 3, 4, 5, 5, 5, 5]));
    }

    #[test]
    fn find_greed_satisfaction_test() {
        assert_eq!(1, find_greed_satisfaction(vec![1, 2, 3], vec![1, 1]));
        assert_eq!(2, find_greed_satisfaction(vec![1, 2], vec![1, 2, 3]));
        assert_eq!(1, find_greed_satisfaction(vec![1, 2], vec![1, 1, 1]));
        assert_eq!(2, find_greed_satisfaction(vec![1, 2, 1], vec![1, 1, 1]));
        assert_eq!(1, find_greed_satisfaction(vec![1, 2, 3], vec![3]));
    }

    #[test]
    fn third_max_test() {
        assert_eq!(1, third_max(vec![3, 2, 1]));
        assert_eq!(3, third_max(vec![3, 2]));
        assert_eq!(1, third_max(vec![3, 2, 2, 2, 2, 1]));
        assert_eq!(3, third_max(vec![3, 2, 2, 2, 2]));
        assert_eq!(-2147483648, third_max(vec![1, 2147483647, -2147483648]));
    }

    #[test]
    fn intersection_all_test() {
        assert_eq!(intersection_all(vec![1, 2, 2, 1], vec![2, 2, 3]), vec![2, 2]);
        assert_eq!(intersection_all(vec![4, 9, 5], vec![9, 4, 9, 8, 4]), vec![4, 9]);

        assert_eq!(intersection_all2(vec![1, 2, 2, 1], vec![2, 2, 3]), vec![2, 2]);
        assert_eq!(intersection_all2(vec![4, 9, 5], vec![9, 4, 9, 8, 4]), vec![4, 9]);
    }

    #[test]
    fn intersection_test() {
        assert_eq!(intersection(vec![1, 2, 2, 1], vec![2, 2]), vec![2]);
        assert_eq!(intersection(vec![4, 9, 5], vec![9, 4, 9, 8, 4]), vec![4, 9]);
    }

    #[test]
    fn iter_both_end_test() {
        let s = "leetcode".to_string();
        let mut iter = s.as_bytes().iter();
        while let (Some(left), Some(right)) = { (iter.next(), iter.next_back()) } {
            println!("{} {}", left, right);
        }
    }

    #[test]
    fn reverse_vowels_test() {
        assert_eq!("eppla".to_string(), reverse_vowels("apple".to_string()));
        assert_eq!("holle".to_string(), reverse_vowels("hello".to_string()));
        assert_eq!("leotcede".to_string(), reverse_vowels("leetcode".to_string()));

        assert_eq!("eppla".to_string(), reverse_vowels2("apple".to_string()));
        assert_eq!("holle".to_string(), reverse_vowels2("hello".to_string()));
        assert_eq!("leotcede".to_string(), reverse_vowels2("leetcode".to_string()));
    }


    #[test]
    fn reverse_string_test() {
        let mut input: Vec<char> = vec!['h', 'e', 'l', 'l', 'o'];
        let output: Vec<char> = vec!['o', 'l', 'l', 'e', 'h'];
        reverse_string(&mut input);
        assert_eq!(input, output);

        let mut input: Vec<char> = vec!['H', 'a', 'n', 'n', 'a', 'h'];
        let output: Vec<char> = vec!['h', 'a', 'n', 'n', 'a', 'H'];
        reverse_string(&mut input);
        assert_eq!(input, output);
    }


    #[test]
    fn array_test() {
        let nums = vec![0, 1, 2, 3, 4, 5];
        let x = [0].iter().chain(nums.iter())
            .scan(0, |s, n| {
                *s += n;
                Some(*s)
            }).collect::<Vec<i32>>();

        println!("{:?}", x)
    }


    #[test]
    fn sum_range_test() {
        assert_eq!(0, NumArray::new(vec![0]).sum_range(0, 0));
        assert_eq!(-2, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(0, 0));
        assert_eq!(-2, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(0, 1));
        assert_eq!(1, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(0, 2));
        assert_eq!(3, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(2, 2));
        assert_eq!(3, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(1, 2));
        assert_eq!(-2, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(2, 3));
        assert_eq!(-3, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(3, 4));
        assert_eq!(0, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(2, 4));
        assert_eq!(-1, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(2, 5));
        assert_eq!(-3, NumArray::new(vec![-2, 0, 3, -5, 2, -1]).sum_range(0, 5));
    }

    #[test]
    fn move_zeroes_test() {
        let mut v = vec![0, 1, 0, 3, 12];
        move_zeroes(&mut v);
        println!("{:?}", v)
    }

    #[test]
    fn contains_duplicate_test() {
        assert!(contains_duplicate(vec![1, 2, 3, 1]));
        assert!(!contains_duplicate(vec![1, 2, 3, 4]));
        assert!(contains_duplicate(vec![1, 1, 1, 3, 3, 4, 3, 2, 4, 2]));
    }
}