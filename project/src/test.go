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


    _printMap(_apriori("retail.dat", 2, 10000)) //data, passes, support

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

    // pairs = _pairItems(temp_map)
    //PAIR START

    current := pair{}
    count := 0
    y := 0
    m := 0

    length := len(temp_map)
    slice := make([]int, length)
    tempSlice := slice

    PList := list{}
    PList.last = &current
    //fmt.Println("\tPLIST",PList)

    for key := range temp_map {
        slice[count] = key
        count++
    }

    for (y < len(tempSlice)){
        m = 0
        newPair := pair{}
        tempPair := pair{}
        for (m < len(tempSlice)){
            if (tempSlice[y] < tempSlice[m]){ // removes duplicates
                tempPair = current 
                PList.current = &newPair
                //PList.last = &tempPair             
                PList.current.prev = &tempPair
                fmt.Println("\t--- current ", *PList.current)//debug lines
                fmt.Println("\t\t----- current.prev", *PList.current.prev)
                fmt.Println("\t\t\t----- PList.last", *PList.last)
                fmt.Println("")

                current.item1 = tempSlice[y]
                current.item2 = tempSlice[m]

                PList.current = &current
            }
            m++
        }  
        y++
        printPairs(PList)       

    }



    //fmt.Println("should be full\n", current)
    temp_map = nil //free memory yay

    //PAIRS END
    //

    var final_map map [pair]int 
    final_map = make(map[pair]int)

    final_map = _match(PList, baskets, support)
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
    prev *pair
}

type list struct {
    current *pair
    last *pair
}

func printPairs(args list){
   //fmt.Println("\t\tpair print",*args.current)
    for args.last.prev != nil{
        args.current = args.last
        fmt.Println(*args.current)
    }
}


func _printMap(args map[pair]int)  {
    fmt.Println("print map")
    for key, value := range args {
        fmt.Println("Key:", key, "Value:", value)
    }
}

/*func _pairItems(args map[int]int) (x pair){

    current := pair{}

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
       // newPair := pair{}
        for (m < len(temp)){
            if (temp[y] < temp[m]){ // removes duplicates
                current.item1 = temp[y]
                current.item2 = temp[m] 

                temp := current 
                current.prev = &temp 
                //current.next = &newPair
                current = *current.next 
            }
            m++
        }  
        y++
    }
    fmt.Println("items paired sucessfully")
    return current 
    }*/

    func _match(args list, baskets map[int][]int, support int) (map[pair]int){
        count := 0
        count2 := 0

        var final map [pair]int 
        final = make(map[pair]int)

    // if (args.item1 == 0 && args.item2 == 0){
    //     args.current.end = true
    //     //fmt.Println(args.current)
    //     //fmt.Println("set pointer back 1")
    // }

    // for key, value := range baskets{
    //     fmt.Println(key, value)
    // }
        fmt.Println(args)
    // for (args.prev != nil){
    //     fmt.Println("matcher ",args)
    //     args = args.prev
    // }

        for key, array := range baskets {
            _ = key      
            count = 0
            for count < len(array){
                if args.current.item1 == array[count]{
                    count2 = 0
                    for count2 < len(array){
                        if args.current.item2 == array[count2]{
                            final[*(args.current)] = final[*(args.current)] + 1
                        }
                        count2++
                    }                
                }
                count++
            }
            if args.current != nil {
                args.current = args.last
            }
        }

        for key := range final {
            if final[key] < support {
                delete(final, key)
            }
        }
    //fmt.Println("\n\nfinal", final[args])
        _printMap(final)

        return final
    }