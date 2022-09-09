package main

func main() {
	var p Print = NewPrintBanner("Go言語でデザインパターン勉強")
	p.PrintWeak()
	p.PrintStrong()
}
