package main
import ("fmt")
func multiple(max int) {
   fmt.Printf("%5s","X") //X
   //First now
   for i:=0;i<=max;i++{
	   fmt.Printf("%5d",i)
   }
   fmt.Println()
   for i:=0;i<=max;i++{
	
	fmt.Printf("%5d",i)
	
	for j:=0;j<=max;j++{
		fmt.Printf("%5d",i*j)
	}
	fmt.Println()

   }
   
}
func main(){
    var number int
    fmt.Print("Enter the table: ")
    fmt.Scanln(&number)
    multiple(number)
}

/*
$ go run function/test.go 
Enter the table: 5
    X    0    1    2    3    4    5
    0    0    0    0    0    0    0
    1    0    1    2    3    4    5
    2    0    2    4    6    8   10
    3    0    3    6    9   12   15
    4    0    4    8   12   16   20
    5    0    5   10   15   20   25
*/
