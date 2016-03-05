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

	buckets, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	printflag, err := strconv.Atoi(args[4])
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
	//(_pcy("../apriori/netflix.data", 2, 0.25, 1))
	(_pcy(args[1], 2, support, buckets, printflag))

}

func _pcy(input_file string, pass int, support float64, buckets int, printflag int) map[int]int {
	check := 0
	items := 0
	basket_total := 0

	var temp_map map[int]int
	temp_map = make(map[int]int)

	var baskets map[int][]int
	baskets = make(map[int][]int)

	var hashMap map[int]int
	hashMap = make(map[int]int)

	file, err := os.Open(input_file)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("reading file")
	defer file.Close()
	if check == 0 { // load all items into the map one by one
		scanner := bufio.NewScanner(file)
		var bufArr []byte
		scanner.Buffer(bufArr, 262144)
		for scanner.Scan() {
			count := 0
			count2 := 0
			var tempHash1 int
			var tempHash2 int
			basket_total = basket_total + 1         //total number of baskets
			words := strings.Fields(scanner.Text()) //seperating the strings
			items = items + len(words)              //number of items
			var result = make([]int, len(words))
			//_ = result
			for count < len(words) {
				w, err := strconv.Atoi(words[count])
				if err != nil {
					fmt.Println("error parsing string")
				}
				result[count] = w
				temp_map[result[count]] = temp_map[result[count]] + 1
				count++
				if count2 == 0 {
					tempHash1 = w
				} else if count2 == 1 {
					tempHash2 = w
					temp := hash(tempHash1, tempHash2, buckets)
					hashMap[temp] = hashMap[temp] + 1
					count2 = 0
				}
				count2++
			}
			baskets[basket_total] = result
		}
		check++
	}
	// fmt.Println("reading file complete")

	temp := (support * float64(basket_total))
	threshold := int(temp)

	// fmt.Println("starting freqent items")
	//frequent items
	for key := range temp_map {
		if temp_map[key] < threshold {
			delete(temp_map, key)
		}
	}
	for x := range hashMap {
		if hashMap[x] > 100 { //10 backup value
			// fmt.Println(x, "Working", hashMap[x])
			hashMap[x] = 1
		} else {
			// fmt.Println(x, "skip", hashMap[x])
			hashMap[x] = 0
		}
	}
	// fmt.Println("done freqent items")

	//frequent items end

	// fmt.Println("temp", temp_map)
	// fmt.Println("hash", hashMap)

	//frequent items end
	// fmt.Println("pairing items")

	slice := _pairItems(temp_map, hashMap, threshold)
	// fmt.Println("pairing complete")

	//fmt.Println(slice, "returned")

	var final_map map[int]int
	final_map = make(map[int]int)

	// fmt.Println("starting matching")

	final_map = _match(slice, threshold, baskets, hashMap, temp_map, buckets)
	// fmt.Println("finished matching")

	// fmt.Println("freeing memory")
	temp_map = nil //free memory yay
	baskets = nil

	if printflag == 1 {
		_printMap(final_map, slice)
		_printStats(items, basket_total, threshold, check, input_file, buckets)
	}
	return final_map
}

func hash(x int, y int, buckets int) int {
	// var keys []int
	//    for k := range hashMap {
	//        keys = append(keys, k)
	//    }
	//    sort.Ints(keys)

	var temp = (x ^ y) % buckets
	// To perform the opertion you want
	// for _, k := range keys {
	//     fmt.Println("Key:", k, "Value:", m[k])
	// }
	return temp

}

func _printStats(items int, numBaskets int, support int, passes int, input string, buckets int) {
	fmt.Println("+---------------------------------------+")
	fmt.Println("|\t\tSTATISTICS\t\t|")
	fmt.Println("+---------------------------------------+")
	fmt.Println("|\t\t", input, "\t\t|")
	fmt.Println("|\t      ITEMS:", items, "\t\t|")
	fmt.Println("|\t      BUCKETS:", buckets, "\t\t|")
	fmt.Println("|\t     THRESHOLD:", support, "\t\t|")
	fmt.Println("|\t      BASKETS:", numBaskets, "\t\t|")
	fmt.Println("+---------------------------------------+")

}

func _pairItems(args map[int]int, hashMap map[int]int, threshold int) (slice [][]int) {
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
	//fmt.Println("hash", hashMap)

	return final_pairs
}

func _printMap(args map[int]int, pairs [][]int) {
	fmt.Println("+---------------------------------------+")
	fmt.Println("\t\tPRINTING MAP\t\t")
	fmt.Println("+---------------------------------------+")
	for key, value := range args {
		fmt.Println("\tKey:", pairs[key], "Value:", value, "\t")
	}

}

func _match(pairs [][]int, support int, baskets map[int][]int, hashMap map[int]int, freqItems map[int]int, buckets int) map[int]int {

	var finalPairs map[int]int
	finalPairs = make(map[int]int)

	for _, array := range baskets { //itterates over all baskets
		for index, element := range pairs { //itterates over frequent pairs
			for x := range array {
				if array[x] == element[0] {
					for y := range array {
						if array[y] == element[1] { // x and y both frequent
							tempHash := hash(x, y, buckets)
							//			fmt.Println(tempHash, hashMap[tempHash])
							if hashMap[tempHash] == 1 {
								finalPairs[index] = finalPairs[index] + 1
							}
						}
					}
				}
			}
		}
	}

	//fmt.Println(finalPairs)

	for key := range finalPairs {
		if finalPairs[key] < support {
			delete(finalPairs, key)
		}
	}
	return finalPairs
}
