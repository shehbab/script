package main
import ("fmt";"os";"strconv")
func code(){
   a, _ := strconv.Atoi(os.Args[1])

   switch 
   {
   case a > 100:
	fmt.Println("big")
	fallthrough
   case a > 0:
	fmt.Println("postive")
	fallthrough
   default :
	fmt.Println("number")
   }
}
func main(){
	code()
}
/*
$ go run function/test.go 99
postive
number

$ go run function/test.go 102
big
postive
number
*/
