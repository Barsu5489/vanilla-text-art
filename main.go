package main

import "fmt"

func main (){
	for i:= 0; i < 8; i++{
        if i == 1 || i == 5{
           fmt.Print("-") 
        }else{
            fmt.Print(" ")
        }
         
    }
    fmt.Println()
     for i := 0; i<= 4; i++{
         
         if i < 4 && i != 2 && i != 1{
                 for i:= 0; i <= 1; i++{
     fmt.Print("| | ")
     
    }
         }else if i == 4{
               fmt.Print("|_| |_|")
         } else if i == 1{
             fmt.Print("| |_| |")
         }else {
             fmt.Print("|  _  |")
         }
     

     fmt.Println()
		}
	}