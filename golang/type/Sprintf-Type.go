package main
import ("fmt";"strconv")
func main() {
    floats := []float64{1.9999,2.4555}
    for _, f := range floats {
        t := int(f)
        s := fmt.Sprintf("%.0f", f)
        i, _ := strconv.Atoi(s)
        fmt.Printf("Type %T and value %v\n",t,t)
		fmt.Printf("s is of Type %T and value %s\n",s,s)
		fmt.Printf("i is of Type %T and value %v\n",i,i)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>")
    }
}


/*
Type int and value 1
s is of Type string and value 2.0
i is of Type int and value 0
>>>>>>>>>>>>>>>>>>>>
Type int and value 2
s is of Type string and value 2.5
i is of Type int and value 0
>>>>>>>>>>>>>>>>>>>>
*/
