package main
import ("fmt";"os";"strconv")
func checkNo(n string){
   // fmt.Printf("No entered is %q\n",n)
    p, _ := strconv.ParseFloat(n, 64)         //Convert String to Float
    fmt.Println("The number entered:",p)
    p1 := int(p)                             //Convert Float to Int
    fmt.Printf("%.1f convert to int value = %d\n",p,p1)
    if p1 % 2 == 0 {
      fmt.Printf("%d is even\n",p1)
    } else {
      fmt.Printf("%d is odd\n",p1)
    }
}
func start(){
   if len(os.Args) != 2 {
      fmt.Println("Enter the Single Number.")
      return
   }
   n := os.Args[1]
   checkNo(n)
}
func main(){
    start()
}
/***
$ go run testing.go 27.6
The number entered: 27.6
27.6 convert to int value = 27
27 is odd
***/
