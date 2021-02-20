package main

import "fmt"

type array struct {
	length int
	data []interface{}
}

func (a *array) Delete(i int) {
	RemoveIndex(a.data, i)
	a.length--
}

func RemoveIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}

func (a *array) Get(i int) interface{}{
	if i > a.length {
		panic("Out of array bound")
	}
	return a.data[i]
}

func (a *array) Add(i interface{}) {
	a.data[a.length] = i
	a.length++
}

func (a *array) Size() int {
	return a.length
}

type Array interface {
	Add(interface{})
	Size() int
	Get(int) interface{}
	Delete(int)
}

func NewArray() Array {
	return &array{length: 0, data: []interface{}{}}
}



func main() {
	arr := NewArray()
	arr.Add(5)
	arr.Add(2)
	arr.Add(3)
	fmt.Println(arr.Get(0))
	fmt.Println(arr.Get(1))
	fmt.Println(arr.Get(2))
	fmt.Printf("Size of array %d\n", arr.Size())
	arr.Delete(0)
	fmt.Println(arr.Get(0))
	fmt.Println(arr.Get(1))
	fmt.Println(arr.Get(2))
	fmt.Printf("Size of array %d", arr.Size())
}
