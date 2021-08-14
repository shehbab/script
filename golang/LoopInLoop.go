package main
import ("fmt";"os";"strconv")
func main(){
  fmt.Printf("%2s","X")  
  a, _ := strconv.Atoi(os.Args[1])
  for i:=0;i<=a;i++{
      fmt.Printf("%2d",i)
  }
  fmt.Println()
  for i:=0;i<=a;i++{
        fmt.Printf("%2d",i)
    for j:=0;j<=a;j++{
        fmt.Printf("%2d",i*j)      
    }
    fmt.Println()
  }
}
/*
$ go run function/test.go 1
 X 0 1
 0 0 0
 1 0 1
 $ go run function/test.go 2
 X 0 1 2
 0 0 0 0
 1 0 1 2
 2 0 2 4
$ go run function/test.go 3
 X 0 1 2 3
 0 0 0 0 0
 1 0 1 2 3
 2 0 2 4 6
 3 0 3 6 9
*/
