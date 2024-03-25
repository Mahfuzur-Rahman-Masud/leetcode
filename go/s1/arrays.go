package s1

import (
	"math"
	"sort"
	"strconv"
)

func distributeCandies(candyType []int) int {
	count := 0
	m := map[int]struct{}{}
	half := len(candyType) / 2

	for _, candy := range candyType {
		if _, ok := m[candy]; ok {
			continue
		}

		m[candy] = struct{}{}
		count++
		if count == half {
			break
		}

	}
	return count
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	rr := len(mat)
	if rr == 0 {
		return mat
	}

	cc := len(mat[0])
	size := rr * cc

	if r >= rr && c >= cc {
		return mat
	} else if c > cc {
		c = size / r
	}

	if r > rr {
		r = rr
	}

	for c*r < size {
		if c < cc {
			c++
		} else {
			rr++
		}
	}

	out := make([][]int, r)

	x, y := 0, 0
	for _, row := range mat {
		for _, val := range row {
			if y == r {
				out = append(out, []int{})
			}

			out[y] = append(out[y], val)
			x++
			if x == c {
				x = 0
				y++
			}

		}
	}

	return out
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	l := len(nums) - 1

	sum := 0

	for ; l > 0; l -= 2 {
		sum += nums[l-1]
	}

	return sum
}

func findRelativeRanks(score []int) []string {
	l := len(score)
	out := make([]string, l)

	sc := make([]int, l)
	copy(sc, score)
	sort.Ints(sc)

	ranks := map[int]int{}

	r := l
	for _, scc := range sc {
		ranks[scc] = r
		r--
	}

	for i, scc := range score {
		r = ranks[scc]
		switch r {
		case 1:
			out[i] = "Gold Medal"
		case 2:
			out[i] = "Silver Medal"
		case 3:
			out[i] = "Bronze Medal"
		default:
			out[i] = strconv.Itoa(r)
		}
	}

	return out
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	end := 0
	affect := 0

	for _, t := range timeSeries {
		if end > t {
			affect = affect + duration + t - end
		} else {
			affect += duration
		}

		end = t + duration
	}

	return affect
}

func constructRectangle2(area int) []int {
	root := int(math.Sqrt(float64(area)))

	left, right := root, root

	for left > 0 && right <= area {
		res := left * right

		if res == area {
			return []int{right, left}
		} else if res > area {
			left--
		} else {
			right++
		}
	}

	return []int{area, 1}
}

func constructRectangle(area int) []int {
	root := int(math.Sqrt(float64(area)))

	if root*root == area {
		return []int{root, root}
	}

	out := []int{area, 1}
	// gap := area - 1

	left, right := 2, area/2

	for left < right {
		res := left * right

		if res == area {
			out[0] = right
			out[1] = left
			left++
			right--
		} else if res > area {
			right--
		} else {
			left++
		}
	}

	return out
}

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
func findMaxConsecutiveOnes(nums []int) int {
	count, max := 0, 0

	for _, n := range nums {
		count = count*n + n

		if count > max {
			max = count
		}
	}

	return max
}

/**
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	len := len(nums1) + len(nums2)

	if l1 == 0 && l2 == 0 {
		return -1
	}

	if len == 1 && l1 == 0 {
		return float64(nums2[0])
	}

	if len == 1 && l2 == 0 {
		return float64(nums1[0])
	}

	left := 0
	right := len - 1
	left1 := 0
	left2 := 0

	right1 := l1 - 1
	right2 := l2 - 1
	lv := 0
	rv := 0

	for left <= right {

		lx, ly := math.MaxInt, math.MaxInt
		if left1 < l1 {
			lx = nums1[left1]
		}

		if left2 < l2 {
			ly = nums2[left2]
		}

		if lx < ly {
			lv = lx
			left1++
		} else {
			lv = ly
			left2++
		}

		rx, ry := math.MinInt, math.MinInt
		if right1 >= 0 {
			rx = nums1[right1]
		}

		if right2 >= 0 {
			ry = nums2[right2]
		}

		if rx > ry {
			rv = rx
			right1--
		} else {
			rv = ry
			right2--
		}

		left++
		right--

	}

	return (float64(lv) + float64(rv)) / 2

}

/**
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/

func maxArea(height []int) int {
	/**
	Observations:
		any y with value of 0 can never for part of the container area and safe to skip
		any y2 <= y1 where  i1 < i2 can not add value as opening node but can still be part of closing node
		 example: 5 0 4 5 7
		 -------  ^      ^     2nd 5 or the 4 cannot be part of solution to the right

	*/

	mx := 0
	l, r := 0, len(height)-1

	for l < r {
		var h int = 0

		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			r--
			h = height[r]
		}

		v := h * (r - l + 1)
		if v > mx {
			mx = v
		}

	}

	return mx

}

type point struct {
	x int
	y int
}

