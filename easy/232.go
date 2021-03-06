package main

type MyQueue struct {
	in, out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	inStack := []int{}
	outStack := []int{}
	return MyQueue{in: inStack, out: outStack}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			n := len(this.in)
			this.out = append(this.out, this.in[n-1])
			this.in = this.in[:n-1]
		}
	}
	n := len(this.out)
	pop := this.out[n-1]
	this.out = this.out[:n-1]
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			n := len(this.in)
			this.out = append(this.out, this.in[n-1])
			this.in = this.in[:n-1]
		}
	}
	n := len(this.out)
	pop := this.out[n-1]
	return pop
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}
