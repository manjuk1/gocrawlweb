package main

import (
	"fmt"
	"github.com/manjuk1/gocrawlweb/links"
)

func main(){
	fmt.Println("Hello World! Starting a Web crawler project")
	fmt.Println("Calling the Link Package Extract method")
	links.Extract("hhtp://abc.com")
}
