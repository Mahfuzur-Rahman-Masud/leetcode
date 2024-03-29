test("count_subarrays", ()=>{
    expect(countSubarrays([1, 3, 2, 3, 3],2)).toBe(6)
    expect(countSubarrays([1, 4, 2, 1],3)).toBe(0)
    expect(countSubarrays([1],1)).toBe(1)
    expect(countSubarrays([1,2],1)).toBe(2)
})


function countSubarrays(nums: number[], k: number): number {
    let max = 0
    for (let i=0; i< nums.length; i++){
        if(nums[i]> max){
            max = nums[i]
        }
    }

    let out = 0
    let mark = 0
    let count = 0
    for (let i=0; i< nums.length; i++){
         let n = nums[i]
        if (n == max){
            count++
        }

        while (count >= k){
            if (nums[mark] == max){
                count--
            }

            mark++
        }

        
        out+=mark

    }

    return out;
};