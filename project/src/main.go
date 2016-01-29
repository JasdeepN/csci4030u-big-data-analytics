package main

import (
"bufio"
"fmt"
"log"
"os"
"strings"
"strconv"
"container/list"

)

func main() {
    var currMap map [string]int 
    currMap = make(map[string]int)

    var oldMap map [string]int 
    oldMap = make(map[string]int)
    // _ = oldMap

    basket_total := 0
    threshold := 2 // items that appear >= threshold = frequent items

   //_ = threshold // fixes the "declared but not used err"

    currMap, basket_total = _passOne("small_data.dat")

   // _printMap(currMap)

    oldMap, currMap = _passTwo(currMap, threshold)

    fmt.Println("oldMap")
    _printMap(oldMap)

    fmt.Println("currMap")
    _printMap(currMap)

    fmt.Println("total baskets: ", basket_total)
    //fmt.Println("Map: ", currMap)
}

func _printMap(args map[string]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _passOne(input_file string) (map[string]int, int)  {
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

func _passTwo(args map[string]int, threshold int) (map[string]int, map[string]int) {
    var newMap map [string]int 
    newMap = make(map[string]int)
    count := 0

    //frequent items
    for key, value := range args {
        if value < threshold {
            delete(args, key)
        }
    }

    //making old item map
    for key := range args {
        count++

        n, err := strconv.Atoi(key); 
        if err != nil { 
            fmt.Println("error parsing string"f) 
        } 
        newMap[strconv.Itoa(count)] = n
    }

    return args, newMap
}

func _pairItems(args map[string]int) (map[string]int){
    tempList := list.new()  
    for key := range args {
        list.PushBack(new pair(key, args[key+1]))//FIX HERE
    }
    return args
}

type pair struct {
    item1 string
    item2 string
}