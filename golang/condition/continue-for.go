/*continue statement representation
*/
package main
import ("fmt";"os";"strconv")
func main(){
	sum:=0
	a, _ := strconv.Atoi(os.Args[1])
	for i:=1;i<=a;i++{
      if i % 2 == 0 {
		  continue //If condition match continue and skip that 
	  }
	  sum += i 
	}
	fmt.Printf("Sum of odd number from 0 to %d is %d\n",a,sum)
}

/*
$ go run function/test.go 10
Sum of odd number from 0 to 10 is 25
*/
