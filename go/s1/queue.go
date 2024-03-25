package s1

type MyQueue struct {
	mark int
	q    []int
}

func ConstructorQueue() MyQueue {
	return MyQueue{
		mark: -1,
		q:    make([]int, 64),
	}
}

func (this *MyQueue) Push(x int) {
	this.mark++
	if cap(this.q) < this.mark+1 {
		q := make([]int, this.mark*2)
		copy(q, this.q)
		this.q = q
	}

	this.q[this.mark] = x
}

func (this *MyQueue) Pop() int {
	if this.mark == -1 {
		return 0
	}

	this.mark--
	out := this.q[0]
	this.q = this.q[1:]
	return out
}

func (this *MyQueue) Peek() int {
	if this.mark == -1 {
		return 0
	}

	return this.q[0]
}

func (this *MyQueue) Empty() bool {
	return this.mark == -1
}
