package main

import (
"bufio"
"fmt"
"log"
"os"
"strconv"
"strings"
"sort"
)

func main() {
    (_apriori("retail.dat", 2, 10000)) //data, passes, support
}

func _apriori(input_file string, pass int, support int) (map[pair]int)  {
    check := 0
    items := 0
    basket_total := 0

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
    }

    for check < pass { //check this after
        check++
        //frequent items
        for key, value := range temp_map {
            count := 0
            for count < (value){
                if temp_map[key] < support {
                    delete(temp_map, key)
                    count++
                } else {
                    count++
                }
            }
        }
        //frequent items end
    }
    // fmt.Println(temp_map)
    // fmt.Println(baskets[1])
    // r := sort.Ints(baskets[1])
    // fmt.Println(baskets[1], r)
    //fmt.Println("temp", temp_map)
    slice := _pairItems(temp_map)

    fmt.Println(slice, "returned")
    temp_map = nil //free memory yay

    var final_map map [pair]int 
    final_map = make(map[pair]int)

    //final_map = _match(slice, baskets, support)
    baskets = nil

    fmt.Println("items", items)
    fmt.Println("passes ", check)
    fmt.Println("threshold ", support)
    fmt.Println("baskets ", basket_total)
   // printPairs(PList)
    return final_map
}

type pair struct {
    item1 int
    item2 int
}

func _printMap(args map[pair]int)  {
    fmt.Println("print map")
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _pairItems(args map[int]int) (slice []int){
    keys := make([]int, 0)
    //final_pairs := make([]int, 0, 2)
    for key := range args {
        keys = append(keys, key)     
    }
    sort.Ints(keys)

    temp_keys := make([]int, len(keys))
    copy(temp_keys, keys)
    fmt.Println(keys)
    // length := len(temp_keys)
    // key_len := len(keys)

    // for i := 0; i < length; i++ {
    //     fmt.Println(temp_keys[i], keys[0], i)
    //     for x := 0; x < key_len; x++ {
    //         fmt.Println("before remove",keys)
    //         keys = keys[1:] //should take off the top
    //         fmt.Println("after remove",keys)
    //         fmt.Println(temp_keys[i], keys[x])
    //         key_len = key_len-1
    //     }
    //     length = length-1
    //     temp_keys = temp_keys[1:]
    //     fmt.Println("\t\t",i, temp_keys)
    // } 

    temp_keys = temp_keys[1:]
    fmt.Println("inital", keys, temp_keys)
    for key, _ := range keys {
        for key2, _ := range temp_keys {
            if keys[key] < temp_keys[key2]{
                fmt.Println(keys[0], temp_keys[key2])
            }
        }
        temp_keys = temp_keys[1:]
        keys = keys[1:]
        fmt.Println(keys, temp_keys)
    }
    return keys 
}

func _match(slice []int, baskets map[int][]int, support int) (map[pair]int){

    var final map [pair]int 
    final = make(map[pair]int)

    for key, array := range baskets {
        _ = key
        _ = array
    }


    for key := range final {
        if final[key] < support {
            delete(final, key)
        }
    }
    //fmt.Println("\n\nfinal", final[slice])

    _printMap(final)

    return final
}
