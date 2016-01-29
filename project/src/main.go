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

    first_pass, basket_total = _readData("small_data.dat", )

    fmt.Println("total baskets: ", basket_total)
    //fmt.Println("Map: ", first_pass)
    
    _printMap(first_pass)

}

func _printMap(input map[string]int)  {
    for key, value := range input {
        fmt.Println("Key:", key, "Value:", value)
    }

}

func _readData(input_file string) (map[string]int, int)  {
    var inital_map map [string]int
    inital_map = make(map[string]int)
    basket_total := 0


    file, err := os.Open(input_file) //101 items only
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
            inital_map[words[count]] = inital_map[words[count]] + 1
            count++ 
        }

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return inital_map, basket_total
}


