package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Go 中字符串是一个只读的 byte类型的切片。作为以UTF-8为编码的文本容器
	// Go 中字符概念称为 rune -它是一个标识Unicode 编码的整数

	const s = "สวัสดี" // s是一个string,分配了一个literal value。泰语中表示 "hello".
	// Go 字符串是 UTF-8编码的文本。

	fmt.Println("Len", len(s)) // 字符串等价于 []byte 会存储在其中的原始字节的长度

	// 对字符串进行索引会在每个索引处生成原始字节值。这个循环构成s中 Unicode的所有字节的 十六进制值
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	// 计算字符串中有多少rune.使用 utf8包。RuneCountInstring 运行时取决于字符串的大小。一个泰语字符由多个 UTF-8 code point来标识
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串并解码每个rune及其在字符串中的偏移量
	for idx, runeValue := range s {
		fmt.Printf("%#U start at %d\n", runeValue, idx)
	}

	// 以下 迭代跟 range 循环一样，
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d \n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
