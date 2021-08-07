package main
import (
	"fmt"
	"os"
    "strconv"
)
func convertStrToInt(i string){
   if len(os.Args) == 2 {
 //strconv.Atoi From String To Int
 //strconv.ParseFloat From String To Float
     	if s, err := strconv.ParseFloat(i, 64); err == nil { 
	    fmt.Printf("%v is of Type \"%T\"\n",s,s)
	    }

	}else{
	   fmt.Println("Enter only one argument\n")
   }
}
func main(){
     convertStrToInt(os.Args[1])
}

/******
$ go run function/test.go 231
231 is of Type "float64"
$ go run function/test.go 231.23
231.23 is of Type "float64"
*****/
