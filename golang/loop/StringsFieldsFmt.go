package main
import ("fmt";"os";"strings")
func main(){
  //Loop String from CmdLine  
  for i:=1;i<len(os.Args);i++{
      fmt.Printf("%q\n",os.Args[i])
  }
  words := strings.Fields(
      "This is a cool document read it on time.",
  )
  //Loop Sentence from field 
  for j:=0;j<len(words);j++{
      fmt.Printf("#%-10d: %10q\n",j+1,words[j])
  }
}

/*
$ go run function/test.go Shehbab Kakkar Good
"Shehbab"
"Kakkar"
"Good"
#1         :     "This"
#2         :       "is"
#3         :        "a"
#4         :     "cool"
#5         : "document"
#6         :     "read"
#7         :       "it"
#8         :       "on"
#9         :    "time."
*/
