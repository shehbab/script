package main
import ("fmt";"time")
func main(){
   switch a := time.Now().Hour() 
   {
   case a > 18:
	fmt.Println("Good Evening ",a)
   case a > 12:
	fmt.Println("Good Afternoon ",a)
   case a > 6:
	fmt.Println("Good Morning ",a)
   default:
	fmt.Println("Good Night ",a)
   }
}
