#[allow(dead_code)]
struct Solution(i32);
#[allow(dead_code)]
impl Solution {
    pub fn is_bad_version(&self, version:i32)->bool{
        version>= self.0
    }

    pub fn first_bad_version(&self, n: i32) -> i32 {
        let (mut high, mut low) = (n, 0);

        while low <=high{
            let mid = low +(high -low)/2;
            if self.is_bad_version(mid){
                high = mid-1;
            }else{
                low = mid+1;
            }
        }

        low

    }
}