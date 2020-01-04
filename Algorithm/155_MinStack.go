package main

/*
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MinStack struct {
	stack       []int
	stackTop    int
	minStack    []int
	minStackTop int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		-1,
		[]int{},
		-1,
	}
}

func (this *MinStack) Push(x int) {
	this.stackTop++
	this.stack = append(this.stack, x)
	if this.minStackTop == -1 || this.minStack[this.minStackTop] >= x {
		this.minStackTop++
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	if this.minStack[this.minStackTop] == this.stack[this.stackTop] {
		this.minStack = this.minStack[:this.minStackTop]
		this.minStackTop--
	}
	this.stack = this.stack[:this.stackTop]
	this.stackTop--
}

func (this *MinStack) Top() int {
	return this.stack[this.stackTop]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.minStackTop]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
