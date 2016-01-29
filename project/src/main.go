package main

import (
"bufio"
"fmt"
"log"
"os"
"strings"
)

func main() {

    var first_pass map [string]int 
    first_pass = make(map[string]int)
  //  _ = first_pass

    basket_total := 0
    threshold := 2 // items that appear >= threshold = frequent items

    _ = threshold // fixes the "declared but not used err"

   file, err := os.Open("small_data.dat") //101 items only
    //file, err := os.Open("retail.dat") //full data

   if err != nil {
    log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        count := 0
        basket_total = basket_total + 1
        words := strings.Fields(scanner.Text())

        for(count < len(words)){
            first_pass[words[count]] = first_pass[words[count]] + 1
            count++ 
        }
       
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("total baskets: ", basket_total)
    //fmt.Println("Map: ", first_pass)
    
    _printMap(first_pass)
   
}

func _printMap(input map[string]int)  {
     for key, value := range input {
        fmt.Println("Key:", key, "Value:", value)
    }
}
