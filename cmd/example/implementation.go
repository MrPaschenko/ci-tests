package lab2
import "fmt"
type StackNode struct {
	data string
	next * StackNode
}
func getStackNode(data string, top * StackNode) * StackNode {
	return &StackNode {
		data,
		top,
	}
}
type MyStack struct {
	top * StackNode
	count int
}
func getMyStack() * MyStack {
	return &MyStack {
		nil,
		0,
	}
}

func(this MyStack) size() int {
	return this.count
}
func(this MyStack) isEmpty() bool {
	if this.size() > 0 {
		return false
	} else {
		return true
	}
}

func(this *MyStack) push(data string) {
	this.top = getStackNode(data, this.top)
	this.count++
}
func(this *MyStack) pop() string {
	var temp string = ""
	if this.isEmpty() == false {
		temp = this.top.data
		this.top = this.top.next
		this.count--
	}
	return temp
}

func(this MyStack) peek() string {
	if !this.isEmpty() {
		return this.top.data
	} else {
		return ""
	}
}

func isOperator(text byte) bool {
	if text == '+' || text == '-' || text == '*' || 
		text == '/' || text == '^' || text == '%' {
		return true
	}
	return false
}

func isOperands(text byte) bool {
	if (text >= '0' && text <= '9') || 
	   (text >= 'a' && text <= 'z') || 
	   (text >= 'A' && text <= 'Z') {
		return true
	}
	return false
}

func prefixToinfix(prefix string) string {
	var size int = len(prefix)
	var s * MyStack = getMyStack()
	var auxiliary string = ""
	var op1 string = ""
	var op2 string = ""
	var isValid bool = true
	for i := size - 1 ; i >= 0 ; i-- {
		if isOperator(prefix[i]) {
			if s.size() > 1 {
				op1 = s.pop()
				op2 = s.pop()
				auxiliary = "(" + op1 + 
				string(prefix[i]) + op2 + ")"
				s.push(auxiliary)
			} else {
				isValid = false
			}
		} else if isOperands(prefix[i]) {
			auxiliary = string(prefix[i])
			s.push(auxiliary)
		} else {
			isValid = false
		}
	}
	infix := s.pop()
	return infix
}