func islandPerimeter2(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	perimeter := 0
	r, c := 0, 0

out:
	for ; r < rows; r++ {
		for ; c < cols; c++ {
			if grid[r][c] == 1 {
				break out
			}
		}
	}

	if grid[r][c] == 0 {
		return perimeter
	}

	points := make([]point, 0, rows*cols/2)
	mark, limit := 0, 1
	points = append(points, point{r, c})

	// infinite loop
	for mark < limit {
		p := points[mark]
		r, c = p.x, p.y

		r--
		if r < 0 || grid[r][c] == 0 {
			perimeter++
		} else {
			points = append(points, point{r, c})
			limit++
		}

		r += 2
		if r == rows || grid[r][c] == 0 {
			perimeter++
		} else {
			points = append(points, point{r, c})
			limit++
		}
		r--

		c--
		if c < 0 || grid[r][c] == 0 {
			perimeter++
		} else {
			points = append(points, point{r, c})
			limit++
		}
		c += 2

		if c == cols || grid[r][c] == 0 {
			perimeter++
		} else {
			points = append(points, point{r, c})
			limit++
		}
		c--

		mark++
	}

	return perimeter
}

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	water := 0
	perimeter := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == water {
				continue
			}

			r--
			if r < 0 || grid[r][c] == water {
				perimeter++
			}
			r += 2
			if r == rows || grid[r][c] == water {
				perimeter++
			}
			r--

			c--
			if c < 0 || grid[r][c] == water {
				perimeter++
			}
			c += 2
			if c == cols || grid[r][c] == water {
				perimeter++
			}
			c--
		}
	}

	return perimeter
}

func findPerimeter(grid [][]int, r, c, rows, cols, sum int) int {
	// var left uint8 = 1
	// var top uint8 = 2
	// var right uint8 = 4
	// var bottom uint8 = 8

	r--
	if r >= 0 && grid[r][c] == 1 {
		sum--
		grid[r][c] = 0
		sum = findPerimeter(grid, r, c, rows, cols, sum+3)
	}

	r += 2
	if r < rows && grid[r][c] == 1 {
		sum--
		grid[r][c] = 0
		sum = findPerimeter(grid, r, c, rows, cols, sum+3)
	}

	r--
	c--

	if c >= 0 && grid[r][c] == 1 {
		sum--
		grid[r][c] = 0
		sum = findPerimeter(grid, r, c, rows, cols, sum+3)
	}

	c++

	if c < cols && grid[r][c] == 1 {
		sum--
		grid[r][c] = 0
		sum = findPerimeter(grid, r, c, rows, cols, sum+3)
	}

	return sum
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	out := 0
	valueIndex := -1
	l := len(s)

	for _, greed := range g {
		valueIndex++
		for ; valueIndex < l; valueIndex++ {
			if s[valueIndex] >= greed {
				out++
				break
			}
		}

		if valueIndex == l-1 {
			break
		}
	}

	return out
}

func findContentChildrenBruteForce(greed []int, value []int) int {

	out := 0
	l := len(value)

	for _, g := range greed {

		for i := 0; i < l; i++ {
			if value[i] >= g {
				out++
				value[i] = 0
				break
			}
		}

	}

	return out
}

func findDisappearedNumbers2(nums []int) []int {
	l := len(nums)
	out := make([]int, 0, l)

	for i := 0; i < l; i++ {
		flagIndex := int(math.Abs(float64(nums[i])) - 1)
		nums[flagIndex] = int(-math.Abs(float64(nums[flagIndex])))
	}

	for i, n := range nums {
		if n > 0 {
			out = append(out, i+1)
		}
	}

	return out
}

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	out := make([]int, 0, l)
	expect := 1
	sort.Ints(nums)

	for _, n := range nums {
		if n > expect {
			for i := expect; i < n; i++ {
				out = append(out, i)
			}
		}

		expect = n + 1
	}

	for i := expect; i <= l; i++ {
		out = append(out, i)
	}

	return out
}

func intersect_all(nums1 []int, nums2 []int) []int {
	out := make([]int, 0, len(nums1))
	counter := [1001]int{}

	for _, n := range nums1 {
		counter[n]++
	}

	for _, n := range nums2 {
		if counter[n] != 0 {
			counter[n]--
			out = append(out, n)
		}
	}

	return out
}

func reverseString(s []byte) {
	l := len(s) - 1
	min := len(s) / 2

	for i := 0; i < min; i++ {
		i2 := l - i
		s[i], s[i2] = s[i2], s[i]
	}
}

func summaryRanges(nums []int) []string {
	rs := -1
	l := len(nums) - 1
	p := 0
	out := []string{}

	for i, n := range nums {

		if rs == -1 {
			rs = n
			p = n
		}

		if n-p > 1 {

			if rs == p {
				out = append(out, strconv.Itoa(p))
				rs = n
			} else {
				out = append(out, strconv.Itoa(rs)+"->"+strconv.Itoa(p))
				rs = n
			}
		}

		p = n

		if l == i {
			if rs == p {
				out = append(out, strconv.Itoa(p))
				rs = n
			} else {
				out = append(out, strconv.Itoa(rs)+"->"+strconv.Itoa(p))
				rs = n
			}
		}

	}

	return out
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if i2, ok := m[n]; ok && i-i2 <= k {
			return true
		}

		m[n] = i
	}
	return false
}

func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}

		m[n] = struct{}{}
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

func moveZeroes(nums []int) {
	z := 0

	for i, n := range nums {
		if n == 0 {
			z++
		} else if z > 0 {
			nums[i] = 0
			nums[i-z] = n
		}
	}
}

type NumArray struct {
	v []int
}

func Constructor2(nums []int) NumArray {
	l := len(nums)
	runningSum := make([]int, l)

	sum := 0
	for i, n := range nums {
		sum += n
		runningSum[i] = sum
	}

	return NumArray{runningSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.v[right]
	}
	return this.v[right] - this.v[left-1]
}
