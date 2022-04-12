package main

import "fmt"

func main() {

	//
	nums := []int{2, 3, 4}
	sum := 0
	// 在 Go 中不用的变量要用空白标识符来代替
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// range 返回两个值：索引和值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 迭代 map 键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// %s 字符串输出
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 只遍历K
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range 迭代字符串返回的是unicode码点
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
