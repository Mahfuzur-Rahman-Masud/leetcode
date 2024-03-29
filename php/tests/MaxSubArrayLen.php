<?php

use PHPUnit\Framework\TestCase;

class MaxSubArrayLen extends TestCase{

    function maxSubarrayLength($nums, $k) {
        $m=[];
        $cur = 0;
        $start = 0; 
        $mx = 0;
        $len = sizeof($nums);


        while($cur < $len){
            $n = $nums[$cur];
            $cnt= $m[$n]??0;
            $m[$n]=$cnt+1;

            if($cnt ==$k){
                for($i=$start;$i< $len; $i++ ){
                    $nx = $nums[$i];
                    $m[$nx] = $m[$nx]-1;

                    if($nx === $n){
                        $start = $i+1;
                        if ($len-$start <= $k){
                            return $mx;
                        }

                        break;
                    }
                }
            }else{
                $l = $cur-$start+1;
                if($l > $mx){
                    $mx = $l;
                }
            }



            $cur ++;
        }
        
        return $mx;
    }

    public function testTrueIsTrue(){
        $this->assertEquals($this->maxSubarrayLength([1,2,3,1,2,3,1,2], 2) ,6);
        $this->assertEquals($this->maxSubarrayLength([1,2,1,2,1,2,1,2], 1) ,2);
        $this->assertEquals($this->maxSubarrayLength([5,5,5,5,5,5,5], 4) ,4);
        $this->assertEquals($this->maxSubarrayLength([1], 1) ,1);
    }
}