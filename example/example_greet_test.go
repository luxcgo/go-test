package greet_test

import (
	"fmt"

	// Needed for initialize side effect
	_ "image/png"
)

var m = make(map[string]string)

func ExampleDog_Hello() {
	m = make(map[string]string)
	m["ETH"] = "以太坊"
	m["BTC"] = "比特币"
	m["DOGE"] = "狗狗币"
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Unordered output:
	// BTC 比特币
	// ETH 以太坊
	// DOGE 狗狗币
}
