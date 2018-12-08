package main

import (
	"fmt"
	"time"
)


func main()  {
	go ready("Tea",2)
	go ready("Coffee",5)
	fmt.Println("I'm waiting")
	time.Sleep(10 * time.Second)

}
func ready(w string,sec int)  {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w,"is ready!")
}