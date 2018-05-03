package main

import (
	"github.com/no1charleswu/go_study/TTPPC/crawler"
)

func main() {
	// testOkex()
	// testHuobi()
	// testGate()
	// testZb()
}

func testHuobi() {
	crawler.Huobi()
}

func testOkex() {
	crawler.Okex()
}

func testGate() {
	crawler.Gate()
}

func testZb() {
	crawler.Zb()
}

func testBinance() {
	crawler.Binance()
}
