package main 

import "fmt"
import "bufio"
import "os"


type CommandType struct {
  name string 
  descp string 
  callback func() error 
}

const cmds := map[string]CommandType{
  "help": {
    name: "help",
    descp: "displays a help message",
    callback: cmdHelp,
  },
  "exit": {
    name: "exit",
    descp: "Exit the Pokedex",
    callback: cmdExit,
  }
}



func main(){
  scanner := bufio.NewScanner(os.Stdin)

  for ;; {
  
  }
}
