package main
import ("fmt";"os")
func findSeason(){
	a := os.Args[1]
	
	if len(os.Args) !=2 {
		fmt.Printf("Enter the single Month\n")    
		return
   }else if  ! elementSearch() {
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
// Function for Searching months
func elementSearch() bool {
	month := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"} 
	for _, element := range month {
		if element == os.Args[1] {
          return true
		}
      
	}
   return false
}
func main(){
	findSeason() //Call the function from main
}

/*
$ go run function/test.go Jul
Raining

$ go run function/test.go Nov
Autumn

$ go run function/test.go Mid
"Mid" is an invalid month
*/
