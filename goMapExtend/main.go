package main

import "fmt"

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

}
