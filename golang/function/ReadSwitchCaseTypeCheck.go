package main
import ("fmt";"strconv")

//Function to Check If the read input is number or not. Then check positive, negative and Zero
func CheckNo(){
  fmt.Print("Enter the number: ")
  var no string
  fmt.Scanln(&no)
  val, _ := strconv.ParseFloat(no, 64)
  if numCheck(no) {
	fmt.Printf("Supplied value %v is a number.\n", val)
	switch {
		case val > 0: 
		  fmt.Println("No is positive")
		case val < 0:
		  fmt.Println("No is negative")
		default:
		  fmt.Println("No is Zero") 		
	      }
  }else {
	fmt.Printf("Supplied value %q is not a number.\n", no)	 
	return
   }
}

//Function return boolean value 
func numCheck(v string) bool {
	_, err := strconv.ParseFloat(v, 64)
	return err == nil
 }

func main(){
	CheckNo()
}

/*
$ go run function/test.go 
Enter the number: Hello
Supplied value "Hello" is not a number.

$ go run function/test.go 
Enter the number: 12.34
Supplied value 12.34 is a number.
No is positive

$ go run function/test.go 
Enter the number: -12
Supplied value -12 is a number.
No is negative
*/
