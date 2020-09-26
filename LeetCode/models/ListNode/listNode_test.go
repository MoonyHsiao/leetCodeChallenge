package listNode

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestListLen(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	a := LineList{}
	count := 0
	for i := range arr {
		a.Append(arr[i])
		count++
	}
	len := ListLength(&a)
	expected := count
	assert.Equal(t, expected, len)
}

func TestNew(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	mylist := LineList{}
	for i := range arr {
		mylist.Append(arr[i])
	}
	//這個不會有事 應該是有初始化
	end := LineList{}
	// 這個會有事
	// var end LineList

	len := ListLength(&mylist)
	fmt.Printf("len:%v\n", len)
	len = ListLength(&end)
	fmt.Printf("len:%v\n", len)

	mylist.Display()
	res := mylist.GetElem(2)
	mylist.Insert(1, 100)
	mylist.Display()
	mylist.Remove(0)
	mylist.Display()
	fmt.Printf("res:%v\n", res)
	res = mylist.GetElem(6)
	end.InsertFront(1000)
	end.InsertAfter(99)
	end.InsertFront(11000)
	end.InsertAfter(88)
	end.Display()
	end.Clear()
	end.Clear()
	fmt.Printf("res:%v\n", res)
	end.InsertFront(11000)
	end.Display()
}
