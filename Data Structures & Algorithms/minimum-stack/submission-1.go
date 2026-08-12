type MinStack struct {
    values []int
    minimum []int
}

func Constructor() *MinStack {
    return &MinStack{
        values: []int{},
        minimum: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.values = append(this.values, val)
    if len(this.minimum) == 0 || val <= this.minimum[len(this.minimum) - 1] {
        this.minimum = append(this.minimum, val)
    }
}

func (this *MinStack) Pop() {
    if this.values[len(this.values) - 1] == this.minimum[len(this.minimum) - 1] {
        this.minimum = this.minimum[:len(this.minimum) -1]
    }
    this.values = this.values[:len(this.values) - 1]
}

func (this *MinStack) Top() int {
    return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
    return this.minimum[len(this.minimum)-1]
}
