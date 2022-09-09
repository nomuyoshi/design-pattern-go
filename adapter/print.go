package main

// 必要とされているインターフェース

type Print interface {
	PrintWeak()
	PrintStrong()
}
