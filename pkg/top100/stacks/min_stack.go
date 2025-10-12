package stacks

// NO.155
type MinStack struct {
	stack []int
	minn  []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minn:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minn) == 0 {
		this.minn = append(this.minn, val)
	} else {
		this.minn = append(this.minn, min(val, this.minn[len(this.minn)-1]))
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minn = this.minn[:len(this.minn)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minn[len(this.stack)-1]
}
