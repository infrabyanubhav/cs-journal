package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	addMap := make(map[string]string)

	var name, address string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your address: ")
	fmt.Scan(&address)

	addMap["name"] = name
	addMap["address"] = address

	jsonData, err := json.Marshal(addMap)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(string(jsonData))

}
