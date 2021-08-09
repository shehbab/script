package main
import ("fmt";"os";"sort")
func findSeason(){
	a := os.Args[1]
	month := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"} 
	sort.Strings(month)
	fmt.Println(sort.SearchStrings(month, a))
   if len(os.Args) !=2 {
		fmt.Printf("Enter the single Month\n")    
		return
   }else if sort.SearchStrings(month, a) == 12{
	    fmt.Printf("%q is an invalid month\n",a)
        return
   }else{

		switch a {
		case "Jan","Feb","Mar":
			fmt.Println("Winter\n")
		case "Apr","May","Jun":
			fmt.Println("Spring\n")
		case "Jul","Aug","Sep":
			fmt.Println("Raining\n") 
		case "Oct","Nov","Dec":
			fmt.Println("Autumn\n") 
		default:
			fmt.Println("Invalid month\n")
		}
   }
}
func main(){
	findSeason()
}
