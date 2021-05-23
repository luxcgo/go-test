package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["ETH"] = "以太坊"
	m["BTC"] = "比特币"
	m["DOGE"] = "狗狗币"
	for k, v := range m {
		fmt.Println(k, v)
	}

	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}

	for name, age := range personAge {
		fmt.Println(name, age)
	}
}
