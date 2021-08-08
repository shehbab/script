package main
import ("fmt"
   // "os"
)//Switch Case | read input
func checKing(city string) {
	switch city {
	case "Delhi":
		fmt.Println("India")
    case "Islamabad":
	    fmt.Println("Pakistan")
    case "Dhaka":
	    fmt.Println("Bengladesh")
	default:
		fmt.Println("No in the list")
 	}
 }
func main(){
	fmt.Print("Enter Your Capital of the Country: ")
    var r string
	fmt.Scanln(&r)
//	checKing(os.Args[1])
	checKing(r)
}
