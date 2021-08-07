package main
import ("fmt";"strings";"os";"strconv")
func conv()  {
	const usage = `
	Feet to Meters:-
	Usage:- go run FeetToMeter.go 23
	`	
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return 
	}
	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: '%s' is not a number.\n", arg)
		//return
	}
	meters := feet * 0.3048
	fmt.Printf("meters: %f \n",meters)
	return  
}
func main(){
    conv() //Calling function
}
