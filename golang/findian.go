package main

import(
	"fmt"
	"strings"
	
)


func main(){
	var s string
	fmt.Scan(&s)
	if (strings.HasPrefix(s,"i") || strings.HasPrefix(s,"I"))  && (strings.Contains(s,"a") || strings.Contains(s,"A")) && (strings.HasSuffix(s,"n") || strings.HasSuffix(s,"N")){
		fmt.Printf("Found!")
	}else{
		fmt.Printf("Not Found!")
	}
}