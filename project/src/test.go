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


    (_apriori("retail.dat", 2, 10000)) //data, passes, support

}

func _apriori(input_file string, pass int, support int) (map[pair]int)  {
    check := 0
    items := 0
    basket_total := 0

    //pairs := pair{}

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
    }
    
    //done making frequent itmes
    slice := make([]pair, len(temp_map))
    slice = _pairItems(temp_map)

    fmt.Println(slice, "END")
    temp_map = nil //free memory yay

    //PAIRS END
    //

    var final_map map [pair]int 
    final_map = make(map[pair]int)

    //final_map = _match(PList, baskets, support)
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

func _pairItems(args map[int]int) (slice []pair){

    current := pair{}
    count := 0
    x := 0
    slice = make([]pair, len(args))
    for key := range args {
        x = 0
        for key2 := range args {
            if key < key2 {
                current.item1 = key
                current.item2 = key2
            }
        } 

        for x < len(slice) { 
            fmt.Println(slice[x].item1, current.item1, slice[x].item2, current.item2)
            if slice[x].item1 == current.item1 && slice[x].item2 == current.item1{
                //break
                //do nothing
                delete(key, slice)
            } else {
                slice[count] = current
                fmt.Println("add",slice[count], count)
            }
            x++
        }
        count++
    }
    fmt.Println("items paired sucessfully")

    for x < len(slice){
        fmt.Println(x, slice[x])

    }
    return slice 
}

/*func _match(slice []pair, baskets map[int][]int, support int) (map[pair]int){
    count := 0
    count2 := 0

    var final map [pair]int 
    final = make(map[pair]int)

    // if (slice.item1 == 0 && slice.item2 == 0){
    //     slice.current.end = true
    //     //fmt.Println(slice.current)
    //     //fmt.Println("set pointer back 1")
    // }

    // for key, value := range baskets{
    //     fmt.Println(key, value)
    // }
    fmt.Println(slice)
    // for (slice.prev != nil){
    //     fmt.Println("matcher ",slice)
    //     slice = slice.prev
    // }

    for key, array := range baskets {
        _ = key      
        count = 0
        for count < len(array){
            if slice.current.item1 == array[count]{
                count2 = 0
                for count2 < len(array){
                    if slice.current.item2 == array[count2]{
                        final[*(slice.current)] = final[*(slice.current)] + 1
                    }
                    count2++
                }                
            }
            count++
        }
        if slice.current != nil {
            slice.current = slice.last
        }
    }

    for key := range final {
        if final[key] < support {
            delete(final, key)
        }
    }
    //fmt.Println("\n\nfinal", final[slice])
    _printMap(final)

    return final
}*/