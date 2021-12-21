package stack

import "testing"


func TestStack(t *testing.T){
	stack := New[int]()
	stack.push(1)
	val,err := stack.pop()
	if err != nil {
		t.Errorf(err.Error())
	}else{
		t.Logf("pop val: %d", val)	
	}
	val,err = stack.peek()
	if err != nil {
		t.Errorf(err.Error())
	}else{
		t.Logf("peek val: %d", val)	
	}
}