package main
import ("fmt";"os";"strings")
func main(){
  words := strings.Fields(
      "This is a cool document read it on time.",
  )
  for i,v := range words{
      fmt.Printf(">%-5d  --- %q\n",i+1,v)
  }
  //Loop Using range (index,value)
  for i, v := range os.Args {
      if i == 0 {
          continue
      }
      fmt.Printf("%q\n",v)
  }
  for _, v := range os.Args[1:] {
    fmt.Printf("%q\n",v)
}

}
/*
go run function/test.go Shehbab Kakkar Good
>1      --- "This"
>2      --- "is"
>3      --- "a"
>4      --- "cool"
>5      --- "document"
>6      --- "read"
>7      --- "it"
>8      --- "on"
>9      --- "time."
"Shehbab"
"Kakkar"
"Good"
"Shehbab"
"Kakkar"
"Good"*/
