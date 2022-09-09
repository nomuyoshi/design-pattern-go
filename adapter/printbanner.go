package main

// アダプターの役目を果たす
// すでにあるBannerと必要としているPrintインターフェース間のズレを埋める

type PrintBanner struct {
	banner Banner
}

func NewPrintBanner(title string) PrintBanner {
	return PrintBanner{
		banner: NewBanner(title),
	}
}

// PrintWeak Printインターフェースの実装
func (pb PrintBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

// PrintStrong Printインターフェースの実装
func (pb PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
