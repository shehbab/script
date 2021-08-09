package main
import ("fmt";"os")
func conv()  bool{
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	for _, element := range strSlice {
		if element == os.Args[1] {
          return true
		}
      
	}
   return false
}

func main(){
   if conv() {
      fmt.Println("Found!")
   }  
}


/*
$ go run testing.go hello

$ go run testing.go Japan
Found!
*/
