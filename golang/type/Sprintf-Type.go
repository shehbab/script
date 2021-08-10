package main
import ("fmt";"strconv")
func main() {
    floats := []float64{1.9999,2.4555}//slice
	for _, f := range floats {//index and element ( from slice in the loop)
        t := int(f)//convert float to int
        s := fmt.Sprintf("%.0f", f)//Round off float and print it as a string
        i, _ := strconv.Atoi(s)//convert string to int
	    fmt.Printf("t is of Type %T and value %v\n",t,t)
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
