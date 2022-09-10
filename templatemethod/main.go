package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== CharDisplay ===")
	charDisplay := NewCharDisplay('G')
	Display(charDisplay)

	fmt.Println("=== StringDisplay ===")
	strDisplay := NewStringDisplay("golang")
	Display(strDisplay)
}

// type AbstractDisplay struct{
//	printer Printer
// }
//
// // Display これが「抽象クラスで定義している処理の枠組みを定めたメソッド」に該当する
// func (a *AbstractDisplay) Display() {
// 	fmt.Println(a.printer.Open())
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(a.printer.Print())
// 	}
// 	fmt.Println(a.printer.Close())
// }

// いろいろ考えてみたけど、GoだとInterfaceをシンプルに実装するのが一番しっくりきた。
// だけどこれを Template Method パターンと言えるのか？

type Printer interface {
	Open() string
	Print() string
	Close() string
}

func Display(p Printer) {
	fmt.Print(p.Open())
	for i := 0; i < 5; i++ {
		fmt.Print(p.Print())
	}
	fmt.Print(p.Close())
}

type CharDisplay struct {
	//	*AbstractDisplay
	char rune
}

func NewCharDisplay(char rune) *CharDisplay {
	return &CharDisplay{
		//	AbstractDisplay: &AbstractDisplay{},
		char: char,
	}
}

func (c *CharDisplay) Open() string {
	return "<<"
}

func (c *CharDisplay) Print() string {
	return string(c.char)
}

func (c *CharDisplay) Close() string {
	return ">>\n"
}

type StringDisplay struct {
	// *AbstractDisplay
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		// AbstractDisplay: &AbstractDisplay{},
		str: str,
	}
}

func (s *StringDisplay) Open() string {
	return s.line()
}

func (s *StringDisplay) line() string {
	line := "+"
	for i := 0; i < utf8.RuneCountInString(s.str); i++ {
		line += "-"
	}
	line += "+\n"

	return line
}

func (s *StringDisplay) Print() string {
	return fmt.Sprintf("|%s|\n", s.str)
}

func (s *StringDisplay) Close() string {
	return s.line()
}
