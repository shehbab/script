package main
import ("fmt";"strconv";"os")
func stringInt(i string) int {
	y, _ := strconv.Atoi(i)
	return y
}
func intString(i int) string {
	r := strconv.Itoa(i)
    return r
}  
func main(){
	s := stringInt(os.Args[1])
	fmt.Printf("String to Int:%T %d\n",s,s)
	s1 := intString(s)
	fmt.Printf("Int to String:%T %s\n",s1,s1)
}
