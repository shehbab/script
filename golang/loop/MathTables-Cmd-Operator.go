package main
import ("fmt";"os";"strconv")
func main(){
  if len(os.Args) !=3 {
     fmt.Println("Usage: [op=*/+-] [size]")
      return
  } else { 
            
            o := os.Args[1]
            a, _ := strconv.Atoi(os.Args[2])
            fmt.Printf("Operator %s ,Multiple of %d",o,a)
            fmt.Println()            
            fmt.Printf("%4s",o)  
            for i:=0;i<=a;i++{
                fmt.Printf("%4d",i)
            }
            fmt.Println()
            for i:=0;i<=a;i++{
                    fmt.Printf("%4d",i)
                for j:=0;j<=a;j++{
                    if  o ==  "*" {
                        fmt.Printf("%4d",i*j)
                    } else if o == "+" {
                        fmt.Printf("%4d",i+j)    
                    } else if o == "-" {
                        fmt.Printf("%4d",i-j)
                    } else if o == "/" {
                        if j != 0 {
                        fmt.Printf("%4d",i/j)
                        } else {
                        fmt.Printf("%4d",j)   
                        }
                    }       
                }
                fmt.Println()
            }
            }
}
/**
$ go run function/test.go "*" 3
Operator * ,Multiple of 3
   *   0   1   2   3
   0   0   0   0   0
   1   0   1   2   3
   2   0   2   4   6
   3   0   3   6   9
$ go run function/test.go "/" 3
Operator / ,Multiple of 3
   /   0   1   2   3
   0   0   0   0   0
   1   0   1   0   0
   2   0   2   1   0
   3   0   3   1   1
**/
