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
	(_apriori("retail.dat", 2, 881)) //data, passes, support
}

func _apriori(input_file string, pass int, support int) map[int]int {
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

	for check < pass { //check this after
		check++
		//frequent items
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
		//frequent items end
	}
	slice := _pairItems(temp_map)

	fmt.Println(slice, "returned")
	temp_map = nil //free memory yay

	var final_map map[int]int
	final_map = make(map[int]int)

	final_map = _match(slice, baskets, support)
	baskets = nil

	fmt.Println("items", items)
	fmt.Println("passes ", check)
	fmt.Println("threshold ", support)
	fmt.Println("baskets ", basket_total)
	return final_map
}

func _printMap(args map[int]int, pairs [][]int) {
	fmt.Println("print map")
	for key, value := range args {
		fmt.Println("Key:", pairs[key], "Value:", value)
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
			if len(temp_keys) > 0 {
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

	match := 0

	for _, array := range baskets { //itterates over all baskets
		for index, element := range pairs { //itterates over frequent pairs
			for x := range array {
				if array[x] == element[0] {
					for y := range array {
						if array[y] == element[1] {
							match++
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

	_printMap(final, pairs)
	return final
}
