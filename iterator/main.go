package main

import (
	"fmt"
	"iterator/book"
)

func main() {
	bookshelf := book.NewBookshelf(nil)
	bookshelf.AppendBook(book.NewBook("Javaで学ぶデザインパターン"))
	bookshelf.AppendBook(book.NewBook("DNSの教科書"))
	bookshelf.AppendBook(book.NewBook("Perfect Ruby"))
	it := bookshelf.Iterator()
	for it.HasNext() {
		b := it.Next()
		fmt.Println(b.GetName())
	}
}
