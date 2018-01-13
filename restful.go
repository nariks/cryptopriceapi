package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var coin string
	fmt.Println("Enter name of cyrptocoin:")
	fmt.Scanln(&coin)
	response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/" + coin)
	errorCheck(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	errorCheck(err)
	fmt.Println(string(contents))
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
