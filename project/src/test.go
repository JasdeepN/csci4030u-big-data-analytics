package main

import (
"bufio"
"fmt"
"log"
"os"
"strconv"
"strings"

)

func main() {
    // var mainMap map [pair]int 
    // mainMap = make(map[pair]int)
    // _ = mainMap

    basket_total := 0


    _printMap2(_apriori("small_data.dat", 2, 2)) //data, passes, support

    fmt.Println("total baskets: ", basket_total)
    //_printMap2("Map: ", mainMap)
}

func _apriori(input_file string, pass int, support int) (map[pair]int)  {
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
            //_printMap(temp_map)
           // _printMap2(baskets)
        }
    }

    //done making frequent itmes
    _ = pairs

    pairs = _pairItems(temp_map)

    temp_map = nil //free memory yay

    var final_map map [pair]int 
    final_map = make(map[pair]int)

    final_map = _match(pairs, baskets)

    baskets = nil

    fmt.Println("passes ", check)
    fmt.Println("threshold ", support)

    return final_map
}

type pair struct {
    prev *pair
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

func _printMap2(args map[pair]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _pairItems(args map[int]int) (x pair){

    current := pair{}
    newPair := pair{}
    count := 0
    y := 0
    m := 0

    length := len(args)
    slice := make([]int, length)
    temp := slice


    for key := range args {
        slice[count] = key
        count++
    }

    for (y < len(temp)){
        m = 0
        for (m < len(temp)){
            if (temp[y] < temp[m]){ // removes duplicates
                current.item1 = temp[y]
                current.item2 = temp[m]
                if m != len(temp) {
                    tempNode := current
                    current = newPair
                    current.prev = &tempNode  
                }
            }
            m++
        }   
        y++
    }
    fmt.Println("items paired sucessfully")
    return current 
}

func _match(args pair, baskets map[int][]int, support int) (map[pair]int){
    count := 0
    count2 := 0

    var final map [pair]int 
    final = make(map[pair]int)

    // if (args.item1 == 0 && args.item2 == 0){
    //     args = *args.prev
    //     fmt.Println("set pointer back 1")
    // }

    for key, array := range baskets {
        _ = key
        
        count = 0
        for count < len(array){
            if args.item1 == array[count]{
                count2 = 0
                for count2 < len(array){
                    if args.item2 == array[count2]{
                        final[args] = final[args] + 1
                    }
                    count2++
                }
            }
            count++
            if args.prev != nil {
                args = *args.prev
            }
        }
                fmt.Println(final[args])
    }
    return final
}