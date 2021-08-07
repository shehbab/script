package main
import ("fmt";"strconv";"os")
func stringInt(i string) int {    //String Type CmdLine Arg to Int 
	y, _ := strconv.Atoi(i)
	return y
}
func intString(i int) string { //Int to String 
	r := strconv.Itoa(i)
    return r
}  
func main(){
	s := stringInt(os.Args[1]) 
	fmt.Printf("String to Int:%T %d\n",s,s) //Two time variable pass to get the Data Type and Variable Value
	s1 := intString(s)
	fmt.Printf("Int to String:%T %s\n",s1,s1)
}
/*
$ go run testing.go 42
String to Int:int 42
Int to String:string 42
*/
