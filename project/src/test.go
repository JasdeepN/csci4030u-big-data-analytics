package main

import (
"bufio"
"fmt"
"log"
"os"
"strconv"
"strings"

)

//Calls _apriori for k number of passes
func main() {
    // var mainMap map [pair]int 
    // mainMap = make(map[pair]int)
    // _ = mainMap

    basket_total := 0


    basket_total = _apriori("retail.dat", 2, 10000) //data, passes, support

    fmt.Println("total baskets: ", basket_total)
//    fmt.Println("Map: ", mainMap)
}

func _apriori(input_file string, pass int, support int) (int)  {
    check := 0
    items := 0
    basket_total := 0

    pairs := pair{}

    var temp_map map [int]int
    temp_map = make(map[int]int)

    var baskets map[int][]int
    baskets = make(map[int][]int)

    file, err := os.Open(input_file) 

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    if check == 0 {// load all items into the map one by one
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
        // fmt.Println(scanner.Text())
            count := 0
            basket_total = basket_total + 1 //total number of baskets
            words := strings.Fields(scanner.Text()) //seperating the strings
            items = items + len(words) //number of items
            var result = make([]int, len(words))
            _ = result
            for(count < len(words)){
                w, err := strconv.Atoi(words[count]); 
                if err != nil { 
                    fmt.Println("error parsing string") 
                }
                result[count] = w
                temp_map[result[count]] = temp_map[result[count]] + 1
                count++ 
            }
            baskets[basket_total] = result
        }
        check++

        fmt.Println("items", items)


        for check < pass {
            check++
    //frequent items
            for key, value := range temp_map {
                count := 0
                count2 := 0
                for count < (value){
                    if temp_map[key] < support {
                        delete(temp_map, key)
                        count++
                        count2++
                    } else {
                         count++
                        count2++
                    }
                }
            }
            fmt.Println("map")
            _printMap(temp_map)
           // _printMap2(baskets)
        }
    }

    //done making frequent itmes
    _ = pairs

    // pairs = _pairItems(temp_map)
    // temp_map = nil //free memory yay

    // var final_map map [pair]int 
    // final_map = make(map[pair]int)

    // final_map = _match(pairs, baskets)

    fmt.Println("passes ", check)
    fmt.Println("threshold ", support)

    return int(basket_total)
}

type pair struct {
    item1 int
    item2 int
}

type list struct{
    prev *pair
    next *pair
    last bool
}


func _printMap(args map[int]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _printMap2(args map[int][]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}
