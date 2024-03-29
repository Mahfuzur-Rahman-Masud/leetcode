function maxSubarrayLength(nums: number[], k: number): number {
    let mx =0
    let l = nums.length
    let v = 0
    let cur=0
    let start = 0

    let m :Record<number,number>= {}

    while (cur < l){
        v = nums[cur]
        let c =m[v]?? 0
        m[v]=c+1

        if (c ===k){
            for(let i =start; i<cur; i++){
                let v2 = nums[i]
                m[v2] = m[v2] -1

                if (v2 === v){
                    start = i+1

                    if(l-start < mx){
                        return mx
                    }


                    break
                }
            }
        }else{
            let len = cur-start+1

            if(len>mx){
                mx = len
            }
        }
        


        cur++
    }


    return mx
};



test("max_sub_array", ()=>{
    expect( maxSubarrayLength([1,2,3,1,2,3,1,2], 2)).toBe(6)
    expect( maxSubarrayLength([1,2,1,2,1,2,1,2], 1)).toBe(2)
    expect( maxSubarrayLength([5,5,5,5,5,5,5], 4)).toBe(4)
    expect( maxSubarrayLength([1], 1)).toBe(1)
    expect( maxSubarrayLength([1], 10)).toBe(1)
})

