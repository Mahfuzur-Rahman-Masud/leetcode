package s1

type MyStack struct {
	mark int
	c    []int
}

func Constructor() MyStack {
	return MyStack{-1, make([]int, 64)}
}

func (this *MyStack) Push(x int) {
	this.mark++
	if cap(this.c) < this.mark+1 {
		c2 := make([]int, this.mark*2)
		copy(c2, this.c)
		this.c = c2
	}

	this.c[this.mark] = x
}

func (this *MyStack) Pop() int {
	if this.mark < 0 {
		return 0
	}

	t := this.c[this.mark]
	this.mark--
	return t
}

func (this *MyStack) Top() int {
	if this.mark < 0 {
		return 0
	}

	return this.c[this.mark]
}

func (this *MyStack) Empty() bool {
	return this.mark == -1
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
