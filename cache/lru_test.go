// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-01-13 15:00
// version: 1.0.0
// desc   :

package cache

import (
	"container/list"
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	back := l.PushBack(2)
	l.PushBack(3)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	l.MoveToFront(back)

	fmt.Println("=============")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func TestPut(t *testing.T) {
	l := NewLruCache(6)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	fmt.Println(l.String())
	fmt.Println("=====")
	l.Put(4, 4)
	fmt.Println(l.String())
	fmt.Println("=====")
	l.Get(3)
	fmt.Println(l.String())
	fmt.Println("=====")
	l.Remove(2)
	fmt.Println(l.String())
}


func TestPut2(t *testing.T) {
	l := NewLruCache(6)
	l.Put(1, 11)
	l.Put(2, 22)
	l.Put(3, 33)
	i := l.List()
	fmt.Println(i)
	fmt.Println("=====")
	l.Put(4, 44)
	i = l.List()
	fmt.Println(i)
	fmt.Println("=====")
	l.Get(3)
	i = l.List()
	fmt.Println(i)
}
