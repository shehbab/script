package main
import ("fmt")
func condition(){
   var ( 
        sum, counter int  //Counter start
       )    
       counter = 1
     fmt.Println("Initial Sum value:",sum)
    for {
        //Condition 
        if counter > 5 {
            break
        }
		
        if counter%2 == 0 {
			counter++
            continue 
        }
        sum += counter
        fmt.Printf("%d Final Sum value: %d\n",counter,sum)
        counter++ // incremental
    }   
    fmt.Println("Total Sum value:",sum)
        
}


func main(){
    condition()
}

/*
$ go run function/test.go
Initial Sum value: 0
1 Final Sum value: 1
3 Final Sum value: 4
5 Final Sum value: 9
Total Sum value: 9*/
