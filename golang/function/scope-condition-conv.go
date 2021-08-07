package main
import ( "fmt";"os";"strconv" )
func converts(){
	if a := os.Args; len(a) != 2{
		fmt.Printf("Give me a number.\n")
		return
	} else if n, err := strconv.Atoi(a[1]); err != nil {
        fmt.Printf("Cannot convert %q.\n",a[1])
		return
	} else {
        fmt.Printf("%s * 2 is %d\n", a[1], n*2)
		return
	}
}
func main(){
    converts()
} 
