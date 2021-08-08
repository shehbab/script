package main
import ("fmt";"os")
func main(){
	switch a := os.Args[1]; a {
	case "Jan", "Feb", "March":
		fmt.Println("Winter")
	case "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Monsoon")
	case "Sep","Oct":
		fmt.Println("Monsoon")
	case "Dec", "Nov":
		fmt.Println("Autumn")
    default: 
	   fmt.Println("Invalid Month")
	}
}
