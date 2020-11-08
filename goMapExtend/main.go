package main

import "fmt"

//1.go map必须make之后才能使用
//2. 遍历方法有2种 方法  for k :=range m   只遍历key
//    for k,v :=range m  同时遍历key 和 value

//只遍历key
func travelMapKeyTest() {
	m := make(map[int]int)
	m[1] = 11
	m[2] = 22
	m[3] = 33

	mk := make([]int, 0, len(m))
	for k := range m {
		mk = append(mk, k)
	}

	fmt.Println(mk)
}

//遍历key和value
func travelMapKeyValueTest() {
	m := make(map[int]int)
	m[1] = 11
	m[2] = 22
	m[3] = 33

	for k, v := range m {
		fmt.Printf("k:%d v:%d", k, v)
	}
}

//在遍历中添加元素 行为是不确定的
/*
a 1
b 2
c 3

*/
func travelMapAddElementTest() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	for k, v := range m {
		m["c"] = 3
		fmt.Println(k, v)
	}
}

//在遍历map的时候删除元素，行为是不确定的
// 1 a 1

func travelMapDeleteElementTest() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	n := 0
	flag := false

	for k, v := range m {
		if !flag {
			flag = true
			if k == "a" {
				delete(m, "b")
			} else {
				delete(m, "a")
			}
		}
		n++
		fmt.Println(n, k, v)
	}
}

func main() {
	keys := []string{"key1", "key1", "key2", "key1", "key2", "key3", "key3", "key4"}
	m := make(map[string][]int)

	for i, key := range keys {
		mv, ok := m[key]
		if !ok {
			mv = make([]int, 0, 10)
			m[key] = mv
			fmt.Printf("!ok %s,%p,len_mv:%d,cap_mv:%d\n", key, mv, len(mv), cap(mv))
		}
		mv = append(mv, i) //更新的是mv的内容，m[key]指向的切片还是!ok里面的空切片，所以这里正确的是m[key]=append(mv,i)
		fmt.Printf("%s,%p,len_mv:%d,cap_mv:%d\n", key, mv, len(mv), cap(mv))
	}
	fmt.Println(m) //map[key1:[] key2:[] key3:[] key4:[]]

	fmt.Print("----------------------\n")
	travelMapKeyTest()
	fmt.Print("----------------------\n")
	travelMapKeyValueTest()
	fmt.Print("----------------------\n")
	travelMapAddElementTest()
	fmt.Print("----------------------\n")
	travelMapDeleteElementTest()
}
