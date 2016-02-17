package main

import (
"bufio"
"fmt"
"log"
"os"
"sort"
"strconv"
"strings"
)

func main() {
    //calls _printMap
	_apriori("retail.dat", 2, 3000) //data, passes, support
}

func _apriori(input_file string, pass int, support int) (map[int]int){
	check := 0
	items := 0
	basket_total := 0

	var temp_map map[int]int
	temp_map = make(map[int]int)

	var baskets map[int][]int
	baskets = make(map[int][]int)

    //open file for reading
	file, err := os.Open(input_file)

	if err != nil {
		log.Fatal(err)
	}

    /*FIRST PASS*/

	defer file.Close()
	if check == 0 { // load all items into the map one by one (first pass)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			count := 0
			basket_total = basket_total + 1         //total number of baskets
			words := strings.Fields(scanner.Text()) //seperating the strings
			items = items + len(words)              //number of items
			var result = make([]int, len(words))
			_ = result
			for count < len(words) {
				w, err := strconv.Atoi(words[count])
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

    /*FIRST PASS END*/

    /*MAKING FREQUENT PAIRS*/

	for check < pass { //NUMBER OF PASSES (need this? prob not in theory could do more passes with some changes)
		check++
		for key, value := range temp_map {
			count := 0
			for count < (value) {
				if temp_map[key] < support {
					delete(temp_map, key)
					count++
				} else {
					count++
				}
			}
		}
	}

    /*END MAKING FREQUENT PAIRS*/
	
    //pair items
	slice := _pairItems(temp_map)

    //dont need item counts anymore garbage collection should take over
	temp_map = nil 

    //final results
	var final_map map[int]int
	final_map = make(map[int]int)

    //matching frequent items in baskets
	final_map = _match(slice, baskets, support)

    //dont need the baskets anymore
	baskets = nil

    //checks
	/*fmt.Println("items", items)
	fmt.Println("passes ", check)
	fmt.Println("threshold ", support)
	fmt.Println("baskets ", basket_total)*/
    _printMap(final_map, slice)
	return final_map
}


func _printMap(args map[int]int, pairs [][]int) {
	fmt.Println("print map")
	for key, value := range args {
		fmt.Println("Key:", pairs[key], "Value:", value)
	}
}

func _pairItems(args map[int]int) (slice [][]int) {
    keys := make([]int, 0)

    final_pairs := make([][]int, 0)
    for key := range args {
        keys = append(keys, key)
    }
    sort.Ints(keys)

    temp_keys := make([]int, len(keys))
    copy(temp_keys, keys)

    if (len(temp_keys)>0){
       temp_keys = temp_keys[1:]
   }

   for key, _ := range keys {
    if (len(temp_keys)>1){
        for key2, _ := range temp_keys {
            temp := make([]int, 2)
            if keys[key] < temp_keys[key2] {
                temp[0] = keys[0]
                temp[1] = temp_keys[key2]
                final_pairs = append(final_pairs, temp)
                temp = nil
            } else {
                temp[0] = temp_keys[key2]
                temp[1] = keys[0]
                final_pairs = append(final_pairs, temp)
                temp = nil
            }
        }
    } else {
        if (len(temp_keys) > 0){
            temp := make([]int, 2)
            temp[0] = keys[0]
            temp[1] = temp_keys[0]
            final_pairs = append(final_pairs, temp)
            temp = nil
        }
            break; //reached the end
        }

        temp_keys = temp_keys[1:]
        keys = keys[1:]
    }
    return final_pairs
}

func _match(pairs [][]int, baskets map[int][]int, support int) map[int]int {
	var final map[int]int
	final = make(map[int]int)

    match := 0

    for _, array := range baskets { //itterates over all baskets
        for index,element := range pairs {  //itterates over frequent pairs
            for x := range array {
                if (array[x] == element[0]){
                    for y := range array {
                        if (array[y] == element[1]){
                            match++
                            final[index] = final[index]+1
                        }
                    }
                }
            }
        }
    }
    for key := range final {
      if final[key] < support {
        delete(final, key)
    }
}
return final
}

