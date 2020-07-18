package main

import "github.com/YoungWizard/godemo/test"
import "fmt"
func main(){
	a,b:=1,2
	c:=test.Max(a,b)
	fmt.Println(c)
}