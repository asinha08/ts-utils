package main

import "fmt"
import "github.com/asinha08/ts-utils/servicecall"

func main() {
	fmt.Println("Hi")
	servicecall.Post("https://trackingstream.com", "text/html", []byte("ddd"))
}
