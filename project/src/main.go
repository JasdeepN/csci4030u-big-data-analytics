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
    var mainMap map [string]int 
    mainMap = make(map[string]int)
    _ = mainMap

    basket_total := 0
    k := 2 // number of passes
   //_ = k // fixes the "declared but not used err"

    mainMap, basket_total = _apriori("retail.dat", k) //data, passes

    fmt.Println("total baskets: ", basket_total)
    //fmt.Println("Map: ")
    //_printMap(mainMap)
}

func _printMap(args map[string]int)  {
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

func _apriori(input_file string, pass int) (map[string]int, int)  {
    check := 0
    items := 0
    threshold := 0.0
    basket_total := 0.0
    s := 0.01
    _ = s

    pairs := pair{}

    var inital_map map [string]int
    inital_map = make(map[string]int)
    _ = inital_map

    var temp_map map [string]int
    temp_map = make(map[string]int)

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
                temp_map[words[count]] = temp_map[words[count]] + 1
                count++ 
                //fmt.Println("count < words[count] ",  inital_map[words[count]])
            }
        }
        _copyMap(temp_map, inital_map) //copy before removing items
        // _printMap(inital_map)
        // fmt.Println("inital_map")

        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
        check++
    }

    //threshold = s * basket_total //881
    threshold = 10000 //test
    
    var oldItems map [string]int 
    oldItems = make(map[string]int)

    for check < pass {
        check++
        count := 0

    //frequent items
        for key, value := range temp_map {
            if float64(value) < threshold {
                delete(temp_map, key)
            }
        }
        pairs = _pairItems(temp_map)
       // _printMap(temp_map)
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
        fmt.Println("passes ", check)
    }
    fmt.Println("threshold ", threshold)
    _printPairs(pairs)

    return oldItems, int(basket_total)

}

func _copyMap(args map[string]int, newMap map[string]int) {
    for k,v := range args {
        newMap[k] = v
    }
}

func _printPairs(x pair){
    current := pair{}

    if (x.prev != nil){
        fmt.Println("(", current.item1, current.item2, ")")
        current = *x.prev

    } //move pointer back 1 so we dont get {0, 0}

    for(current.prev != nil){
        if current.item1 != 0 && current.item2 != 0 {
            fmt.Println("(", current.item1, current.item2, ")")
            current = *current.prev
        }

    }

    if (current.prev == nil) {
        fmt.Println("(", current.item1, current.item2, ")")
    }
}

func _pairCheck(n int, w int, args pair) (bool){
    if (args.prev != nil){
        args = *args.prev
    } 
    for(args.prev != nil){
       // fmt.item1
        if args.item1 == n && args.item2 == w {
            fmt.Println("duplicate")
            return false
        } else {
            args = *args.prev

        }
    }
    return true
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
    fmt.Println(slice)

    y = 0
    m := 0

    for (y < len(slice)){
        m = 0
        for (m < len(slice)){
            if (slice[y] != slice[m]){ // removes duplicates
                if (slice[y] != slice[m]) {
               
                    current.item1 = slice[y]
                    current.item2 = slice[m]
                   // fmt.Println("---", current.item1, current.item2, y, m)
                   // fmt.Println(current.item1 == slice[y] && current.item2 == slice[m])
                    //fmt.Println("different")

                    tempNode := current
                    current = newPair
                    current.prev = &tempNode
                } 
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

