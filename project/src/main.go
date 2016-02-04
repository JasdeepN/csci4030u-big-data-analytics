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
    var mainMap map [pair]int 
    mainMap = make(map[pair]int)
    _ = mainMap

    basket_total := 0
    k := 2 // number of passes
   //_ = k // fixes the "declared but not used err"

    mainMap, basket_total = _apriori("retail.dat", k) //data, passes

    fmt.Println("total baskets: ", basket_total)
    //fmt.Println("Map: ")
    //_printMapPair(mainMap)
}

func _printMap(args map[string]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _printMapPair(args map[pair]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _apriori(input_file string, pass int) (map[pair]int, int)  {
    check := 0
    items := 0
    threshold := 0.0
    basket_total := 0.0

    pairs := pair{}

    var temp_map map [string]int
    temp_map = make(map[string]int)

    var baskets map[string]int
    baskets = make(map[string]int)

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
            items = items + len(words) //number of items

            for(count < len(words)){
                baskets[words[count]] = baskets[words[count]] + 1
                count++ 
            }
        }
        _copyMap(baskets, temp_map)
        //_printMap(temp_map)
        //fmt.Println("temp_map")

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
        check++
    }

    //threshold = 0.01 * basket_total //881
    threshold = 10000   //test
    
    var oldItems map [string]int 
    oldItems = make(map[string]int)

    var final_map map [pair]int 
    final_map = make(map[pair]int)

    for check < pass {
        check++
        count := 0

    //frequent items
        for key, value := range temp_map {
            if float64(value) < threshold {
                delete(temp_map, key)
            }
        }

    //done making frequent itmes
        _ = pairs
    //making old item map
        for key := range temp_map {
            count++
            n, err := strconv.Atoi(key); 
            if err != nil { 
                fmt.Println("error parsing string") 
            } 
            oldItems[strconv.Itoa(count)] = n
        }
    }

    pairs = _pairItems(temp_map)
    pairs = _match(pairs, baskets)

    fmt.Println("passes ", check)
    fmt.Println("threshold ", threshold)
    //_printPairs(pairs)
    //_printMapPair(final_map)
    return final_map, int(basket_total)

}

func _match(args pair, baskets map[string]int) (result pair){
    for key, value := range baskets {
        fmt.Println("key", key, value)
    }
    return args
}

func _copyMap(args map[string]int, newMap map[string]int) {
    for k,v := range args {
        newMap[k] = v
    }
}

func _printPairs(args pair){
    if args.item1 == 0 && args.item2 == 0 {

       // fmt.Println("(", args.item1, args.item2, ")")
        args = *args.prev
    //move pointer back 1 so we dont get {0, 0}

        for(args.prev != nil){
            if args.item1 != 0 && args.item2 != 0 {
                fmt.Println("(", args.item1, args.item2, ")")
                args = *args.prev
            }
        }

        if (args.prev == nil) {
            fmt.Println("(", args.item1, args.item2, ")")
        }
    }
}

func _pairItems(args map[string]int) (x pair){
    current := pair{}
    newPair := pair{}
    count := 0
    y := 0
    length := len(args)
    fmt.Println(length)
    slice := make([]int, length)
    w := 0
    _ = w

    for key := range args {
        w, err := strconv.Atoi(key); 
        if err != nil { 
            fmt.Println("error parsing string") 
        }
        slice[count] = w
        //fmt.Println(slice[y])
        count++
    }
       //for key2 := range args {
    //fmt.Println(slice)

    y = 0
    m := 0

    temp := slice

    for (y < len(temp)){
        m = 0
        for (m < len(temp)){
            if (temp[y] < temp[m]){ // removes duplicates
                //fmt.Println((temp[y] != slice[m]), temp[y], temp[m])

                current.item1 = temp[y]
                current.item2 = temp[m]

                tempNode := current
                current = newPair
                current.prev = &tempNode  
            }
            m++
        }   
        y++
    }
   //_printPairs(current)
    return current 
}

type pair struct {
    prev *pair
    item1 int
    item2 int
}

