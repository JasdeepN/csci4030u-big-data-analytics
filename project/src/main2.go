package main

import (
"bufio"
"fmt"
"log"
"os"
"strings"
"strconv"
// "container/list"

)

func main() {
    var currMap map [string]int 
    currMap = make(map[string]int)
    _ = currMap

    // var oldMap map [string]int 
    // oldMap = make(map[string]int)
    // _ = oldMap

    basket_total := 0
    k := 2 // number of passes
    s := 0.01 * 1000.0; //threshold - items that appear >= k = frequent items

    _ = s

    fmt.Println("s = ", s)
   //_ = k // fixes the "declared but not used err"

    currMap, basket_total = _apriori("medium_data.dat", k, s) //data, passes, support


    fmt.Println("total baskets: ", basket_total)
    //fmt.Println("Map: ", currMap)
}

func _printMap(args map[string]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _apriori(input_file string, pass int, threshold float64) (map[string]int, int)  {
    check := 0
    records := 0

    var inital_map map [string]int
    inital_map = make(map[string]int)
    basket_total := 0

    file, err := os.Open(input_file) 

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if check == 0{// load all items into the map one by one
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
        // fmt.Println(scanner.Text())
            count := 0
            basket_total = basket_total + 1 //total number of baskets
            words := strings.Fields(scanner.Text()) //seperating the strings
            records = records + len(words) //number of records

            for(count < len(words)){
                inital_map[words[count]] = inital_map[words[count]] + 1
                count++ 
                //fmt.Println("count < words[count] ",  inital_map[words[count]])
            }
        }
       // _printMap(inital_map)

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
        check++
    }
    var newMap map [string]int 
    newMap = make(map[string]int)
    for check < pass {
        check++
        count := 0

    //frequent items
        for key, value := range inital_map {
            if float64(value) < threshold {
                delete(inital_map, key)
            }
        }
        _printMap(inital_map)

    //making old item map
        for key := range inital_map {
            count++

            n, err := strconv.Atoi(key); 
            if err != nil { 
                fmt.Println("error parsing string") 
            } 
            newMap[strconv.Itoa(count)] = n
        }
        fmt.Println("check ", check)
    }
    fmt.Println("records ", records)

    // fmt.Println("oldMap")
    // _printMap(inital_map)

    // fmt.Println("currMap")
    // _printMap(newMap)
    return newMap, basket_total
}