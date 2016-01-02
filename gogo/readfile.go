package main

import (
        "fmt"
        "os"
        "bufio"
)

func main(){
    input_file , input_error := os.Open(os.Args[1]) 
    if input_error != nil {
        fmt.Printf("An error occurred on opening the inputfile \n")
        return 
    }
    reader := bufio.NewReader(input_file)
    for { 
        inputString , readerError := reader.ReadString('\n')  
        if readerError != nil {
                break
        }

        fmt.Printf("%s" , inputString)
    }
}
