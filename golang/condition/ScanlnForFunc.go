package main
import ("fmt")
func multiple(i int) {
    c := 1
    for c<=10 {
	fmt.Printf("%d * %d = %d\n",i,c,c*i)
	c++
	}
}
func main(){
    var number int
    fmt.Print("Enter the table: ")
    fmt.Scanln(&number)
    multiple(number)
}

/***
$ go run function/test.go 
Enter the table: 20
20 * 1 = 20
20 * 2 = 40
20 * 3 = 60
20 * 4 = 80
20 * 5 = 100
20 * 6 = 120
20 * 7 = 140
20 * 8 = 160
20 * 9 = 180
20 * 10 = 200
***/
