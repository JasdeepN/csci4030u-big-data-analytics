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
	args := os.Args
	//argsWithoutProg := os.Args[1:]

	printflag, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	support, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//fmt.Println("EXECUTING", args[0])
	(_apriori(args[1], 2, support, printflag)) //data, passes, support, printflag
}

func _apriori(input_file string, pass int, support float64, printflag int) map[int]int {
	check := 0
	items := 0
	basket_total := 0

	var temp_map map[int]int
	temp_map = make(map[int]int)

	var baskets map[int][]int
	baskets = make(map[int][]int)

	file, err := os.Open(input_file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	if check == 0 { // load all items into the map one by one
		scanner := bufio.NewScanner(file)
		var bufArr []byte
		scanner.Buffer(bufArr, 262144)
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

	// var temp float64
	temp := (support * float64(basket_total))
	threshold := int(temp)

	for check < pass { //check this after
		check++
		//frequent items
		for key, value := range temp_map {
			count := 0
			for count < (value) {
				if temp_map[key] < threshold {
					delete(temp_map, key)
					count++
				} else {
					count++
				}
			}
		}
		//frequent items end
	}
	slice := _pairItems(temp_map)

	//fmt.Println(slice, "returned")
	temp_map = nil //free memory yay

	var final_map map[int]int
	final_map = make(map[int]int)

	final_map = _match(slice, baskets, threshold)
	baskets = nil

	/*	fmt.Println("items", items)
		fmt.Println("passes ", check)
		fmt.Println("threshold ", threshold)
		fmt.Println("baskets ", basket_total, "\n")*/
	if printflag == 1 {
		_printMap(final_map, slice)
		_printStats(items, basket_total, threshold, check, input_file)
	}
	return final_map
}

func _printStats(items int, numBaskets int, support int, passes int, input string) {
	fmt.Println("+---------------------------------------+")
	fmt.Println("|\t\tSTATISTICS\t\t|")
	fmt.Println("+---------------------------------------+")
	fmt.Println("|\t\t", input, "\t\t|")
	fmt.Println("|\t      ITEMS:", items, "\t\t|")
	fmt.Println("|\t        PASSES:", passes, "\t\t|")
	fmt.Println("|\t     THRESHOLD:", support, "\t\t|")
	fmt.Println("|\t      BASKETS:", numBaskets, "\t\t|")
	fmt.Println("+---------------------------------------+")

}

func _printMap(args map[int]int, pairs [][]int) {
	fmt.Println("+---------------------------------------+")
	fmt.Println("\t\tPRINTING MAP\t\t")
	fmt.Println("+---------------------------------------+")
	for key, value := range args {
		fmt.Println("\tKey:", pairs[key], "Value:", value, "\t")
	}

}

func _pairItems(args map[int]int) (slice [][]int) {
	var keys []int

	var final_pairs [][]int

	for key := range args {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	temp_keys := make([]int, len(keys))
	copy(temp_keys, keys)

	if len(temp_keys) > 0 {
		temp_keys = temp_keys[1:]
	}

	for key := range keys {
		_ = key
		if len(temp_keys) > 1 {
			for key2 := range temp_keys {
				temp := make([]int, 2)
				temp[0] = keys[0]
				temp[1] = temp_keys[key2]

				final_pairs = append(final_pairs, temp)
				temp = nil
			}
		} else {
			if len(temp_keys) == 1 {
				temp := make([]int, 2)
				temp[0] = keys[0]
				temp[1] = temp_keys[0]
				final_pairs = append(final_pairs, temp)
				temp = nil
			}
			break //reached the end
		}

		temp_keys = temp_keys[1:]
		keys = keys[1:]
	}
	return final_pairs
}

func _match(pairs [][]int, baskets map[int][]int, support int) map[int]int {

	var final map[int]int
	final = make(map[int]int)

	for _, array := range baskets { //itterates over all baskets
		for index, element := range pairs { //itterates over frequent pairs
			for x := range array {
				if array[x] == element[0] {
					for y := range array {
						if array[y] == element[1] {
							final[index] = final[index] + 1
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

	//_printMap(final, pairs)
	return final
}
