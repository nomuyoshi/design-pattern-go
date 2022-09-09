package main

import "fmt"

// すでに提供されているもの

type Banner struct {
	title string
}

func NewBanner(title string) Banner {
	return Banner{
		title: title,
	}
}

// ShowWithParen () で囲って出力
func (b Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.title)
}

// ShowWithAster * で囲って出力
func (b Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.title)
}
