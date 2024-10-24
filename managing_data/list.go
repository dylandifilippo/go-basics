package main

import "fmt"

func main() {
	websites := []map[string]string{
		{
			"site": "GolangCode",
			"link": "https://golangcode.com",
		},
		{
			"site": "Google",
			"link": "https://google.com",
		},
		{
			"site": "Twitter",
			"link": "https://twitter.com",
		},
	}
	for key, value := range websites {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
